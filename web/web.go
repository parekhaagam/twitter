package web

import (
	"github.com/parekhaagam/twitter/web/controllers"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
	"strings"
	"github.com/parekhaagam/twitter/globals"
	pb "github.com/parekhaagam/twitter/contracts/authentication"
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

	globals.InitGlobals()
	mx.HandleFunc("/twitter", controllers.Login)
	mx.HandleFunc("/signup", controllers.Signup)
	mx.HandleFunc("/signupValidation", controllers.ValidateSignup("/show-users"))
	mx.HandleFunc("/loginValidation", controllers.ValidateLogin("/feed"))
	mx.HandleFunc("/*", controllers.Login)
	mx.HandleFunc("/show-users", AuthenticationMiddleware(controllers.Show_users))
	mx.HandleFunc("/follow", AuthenticationMiddleware(controllers.Follow_users(controllers.Feed)))
	mx.HandleFunc("/feed", AuthenticationMiddleware(controllers.Feed))
	mx.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	return ws, nil

}

func (w *Web) Start() error {
	return w.srv.ListenAndServe()
}

func (w *Web) Shutdown(ctx context.Context) error {
	return w.srv.Shutdown(ctx)
}

func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token,err := r.Cookie("token")
		if err==nil {
			conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := pb.NewAuthClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			reply, err := c.Authenticate(ctx, &pb.AuthenticateRequest{Token:token.Value })
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			if reply.Success{
				next.ServeHTTP(w, r)
			}else {
				w.Write([]byte("Invalid Token"))
			}
		}else {
			w.Write([]byte("Invalid Token"))
		}


	})
}