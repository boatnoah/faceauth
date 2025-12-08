package main

import (
	"github.com/boatnoah/faceauth/internal/store"
	"go.uber.org/zap"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			addr: "testing.string.",
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	store := store.NewStorage(nil)

	a := app{config: cfg, store: store, logger: logger}

	mux := a.mount()
	logger.Fatal(a.run(mux))
}
