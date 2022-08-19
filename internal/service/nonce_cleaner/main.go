package nonce_cleaner

import (
	"context"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokene/nonce-auth-svc/internal/config"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	"gitlab.com/tokene/nonce-auth-svc/internal/data/pg"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/types"
)

func NewNonceCleaner(cfg config.Config) types.Service {
	return nonceCleaner{
		q:      pg.NewMasterQ(cfg.DB()).Nonce(),
		logger: cfg.Log(),
	}
}

type nonceCleaner struct {
	q      data.NonceQ
	logger *logan.Entry
}

func (s nonceCleaner) Run(ctx context.Context) error {
	running.WithBackOff(context.TODO(),
		logan.New(),
		"nonce-cleaner",
		s.runNonceCleaner,
		3*time.Second,
		3*time.Second,
		3*time.Second,
	)
	return nil
}

func (s nonceCleaner) runNonceCleaner(ctx context.Context) error {
	// filter must be applied only once

	for {
		s.q.FilterExpired()
		s.q.Delete()
		time.Sleep(time.Second)
	}
}
