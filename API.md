# cURL

### Install
```$ brew install jq > curl ... | jq```

##### CreateTodo
```
$ curl -H "Content-Type: application/json" \
-X POST "localhost:8080/twirp/pepeunlimited.todo.TodoService/CreateTodo" \
-d '{"standard_vat": "24.00"}'
```