package validator

import (
	"github.com/pepeunlimited/rpc-starter-kit/pkg/todorpc"
)

type TodoServerValidator struct {}


func NewTodoServerValidator() TodoServerValidator {
	return TodoServerValidator{}
}

func (TodoServerValidator) CreateTodo(params *todorpc.CreateTodoParams) error {
	return nil
}

func (TodoServerValidator) GetTodo(params *todorpc.GetTodoParams) error {
	return nil
}

func (TodoServerValidator) UpdateTodo(params *todorpc.UpdateTodoParams) error {
	return nil
}

func (TodoServerValidator) DeleteTodo(params *todorpc.DeleteTodoParams) error {
	return nil
}