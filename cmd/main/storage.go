package main

import (
	"github.com/parekhaagam/twitter/constants"
	"github.com/parekhaagam/twitter/web/controllers"
)

func main(){

	cfg := &controllers.Config{
		HTTPAddr: constants.StorageServerEndpoint,
	}
	err := controllers.NewStorageServer(cfg)
	if err != nil {
		panic(err)
	}

}
