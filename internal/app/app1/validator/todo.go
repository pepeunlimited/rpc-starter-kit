package validator

import (
	"github.com/pepeunlimited/rpc-starter-kit/rpc"
)

type TodoServerValidator struct {}


func NewTodoServerValidator() TodoServerValidator {
	return TodoServerValidator{}
}

func (TodoServerValidator) CreateTodo(params *rpc.CreateTodoParams) error {
	return nil
}

func (TodoServerValidator) GetTodo(params *rpc.GetTodoParams) error {
	return nil
}

func (TodoServerValidator) UpdateTodo(params *rpc.UpdateTodoParams) error {
	return nil
}

func (TodoServerValidator) DeleteTodo(params *rpc.DeleteTodoParams) error {
	return nil
}