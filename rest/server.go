package rest

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Nariett/arox-pkg/config"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func StartHttpServer(lc fx.Lifecycle, cfg *config.Config, r chi.Router, shutdowner fx.Shutdowner) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("Starting HTTP server at :%s\n", cfg.Port)
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Printf("HTTP server error: %v", err)
					_ = shutdowner.Shutdown()
				}
			}()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down HTTP server...")

			return srv.Shutdown(ctx)
		},
	})
}
