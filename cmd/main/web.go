package main

import (
	"github.com/parekhaagam/twitter/globals"
	"github.com/parekhaagam/twitter/web_server"
	"time"
)

func main(){

	cfg := &web_server.Config{
		HTTPAddr: globals.WebEndpoint,
	}

	webSrv, err := web_server.New(cfg)
	if err != nil {
		panic(err)
	}

	err = webSrv.Start()
	if err != nil {
		panic(err)
	}
	time.Sleep(115 * time.Second)
}


