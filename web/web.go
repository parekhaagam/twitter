package web

import (
	"./controllers"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"../globals"
	"html/template"
)

type Web struct {
	srv *http.Server

	//contentDaemon *contentd.Contentd
}

type loginPage struct {
	Email    string
	Password string
}

func initGlobals() {
	globals.Followers = make(map[string][]globals.User)
	globals.UsersRecord = make(map[string]string)
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

	if controllers.UserExist(emailId, password){
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


	if controllers.InsertUser(emailId, password){
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

	initGlobals()
	mx.HandleFunc("/", controllers.Login)
	mx.HandleFunc("/signup", controllers.Signup)
	mx.HandleFunc("/signupValidation", controllers.ValidateSignup)
	mx.HandleFunc("/loginValidation", controllers.ValidateLogin)
	mx.HandleFunc("/*", controllers.Signup)
	mx.HandleFunc("/show-users", controllers.Show_users)
	mx.HandleFunc("/follow", controllers.Follow_users)
	return ws, nil
}

func (w *Web) Start() error {
	return w.srv.ListenAndServe()
}

func (w *Web) Shutdown(ctx context.Context) error {
	return w.srv.Shutdown(ctx)
}