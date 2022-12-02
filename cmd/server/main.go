package main

import (
	"context"

	"github.com/kevinmidboe/planetposen-mail/config"
	"github.com/kevinmidboe/planetposen-mail/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	// log.SetFormatter(logrustic.NewFormatter("planetposen-mail"))

	log.Info("Starting ...")

	ctx := context.Background()
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	var s server.Server

	if err := s.Create(ctx, config); err != nil {
		log.Fatal(err.Error())
	}

	if err := s.Serve(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
