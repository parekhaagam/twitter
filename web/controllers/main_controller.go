package controllers

import (
	"container/list"
	"fmt"
	"log"
	"net/http"
	"html/template"
	"strings"
	"../../web"
	"strconv"
)

type loginPage struct {
	Email    string
	Password string
}

var userRecord  = &web.UsersRecord{
	make(map[string]string),
}


var limit = 25

func Show_users(w http.ResponseWriter, r *http.Request) {
	pageNumber,err := strconv.Atoi(strings.Split(r.URL.Path, "page=")[1])
	if err != nil{
		if (pageNumber < 1){
			pageNumber = 1
		}
		fmt.Printf("displayig content for pagge number:%d", pageNumber)
		log.Println("string to interger conversion not happen properly")
	}

	// for getting users
	users := userRecord.GetUsers(pageNumber, 25)
	// for getting numbers of pages possible
	possiblePages := userRecord.GetUsersNumber() / limit
}


func Signup(w http.ResponseWriter, r *http.Request) {


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

func Login(w http.ResponseWriter, r *http.Request) {

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


func ValidateLogin(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	emailId := strings.Join(r.Form["EmailId"], "")
	password := strings.Join(r.Form["password"], "")

	if userRecord.UserExist(emailId, password){
		fmt.Println("")
		w.Write([]byte("valid userid"))
	}else{
		w.Write([]byte("Invalid UserId"))
	}

}

func ValidateSignup(w http.ResponseWriter, r *http.Request){

	r.ParseForm()
	emailId := strings.Join(r.Form["EmailId"], "")
	password := strings.Join(r.Form["password"], "")


	if userRecord.InsertUser(emailId, password){
		w.Write([]byte("valid userid"))
	}else{
		w.Write([]byte("Invalid userid"))
	}

}