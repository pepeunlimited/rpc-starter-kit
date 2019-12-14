package server

import (
	"context"
	"github.com/pepeunlimited/rpc-starter-kit/rpc"
	"testing"
	"time"
)

func TestTodoServer_CreateTodo(t *testing.T) {
	server := NewTodoServer()
	todo, err := server.CreateTodo(context.TODO(), &rpc.CreateTodoParams{Todo: &rpc.Todo{
		Name:      "HelloWorld!",
		IsDone:    true,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
		IsDeleted: false,
		Id:        1,
	}})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if server.todos[1] == nil {
		t.FailNow()
	}
	if server.todos[1].Name != todo.Name {
		t.FailNow()
	}
}
