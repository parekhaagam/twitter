package web

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"./controllers"
)

type Web struct {
	srv *http.Server

	//contentDaemon *contentd.Contentd
}

type loginPage struct {
	Email    string
	Password string
}

var userRecord  = &UsersRecord{
	users: make(map[string]string),
}

func signup(w http.ResponseWriter, r *http.Request) {


	t, err := template.ParseFiles("web/html/signup.html")
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

	t:= template.Must(template.ParseFiles("web/html/login.html"))
	mLoginPage := loginPage{
		Email:    "EmailId",
		Password: "password",
	}
	err := t.Execute(w, mLoginPage)
	if err != nil {
		log.Print("error while executing ", err)
	}

}


func validateLogin(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	emailId := strings.Join(r.Form["EmailId"], "")
	password := strings.Join(r.Form["password"], "")

	if userRecord.userExist(emailId, password){
		fmt.Println("")
		w.Write([]byte("valid userid"))
	}else{
		w.Write([]byte("Invalid UserId"))
	}

}

func validateSignup(w http.ResponseWriter, r *http.Request){

	r.ParseForm()
	emailId := strings.Join(r.Form["EmailId"], "")
	password := strings.Join(r.Form["password"], "")


	if userRecord.insertUser(emailId, password){
		w.Write([]byte("valid userid"))
	}else{
		w.Write([]byte("Invalid userid"))
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
	mx.HandleFunc("/signupValidation", validateSignup)
	mx.HandleFunc("/loginValidation", validateLogin)
	mx.HandleFunc("/*", signup)
	mx.HandleFunc("/show-users", controllers.Show_users)
	return ws, nil
}

func (w *Web) Start() error {
	return w.srv.ListenAndServe()
}

func (w *Web) Shutdown(ctx context.Context) error {
	return w.srv.Shutdown(ctx)
}