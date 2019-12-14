# gRPC-starter-kit



# Golang

* https://www.callicoder.com/golang-installation-setup-gopath-workspace/
* https://golang.github.io/dep/docs/new-project.html
* https://developers.google.com/protocol-buffers/docs/gotutorial
* https://github.com/golang/dep/blob/master/docs/Gopkg.toml.md#required
1. ```$ brew install go```
2. ```$ brew install dep```
3. ```$ dep init```
4. ```$ dep ensure```

> ```& in front of variable name is used to retrieve the address of where this variable’s value is stored. That address is what the pointer is going to store.```

> ```* in front of a type name, means that the declared variable will store an address of another variable of that type (not a value of that type).```

> ``` * in front of a variable of pointer type is used to retrieve a value stored at given address. In Go speak this is called dereferencing.```

# NATS

* https://nats.io/documentation/tutorials/gnatsd-install/
1. ```$ brew install gnatsd```

# Twrip
* https://twitchtv.github.io/twirp/docs/install.html
1. ```$ brew install protobuf```
2. ```$ go get -u github.com/golang/protobuf/protoc-gen-go```
3. ```$ go get -u github.com/twitchtv/twirp/protoc-gen-twirp```
4. ```$ protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. todo.proto```
##### If you're not using retool, you can also do a system-wide install with checking out the package new version and using go install:

1. ```$ cd $GOPATH/src/github.com/twitchtv/twirp```
2. ```$ git checkout v5.2.0```
3. ```$ go install ./protoc-gen-twirp```

# Docker

##### Download the wait-for-it.sh (https://github.com/vishnubob/wait-for-it)
```$ curl -LJO https://github.com/vishnubob/wait-for-it/raw/master/wait-for-it.sh```

##### Clean

* https://gist.github.com/bastman/5b57ddb3c11942094f8d0a97d461b430

# Git

##### Change origin

1. ```$ git remote set-url origin git@bitbucket.org:siimooo/go-starter-kit-k8.git```
2. ```$ git push --set-upstream origin master```
3. ```$ git tag -l | xargs -n 1 git push --delete origin```
4. ```$ git tag | xargs git tag -d```

##### Add tag

1. ```$ git tag -a v1.4 -m "my version 1.4"```
2. ```$ git push --tags ```

# Release

1. ```$ git add -A```
2. ```$ git commit -m "bump the version to v0.2"```
3. ```$ git tag -a v0.2 -m "upgrade 0.2v"```
4. ```$ git push```
5. ```$ git describe```
6. ```$go-starter-kit/scripts/release: sh build.sh``` 
7. ```$go-starter-kit/scripts/release: sh tag.sh {from describe}```
8. ```$go-starter-kit/scripts/release: sh push.sh {from describe}```
9. ```$ git push --tags``` 
 
 # MySQL
 
 ```$ mysql -u root -proot```
 
 # Backup
 ```$ docker exec CONTAINER /usr/bin/mysqldump -u root --password=root DATABASE > backup.sql```
 
 # Restore
 ```$ cat backup.sql | docker exec -i CONTAINER /usr/bin/mysql -u root --password=root DATABASE```
 
 #### Bash
 ``` $ docker exec -it mysql bash ```
 
 #### Text Types:
 
 #####VARCHAR(X)
    Case: user name, email, country, subject, password
 
 #####TEXT:
    Case: messages, emails, comments, formatted text, html, code, images, links
 
 #####MEDIUMTEXT:
    Case: large json bodies, short to medium length books, csv strings
 
 #####LONGTEXT:
    Case: textbooks, programs, years of logs files, harry potter and the goblet of fire, scientific research logging
 
 #### Logs:
 ```$ SET global general_log = 1;```
 ```$ SHOW VARIABLES LIKE "general_log%";```

### Minikube
1. Install the hyperkit VM manager using brew: ```$ brew install hyperkit``` 
2. Then install the most recent version of minikube's fork of the hyperkit driver:
```$ curl -LO https://storage.googleapis.com/minikube/releases/latest/docker-machine-driver-hyperkit && sudo install -o root -g wheel -m 4755 docker-machine-driver-hyperkit /usr/local/bin/```
3. ```$ minikube start --vm-driver hyperkit``` 
or, to use hyperkit as a default driver for minikube:
4. ```$ minikube config set vm-driver hyperkit```

# Utils
1. Recursive delete dir ```$ rm -rf some_dir/```
2. ```$ ps aux | grep 'go'```
3. ```$ kill -9 {PID}```

#### HyperKit time issue fix
``` $ DSTR=$(date -u); minikube ssh "sudo date --set=\"$DSTR\"" ```