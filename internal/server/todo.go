package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alcb1310/todo/internal/database"
)

func (s TodoService) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := s.DB.GetAllTodos()
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(todos)
}

func (s TodoService) CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := database.Todo{
		Title:     "New Todo",
		Completed: false,
	}

	s.DB.CreateTodo(&todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (s TodoService) GetOneTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Getting todo %s", id)))
}

func (s TodoService) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Updating todo %s", id)))
}

func (s TodoService) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "id")
	w.WriteHeader(http.StatusNoContent)
}
