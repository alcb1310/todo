package database

import "errors"

type Todo struct {
	ID        uint
	Title     string
	Completed bool
}

type todoService struct {
	todos  []Todo
	nextId uint
}

type TodoDatabaseService interface {
	GetAllTodos() []Todo
	CreateTodo(todo *Todo) error
	GetOneTodo(id uint) (Todo, error)
	UpdateTodo(id uint, todo *Todo) error
	DeleteTodo(id uint) error
}

func NewTodoDatabaseService() TodoDatabaseService {
	return &todoService{
		todos:  []Todo{},
		nextId: 1,
	}
}

func (s *todoService) GetAllTodos() []Todo {
	return s.todos
}

func (s *todoService) CreateTodo(todo *Todo) error {
	todo.ID = s.nextId
	s.nextId++

	s.todos = append(s.todos, *todo)
	return nil
}

func (s *todoService) GetOneTodo(id uint) (Todo, error) {
	for _, todo := range s.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return Todo{}, errors.New("todo not found")
}

func (s *todoService) UpdateTodo(id uint, todo *Todo) error {
	foundTodo := Todo{}
	todoIdx := -1

	for i, t := range s.todos {
		if t.ID == id {
			foundTodo = t
			todoIdx = i
			break
		}
	}
	if todoIdx == -1 {
		return errors.New("todo not found")
	}

	foundTodo.ID = todo.ID
	foundTodo.Title = todo.Title
	foundTodo.Completed = todo.Completed

	s.todos[todoIdx] = foundTodo

	return nil
}

func (s *todoService) DeleteTodo(id uint) error {
	todoIdx := -1

	for i, t := range s.todos {
		if t.ID == id {
			todoIdx = i
			break
		}
	}
	if todoIdx == -1 {
		return errors.New("todo not found")
	}

	s.todos = append(s.todos[:todoIdx], s.todos[todoIdx+1:]...)
	return nil
}
