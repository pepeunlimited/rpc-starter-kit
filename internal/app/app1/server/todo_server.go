package server

import (
	"context"
	"github.com/pepeunlimited/rpc-starter-kit/internal/app/app1/validator"
	"github.com/pepeunlimited/rpc-starter-kit/rpctodo"
	"github.com/twitchtv/twirp"
)

type TodoServer struct {
	todos map[int64]*rpctodo.Todo
	validator validator.TodoServerValidator
}

func (server TodoServer) UpdateTodo(ctx context.Context, params *rpctodo.UpdateTodoParams) (*rpctodo.Todo, error) {
	return nil, nil
}

func (server TodoServer) DeleteTodo(ctx context.Context, params *rpctodo.DeleteTodoParams) (*rpctodo.Todo, error) {
	return nil, nil
}

func (server TodoServer) GetTodo(ctx context.Context, params *rpctodo.GetTodoParams) (*rpctodo.Todo, error) {
	if err := server.validator.GetTodo(params); err != nil {
		return nil, err
	}
	todo := server.todos[params.Id]
	if todo == nil {
		return nil, twirp.NotFoundError("todo's not exist")
	}
	return todo, nil
}

func (server TodoServer) CreateTodo(ctx context.Context, params *rpctodo.CreateTodoParams) (*rpctodo.Todo, error) {
	if err := server.validator.CreateTodo(params); err != nil {
		return nil, err
	}
	server.todos[params.Todo.Id] = params.Todo
	return params.Todo, nil
}

func NewTodoServer() TodoServer {
	todos := make(map[int64]*rpctodo.Todo)
	return TodoServer{
		todos: todos,
		validator: validator.NewTodoServerValidator(),
	}
}