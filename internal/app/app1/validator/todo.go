package validator

import (
	"github.com/pepeunlimited/rpc-starter-kit/rpctodo"
)

type TodoServerValidator struct {}


func NewTodoServerValidator() TodoServerValidator {
	return TodoServerValidator{}
}

func (TodoServerValidator) CreateTodo(params *rpctodo.CreateTodoParams) error {
	return nil
}

func (TodoServerValidator) GetTodo(params *rpctodo.GetTodoParams) error {
	return nil
}

func (TodoServerValidator) UpdateTodo(params *rpctodo.UpdateTodoParams) error {
	return nil
}

func (TodoServerValidator) DeleteTodo(params *rpctodo.DeleteTodoParams) error {
	return nil
}