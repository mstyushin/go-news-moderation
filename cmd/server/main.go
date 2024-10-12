package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mstyushin/go-news-moderation/pkg/api"
	"github.com/mstyushin/go-news-moderation/pkg/config"
)

const (
	AppName = "go-news-moderation"
)

type server struct {
	api *api.API
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	if cfg == nil {
		os.Exit(0)
	}

	log.Printf("starting %s service\n", AppName)
	log.Println(config.VersionString())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		cancel()
	}()

	server := api.New(cfg)
	if err := server.Run(ctx); err != nil {
		log.Println("Got error:", err)
		os.Exit(0)
	}
}
