package main

import (
	"log"

	"github.com/boatnoah/faceauth/internal/store"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			addr: "",
		},
	}

	store := store.NewStorage(nil)

	a := app{config: cfg, store: store}

	mux := a.mount()
	log.Fatal(a.run(mux))
}
