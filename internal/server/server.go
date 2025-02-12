package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

const ctxTimeout = 5

type Server struct {
	mux *mux.Router
	db  *sqlx.DB
}

func New(db *sqlx.DB) *Server {
	return &Server{
		mux: mux.NewRouter(),
		db:  db,
	}
}

func (s *Server) Start() error {
	srv := http.Server{
		Addr:         ":8080",
		Handler:      s.mux,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 200 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Can't start http server: %v", err)
		}
	}()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM)
	<-quitCh

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	return srv.Shutdown(ctx)
}
