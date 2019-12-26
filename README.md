# rpc-stater-kit

A starting point project to create `todo`-service

## Go Directories

### `/build`

### `/cmd`

### `/deployments`
Deployments folder contains files for the `k8s` environment; `service`, `deployment` and `ingress`. `dev` folder have an additional `external-mysql-service.yaml` file which allow access from local cluster to external MySQL database.  

`docker-compose up` spin up the MySQL database for local development 

### `/init`

### `/internal`

#### `/ent`
Speed up implementing the database access using [`ent`](https://github.com/facebookincubator/ent). Of course you can implement repositories with a raw sql statements, but it is very time consuming and boring repeat x10 times same CRUD functions.

#### [`ent`](https://github.com/facebookincubator/ent)
- `$ entc generate ./ent/schema`

### `/rpc`

#### [`twirp`](https://github.com/twitchtv/twirp)
-  `$ protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. user.proto`


### `/scripts`

#### `misc`
```$ brew install jq > curl ... | jq```
##### Show MySQL's logs
```$ SET global general_log = 1;```  
```$ SHOW VARIABLES LIKE "general_log%";```