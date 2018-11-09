package web

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Web struct {
	srv *http.Server

	//contentDaemon *contentd.Contentd
}

type loginPage struct {
	Email    string
	Password string
}


func signup(w http.ResponseWriter, r *http.Request) {


	t, err := template.ParseFiles("web/signup.html")
	if err != nil{
		log.Print("sign up page not loaded properly", err)
	}
	mLoginPage := loginPage{
		Email:    "EmailId",
		Password: "password",
	}
	err = t.Execute(w, mLoginPage)
	if err != nil {
		log.Print("error while executing ", err)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(".")
	fmt.Println(files)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	log.Print("inside login function")

	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
//	templates := template.Must(template.ParseFiles(os.Getwd("templates/index.html")))
	t:= template.Must(template.ParseFiles("web/login.html"))
/*	if err != nil {
		log.Print("Login page not loaded properly", err)
	}
*/
	mLoginPage := loginPage{
		Email:    "EmailId",
		Password: "password",
	}
	err = t.Execute(w, mLoginPage)
	if err != nil {
		log.Print("error while executing ", err)
	}

}

func New(cfg *Config) (*Web, error) {
	mx := http.NewServeMux()
	s := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: mx,
	}

	ws := &Web{
		srv: s,
	}

	mx.HandleFunc("/", login)
	mx.HandleFunc("/signup", signup)

	return ws, nil
}

func (w *Web) Start() error {
	return w.srv.ListenAndServe()
}

func (w *Web) Shutdown(ctx context.Context) error {
	return w.srv.Shutdown(ctx)
}