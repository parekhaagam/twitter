package controllers

import (
	"log"
	"net/http"
	"html/template"
)

type loginPage struct {
	Email    string
	Password string
}

func Show_users(w http.ResponseWriter, r *http.Request) {
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