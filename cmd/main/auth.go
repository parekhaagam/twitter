package main

import (
	"github.com/parekhaagam/twitter/auth_server"
	"github.com/parekhaagam/twitter/globals"
)

func main(){

	//contract.Get()
	cfg := &auth_server.Config{
		HTTPAddr: globals.AuthServerEndpoint,
	}
	err := auth_server.NewAuthServer(cfg)
	if err != nil {
		panic(err)
	}



}


