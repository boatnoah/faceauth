package main

import (
	"net/http"
	"time"

	"errors"
	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type app struct {
	config config
	logger *zap.SugaredLogger
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr string
}

func (a *app) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", a.healthCheckHandler)
	})
	return r
}

func (a *app) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	a.logger.Infow("Server has started", "addr", a.config.addr)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	a.logger.Infow("Server has stopped", "addr", a.config.addr)

	return nil

}
