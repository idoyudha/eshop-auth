package main

import (
	"log"

	"github.com/idoyudha/eshop-auth/config"
	"github.com/idoyudha/eshop-auth/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
