package main

import (
	"github.com/parekhaagam/twitter/constants"
	"github.com/parekhaagam/twitter/web/auth"
	"github.com/parekhaagam/twitter/web/auth/storage"
)

func main(){

	storage.Get()
	cfg := &auth.Config{
		HTTPAddr: constants.AuthServerEndpoint,
	}
	err := auth.NewAuthServer(cfg)
	if err != nil {
		panic(err)
	}



}


