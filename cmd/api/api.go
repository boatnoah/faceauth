package main

import (
	"log"
	"net/http"
	"time"

	"errors"

	"github.com/boatnoah/faceauth/internal/auth"
	"github.com/boatnoah/faceauth/internal/store"
	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
)

type app struct {
	config        config
	store         store.Storage
	authenticator auth.Authenticator
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

	log.Printf("Server has started addr=%v", a.config.addr)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	log.Printf("Server has stopped addr=%v", a.config.addr)

	return nil

}
