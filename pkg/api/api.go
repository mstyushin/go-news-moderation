package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/mstyushin/go-news-moderation/pkg/config"

	"github.com/gorilla/mux"
)

type API struct {
	cfg *config.Config
	mux *mux.Router
}

func New(cfg *config.Config) *API {
	api := API{
		cfg: cfg,
		mux: mux.NewRouter(),
	}

	return &api
}

func (api *API) Run(ctx context.Context) error {
	errChan := make(chan error)
	srv := api.serve(ctx, errChan)

	select {
	case <-ctx.Done():
		log.Println("gracefully shutting down")
		srv.Shutdown(ctx)
		return ctx.Err()
	case err := <-errChan:
		log.Println(err)
		return err
	}
}

func (api *API) serve(ctx context.Context, errChan chan error) *http.Server {
	api.initMux()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", api.cfg.HttpPort),
		Handler: api.mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, fmt.Sprintf(":%v", api.cfg.HttpPort), l.Addr().String())
			return ctx
		},
	}

	go func(s *http.Server) {
		if err := s.ListenAndServe(); err != nil {
			errChan <- err
		}
	}(httpServer)

	log.Println("serving HTTP server at", api.cfg.HttpPort)

	return httpServer
}

func (api *API) initMux() {
	api.mux.HandleFunc("/moderation", api.check).Methods(http.MethodPost, http.MethodOptions)
	api.mux.Use(URLSchemaMiddleware(api.mux))
	api.mux.Use(RequestIDLoggerMiddleware(api.mux))
	api.mux.Use(LoggerMiddleware(api.mux))
}
