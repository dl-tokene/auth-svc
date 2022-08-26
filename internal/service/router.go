package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/nonce-auth-svc/internal/config"
	"gitlab.com/tokene/nonce-auth-svc/internal/data/pg"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/handlers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxDB(pg.NewMasterQ(cfg.DB())),
			helpers.CtxServiceConfig(cfg.ServiceConfig()),
		),
	)
	r.Route("/integrations/nonce-auth-svc", func(r chi.Router) {
		r.Post("/nonce", handlers.GetNonce)
		r.Post("/register", handlers.Register)
		r.Post("/refresh_token", handlers.RefreshToken)
		r.Post("/login", handlers.Login)
	})

	return r
}
