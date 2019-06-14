# How to install

1. install xampp for mac
2. in phpMyAdmin, create DB `DreamTracker` and import from sql in the root of the repo
3. install go `brew install go`
4. make sure you have GOPATH properly set-up, for simplicity use this command `export GOPATH=$HOME/go`
5. and move to that gopath `cd $GOPATH/src`
6. clone this repo into folder `server`
7. move into that `server` folder
8. run `go get`
9. make sure you check `config/db.go` and `config/server.go` and change if needed
10. run `go run main.go` to run in your local machine
11. now you can use localhost:8088/ as your base url
