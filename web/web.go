package web

import (
	"../globals"
	"./auth"
	"./controllers"
	"context"
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

type Web struct {
	srv *http.Server

	//contentDaemon *contentd.Contentd
}

type loginPage struct {
	Email    string
	Password string
}

func InitGlobals() {
	globals.Followers = make(map[string][]globals.User)
	globals.UsersRecord = make(map[string]string)
	globals.UserTweet = make(map[string][]globals.Tweet)
	globals.AllUsers = insertDummies()
}

func insertDummies() (allUsers[] globals.User){
	globals.UsersRecord["manish.n"] = "admin"
	globals.UsersRecord["dhoni007"] = "admin"
	globals.UsersRecord["srk"] = "admin"
	globals.UsersRecord["chandler"] = "admin"

	allUsers = append(allUsers, globals.User{"manish.n"})
	allUsers = append(allUsers, globals.User{"dhoni007"})
	allUsers = append(allUsers, globals.User{"srk"})
	allUsers = append(allUsers, globals.User{"chandler"})


	tweet1 := globals.Tweet{TID:uuid.New().String(), Content:"Mujhe bhi T20 khelna h", Timestamp:time.Now().Unix(), UserId:"dhoni007"}
	tweet2 := globals.Tweet{TID:uuid.New().String(), Content:"Zero releasing December 2018", Timestamp:time.Now().Unix(), UserId:"srk"}
	tweet3 := globals.Tweet{TID:uuid.New().String(), Content:"Could I be wearing anymore clothes", Timestamp:time.Now().Unix(), UserId:"chandler"}
	tweet4 := globals.Tweet{TID:uuid.New().String(), Content:"SDE at Google", Timestamp:time.Now().Unix(), UserId:"manish.n"}
	tweet5 := globals.Tweet{TID:uuid.New().String(), Content:"Virat ne mujhe nikal diya T20 team se", Timestamp:time.Now().Unix(), UserId:"dhoni007"}

	globals.UserTweet["dhoni007"] = append(globals.UserTweet["dhoni007"], tweet1)
	globals.UserTweet["dhoni007"] = append(globals.UserTweet["dhoni007"], tweet5)
	globals.UserTweet["manish.n"] = append(globals.UserTweet["manish.n"], tweet4)
	globals.UserTweet["chandler"] = append(globals.UserTweet["chandler"], tweet3)
	globals.UserTweet["srk"] = append(globals.UserTweet["srk"], tweet2)

	return allUsers
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

	InitGlobals()
	mx.HandleFunc("/twitter", controllers.Login)
	mx.HandleFunc("/signup", controllers.Signup)
	mx.HandleFunc("/signupValidation", controllers.ValidateSignup(controllers.Show_users))
	mx.HandleFunc("/loginValidation", controllers.ValidateLogin(controllers.Feed))
	mx.HandleFunc("/*", controllers.Signup)
	mx.HandleFunc("/show-users", auth.AuthenticationMiddleware(controllers.Show_users))
	mx.HandleFunc("/follow", auth.AuthenticationMiddleware(controllers.Follow_users(controllers.Feed)))
	mx.HandleFunc("/feed", auth.AuthenticationMiddleware(controllers.Feed))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	return ws, nil

}

func (w *Web) Start() error {
	return w.srv.ListenAndServe()
}

func (w *Web) Shutdown(ctx context.Context) error {
	return w.srv.Shutdown(ctx)
}