package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"gitlab.com/distributed_lab/ape"
	gosdk "gitlab.com/tokene/go-sdk"
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
			helpers.CtxNodeAdmins(gosdk.NewNodeAdminsMock(common.HexToAddress("0x750Bd531CEA1f68418DDF2373193CfbD86A69058"))), //TODO change when admin's smart contracts ready
			helpers.CtxDoormanConnector(cfg.DoormanConnector()),
		),
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		}),
	)
	r.Route("/integrations/nonce-auth-svc", func(r chi.Router) {
		r.Post("/nonce", handlers.GetNonce)
		r.Post("/register", handlers.Register)
		r.Post("/refresh_token", handlers.RefreshToken)
		r.Post("/login", handlers.Login)
		r.Post("/admin_login", handlers.AdminLogin)
		r.Post("/created_at", handlers.CreatedAt)
	})

	return r
}
