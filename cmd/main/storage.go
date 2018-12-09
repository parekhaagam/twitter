package main

import (
	"github.com/parekhaagam/twitter/app_server"
	"github.com/parekhaagam/twitter/constants"
)

func main(){

	cfg := &app_server.Config{
		HTTPAddr: constants.StorageServerEndpoint,
	}
	err := app_server.NewStorageServer(cfg)
	if err != nil {
		panic(err)
	}

}
