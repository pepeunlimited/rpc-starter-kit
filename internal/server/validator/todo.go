package validator

import (
	"github.com/pepeunlimited/rpc-starter-kit/pkg/rpc/todo"
)

type TodoServerValidator struct {}


func NewTodoServerValidator() TodoServerValidator {
	return TodoServerValidator{}
}

func (TodoServerValidator) CreateTodo(params *todo.CreateTodoParams) error {
	return nil
}

func (TodoServerValidator) GetTodo(params *todo.GetTodoParams) error {
	return nil
}

func (TodoServerValidator) UpdateTodo(params *todo.UpdateTodoParams) error {
	return nil
}

func (TodoServerValidator) DeleteTodo(params *todo.DeleteTodoParams) error {
	return nil
}