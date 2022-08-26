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
	s.logger.Debug("Calling delete of expired nonces")
	running.WithBackOff(ctx,
		s.logger,
		"nonce-cleaner",
		s.runNonceCleaner,
		1*time.Minute,
		1*time.Second,
		5*time.Second,
	)

	return nil
}

func (s nonceCleaner) runNonceCleaner(ctx context.Context) error {
	s.q.FilterExpired() //Clearing previous sql condition inside function
	return s.q.Delete()
}
