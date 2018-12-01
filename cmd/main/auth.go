package main

import (
	"github.com/parekhaagam/twitter/web/auth"
)

func main(){

	cfg := &auth.Config{
		HTTPAddr: "localhost:9000",
	}
	err := auth.NewAuthServer(cfg)
	if err != nil {
		panic(err)
	}

}


