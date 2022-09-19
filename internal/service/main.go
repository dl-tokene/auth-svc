package service

import (
	"context"
	"net"
	"net/http"
	"sync"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/nonce-auth-svc/internal/config"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/nonce_cleaner"
	serv "gitlab.com/tokene/nonce-auth-svc/internal/service/types"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
}

func (s *service) run(cfg config.Config) error {
	s.log.Info("Service started")
	r := s.router(cfg)

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}
func runService(service serv.Service, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		_ = service.Run(context.Background())
	}()
}
func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
	}
}

func Run(cfg config.Config) {
	wg := &sync.WaitGroup{}

	nonceCleaner := nonce_cleaner.NewNonceCleaner(cfg)
	runService(nonceCleaner, wg)

	if err := newService(cfg).run(cfg); err != nil {
		panic(err)
	}
}
