package web

import (
	"./controllers"
	"context"
	"net/http"
)

type Web struct {
	srv *http.Server
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

	mx.HandleFunc("/", controllers.Login)
	mx.HandleFunc("/signup", controllers.Signup)
	mx.HandleFunc("/signupValidation", controllers.ValidateSignup)
	mx.HandleFunc("/loginValidation", controllers.ValidateLogin)
	mx.HandleFunc("/*", controllers.Signup)
	mx.HandleFunc("/show-users", controllers.Show_users)
	return ws, nil
}

func (w *Web) Start() error {
	return w.srv.ListenAndServe()
}

func (w *Web) Shutdown(ctx context.Context) error {
	return w.srv.Shutdown(ctx)
}