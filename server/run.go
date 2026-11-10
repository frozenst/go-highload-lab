// Package server contains the HTTP server setup and lifecycle.
package server

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/frozenst/go-highload-lab/server/handles"
)

// newRouter registers all application routes.
func newRouter() *router.Router {
	r := router.New()
	r.GET("/health", handles.Health)
	return r
}

// Start runs the HTTP server and blocks until SIGINT or SIGTERM.
func Start() {
	srv := &fasthttp.Server{
		Handler:      newRouter().Handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		slog.Info("server started", "addr", ":53347")
		if err := srv.ListenAndServe(":53347"); err != nil {
			slog.Error("server failed", "err", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.ShutdownWithContext(ctx); err != nil {
		slog.Error("shutdown failed", "err", err)
	}
}
