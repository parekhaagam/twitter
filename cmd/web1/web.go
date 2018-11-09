package main

import (
	"../../web"
	"time"
)

func main(){

	cfg := &web.Config{
		HTTPAddr: "localhost:8080",
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


