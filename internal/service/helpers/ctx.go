package helpers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokene/doorman/connector"
	gosdk "gitlab.com/tokene/go-sdk"
	"gitlab.com/tokene/nonce-auth-svc/internal/config"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	dbCtxKey
	nodeAdminsCtxKey
	serviceConfigCtxKey
	doormanConnectorCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}
func CtxDB(entry data.MasterQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbCtxKey, entry)
	}
}
func DB(r *http.Request) data.MasterQ {
	return r.Context().Value(dbCtxKey).(data.MasterQ).New()
}
func CtxServiceConfig(entry *config.ServiceConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, serviceConfigCtxKey, entry)
	}
}
func CtxNodeAdmins(entry gosdk.NodeAdminsI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, nodeAdminsCtxKey, entry)
	}
}
func NodeAdmins(r *http.Request) gosdk.NodeAdminsI {
	return r.Context().Value(nodeAdminsCtxKey).(gosdk.NodeAdminsI)
}
func ServiceConfig(r *http.Request) *config.ServiceConfig {
	return r.Context().Value(serviceConfigCtxKey).(*config.ServiceConfig)
}
func CtxDoormanConnector(entry connector.ConnectorI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, doormanConnectorCtxKey, entry)
	}
}
func DoormanConnector(r *http.Request) connector.ConnectorI {
	return r.Context().Value(doormanConnectorCtxKey).(connector.ConnectorI)
}
