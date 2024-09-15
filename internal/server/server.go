package server

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/alcb1310/todo/internal/database"
	"github.com/go-chi/chi/v5"
)

type TodoService struct {
  DB database.TodoDatabaseService
}

func NewServer() {
	r := chi.NewRouter()

  ts := TodoService{
    DB: database.NewTodoDatabaseService(),
  }

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World using air"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/todos", func(r chi.Router) {
        r.Get("/", ts.GetAllTodos)
				r.Post("/", ts.CreateTodo)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", ts.GetOneTodo)
          r.Put("/", ts.UpdateTodo)
          r.Delete("/", ts.DeleteTodo)
				})
			})
		})
	})

	slog.Info("Starting server", "port", ":3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		slog.Error("Error starting the server", "err", err)
		os.Exit(1)
	}
}
