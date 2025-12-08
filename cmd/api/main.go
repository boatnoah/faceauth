package main

import (
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

	a := app{config: cfg, logger: logger}

	mux := a.mount()
	logger.Fatal(a.run(mux))

}
