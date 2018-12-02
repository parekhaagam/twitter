package main

import (
	"github.com/parekhaagam/twitter/web/controllers"
)

func main(){

	cfg := &controllers.Config{
		HTTPAddr: "localhost:9002",
	}
	err := controllers.NewStorageServer(cfg)
	if err != nil {
		panic(err)
	}

}
