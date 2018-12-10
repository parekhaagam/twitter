package main

import (
	"github.com/parekhaagam/twitter/app_server"
	"github.com/parekhaagam/twitter/globals"
	"log"
)

func main(){

	cfg := &app_server.Config{
		HTTPAddr: globals.StorageServerEndpoint,
	}
	err := app_server.NewStorageServer(cfg)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

}
