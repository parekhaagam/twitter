package main

import (
	"github.com/parekhaagam/twitter/constants"
	"github.com/parekhaagam/twitter/web/auth"
)

func main(){

	cfg := &auth.Config{
		HTTPAddr: constants.AuthServerEndpoint,
	}
	err := auth.NewAuthServer(cfg)
	if err != nil {
		panic(err)
	}

}


