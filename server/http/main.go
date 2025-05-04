package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	authhttp "nookli/server/http/auth"
	// wshttp "nookli/server/http/workspace"
)

func main() {
	r := chi.NewRouter()
	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(15*time.Second),
	)

	// Health check
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// Public Auth
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Mount("/", authhttp.Router())
	})

	// Protected Workspaces
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(authhttp.RequireAuth) // your auth middleware
		// r.Mount("/workspaces", wshttp.Router())
	})

	srv := &http.Server{Addr: ":8080", Handler: r}
	go func() {
		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Shutting down…")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
