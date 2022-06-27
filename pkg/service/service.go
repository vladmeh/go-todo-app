package service

import (
	"github.com/vladmeh/go-todo-app"
	"github.com/vladmeh/go-todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, id int) (todo.TodoList, error)
	Update(userId int, id int, input todo.UpdateListInput) error
	Delete(userId int, id int) error
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Update(userId, itemId int, input todo.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
		TodoList:      NewTodoListService(r.TodoList),
		TodoItem:      NewTodoItemService(r.TodoItem, r.TodoList),
	}
}
