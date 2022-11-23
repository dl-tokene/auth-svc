package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	gosdk "gitlab.com/tokene/go-sdk"
	"gitlab.com/tokene/nonce-auth-svc/internal/config"
	"gitlab.com/tokene/nonce-auth-svc/internal/data/pg"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/handlers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/util"
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
			helpers.CtxNodeAdmins(gosdk.NewNodeAdminsMock(util.StringSliceToAddresses(cfg.AdminsConfig().Admins)...)), //TODO change when admin's smart contracts ready
			helpers.CtxDoormanConnector(cfg.DoormanConnector()),
		),
	)
	r.Route("/integrations/nonce-auth-svc", func(r chi.Router) {
		r.Post("/nonce", handlers.GetNonce)
		r.Get("/refresh_token", handlers.RefreshToken)
		r.Post("/admin_login", handlers.AdminLogin)

	})

	return r
}
