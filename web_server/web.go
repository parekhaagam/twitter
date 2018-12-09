package web_server

import (
	"context"
	"github.com/parekhaagam/twitter/globals"
	pb "github.com/parekhaagam/twitter/web_server/contracts/authentication"
	"github.com/parekhaagam/twitter/web_server/controllers"
	"google.golang.org/grpc"
	"log"
	"net/http"
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
var authClient pb.AuthClient


func New(cfg *Config) (*Web, error) {
	mx := http.NewServeMux()
	s := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: mx,
	}

	ws := &Web{
		srv: s,
	}

	//globals.InitGlobals()
	mx.HandleFunc("/twitter", controllers.Login)
	mx.HandleFunc("/signup", controllers.Signup)
	mx.HandleFunc("/signupValidation", controllers.ValidateSignup("/show-users"))
	mx.HandleFunc("/loginValidation", controllers.ValidateLogin("/feed"))
	mx.HandleFunc("/*", controllers.Login)
	mx.HandleFunc("/show-users", AuthenticationMiddleware(controllers.Show_users))
	mx.HandleFunc("/follow", AuthenticationMiddleware(controllers.Follow_users(controllers.Feed)))
	mx.HandleFunc("/feed", AuthenticationMiddleware(controllers.Feed))
	mx.Handle("/web_server/public/", http.StripPrefix("/web_server/public/", http.FileServer(http.Dir("web_server/public"))))
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
			authenticationClient := getAuthServerConnection()
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			reply, err := authenticationClient.Authenticate(ctx, &pb.AuthenticateRequest{Token:token.Value })
			if err != nil {
				log.Fatalf("Something went wrong ---> %v", err)
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

func getAuthServerConnection() (pb.AuthClient){
	if  authClient == nil{
		conn, err := grpc.Dial(globals.AuthServerEndpoint, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("unable to connect auth servers: %v", err)
		}
		authClient = pb.NewAuthClient(conn)
	}
	return authClient
}