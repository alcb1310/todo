package server

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func NewServer() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World using air"))
	})

	slog.Info("Starting server", "port", ":3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		slog.Error("Error starting the server", "err", err)
		os.Exit(1)
	}
}
