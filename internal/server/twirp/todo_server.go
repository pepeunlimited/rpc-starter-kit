package twirp

import (
	"context"
	"github.com/pepeunlimited/rpc-starter-kit/internal/server/validator"
	"github.com/pepeunlimited/rpc-starter-kit/pkg/todorpc"
	"github.com/twitchtv/twirp"
)

type TodoServer struct {
	todos     map[int64]*todorpc.Todo
	validator validator.TodoServerValidator
}

func (server TodoServer) UpdateTodo(ctx context.Context, params *todorpc.UpdateTodoParams) (*todorpc.Todo, error) {
	return nil, nil
}

func (server TodoServer) DeleteTodo(ctx context.Context, params *todorpc.DeleteTodoParams) (*todorpc.Todo, error) {
	return nil, nil
}

func (server TodoServer) GetTodo(ctx context.Context, params *todorpc.GetTodoParams) (*todorpc.Todo, error) {
	if err := server.validator.GetTodo(params); err != nil {
		return nil, err
	}
	todo := server.todos[params.Id]
	if todo == nil {
		return nil, twirp.NotFoundError("todo's not exist")
	}
	return todo, nil
}

func (server TodoServer) CreateTodo(ctx context.Context, params *todorpc.CreateTodoParams) (*todorpc.Todo, error) {
	if err := server.validator.CreateTodo(params); err != nil {
		return nil, err
	}
	server.todos[params.Todo.Id] = params.Todo
	return params.Todo, nil
}

func NewTodoServer() TodoServer {
	todos := make(map[int64]*todorpc.Todo)
	return TodoServer{
		todos:     todos,
		validator: validator.NewTodoServerValidator(),
	}
}