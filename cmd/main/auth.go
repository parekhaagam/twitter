package main

import (
	"github.com/parekhaagam/twitter/constants"
	"github.com/parekhaagam/twitter/auth_server"
)

func main(){

	//storage.Get()
	cfg := &auth_server.Config{
		HTTPAddr: constants.AuthServerEndpoint,
	}
	err := auth_server.NewAuthServer(cfg)
	if err != nil {
		panic(err)
	}



}


