package main

import (
	"github.com/parekhaagam/twitter/web"
	"time"
)

func main(){

	cfg := &web.Config{
		HTTPAddr: "localhost:8090",
	}

	webSrv, err := web.New(cfg)
	if err != nil {
		panic(err)
	}

	err = webSrv.Start()
	if err != nil {
		panic(err)
	}
	time.Sleep(115 * time.Second)
}


