package helpers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokene/doorman/internal/config"
	gosdk "gitlab.com/tokene/go-sdk"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	serviceConfigCtxKey
	nodeAdminsCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}
func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxServiceConfig(entry *config.ServiceConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, serviceConfigCtxKey, entry)
	}
}
func ServiceConfig(r *http.Request) *config.ServiceConfig {
	return r.Context().Value(serviceConfigCtxKey).(*config.ServiceConfig)
}

func CtxNodeAdmins(entry gosdk.NodeAdminsI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, nodeAdminsCtxKey, entry)
	}
}
func NodeAdmins(r *http.Request) gosdk.NodeAdminsI {
	return r.Context().Value(nodeAdminsCtxKey).(gosdk.NodeAdminsI)
}
