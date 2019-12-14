# rpc-starter-kit (twirp)

A starter kit project to create services using `rpc` with a ([`twirp`](https://github.com/twitchtv/twirp)).

## Go Directories

### `/build`

### `/cmd`

### `/deployments`

### `/init`

### `/internal`

#### `/ent`
Speed up implementing the database access using [`ent`](https://github.com/facebookincubator/ent). Of course you can implement repositories with a raw sql's queries, but it is very time consuming and boring repeat x10 times same CRUD functions.

#### [`ent`](https://github.com/facebookincubator/ent)
- `$ entc generate ./ent/schema`

### `/rpc`

#### [`twirp`](https://github.com/twitchtv/twirp)
-  `$ protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. todo.proto`


### `/scripts`

#### `misc`
```$ brew install jq > curl ... | jq```