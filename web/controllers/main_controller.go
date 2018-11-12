package controllers

import (
	"../../globals"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const WEB_HTML_DIR  = "web/html"

type loginPage struct {
	Email    string
	Password string
}


var limit = 25

func Signup(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(WEB_HTML_DIR+"/signup.html")
	if err != nil{
		log.Print("Sign up page not loaded properly", err)
	}
	mLoginPage := loginPage{
		"EmailId",
		"password",
	}
	err = t.Execute(w, mLoginPage)
	if err != nil {
		log.Print("error while executing ", err)
	}
}


func Show_users(w http.ResponseWriter, r *http.Request) {
	//pageNumber,err := strconv.Atoi(strings.Split(r.URL.Path, "page=")[1])
	//if err != nil{
	//	if (pageNumber < 1){
	//		pageNumber = 1
	//	}
	//	fmt.Printf("displayig content for pagge number:%d", pageNumber)
	//	log.Println("string to interger conversion not happen properly")
	//}
	//
	//// for getting users
	////users := userRecord.GetUsers(pageNumber, 25)

	t, err := template.ParseFiles(WEB_HTML_DIR+"/users_to_follow.html")
	if err != nil{
		log.Print("500 Iternal Server Error", err)
	}
	err = t.Execute(w, Get_all_users())
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

	if UserExist(emailId, password){
		fmt.Println("")
		w.Write([]byte("valid userid"))
	}else{
		w.Write([]byte("Invalid UserId"))
	}

}

func ValidateSignup(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	emailId := strings.Join(r.Form["EmailId"], "")
	password := strings.Join(r.Form["password"], "")

	if InsertUser(emailId, password) {
		w.Write([]byte("valid userid"))
	} else {
		w.Write([]byte("Invalid userid"))
	}

}

func Follow_users(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	selected := r.Form["follow-chkbx"]

	followers := globals.Followers
	currUserName := "manish.n" //should come from session @agam
	follows,ok := followers[currUserName]
	if !ok {
		follows = []globals.User{}
	}

	for _, userName := range selected {
		follows = append(follows, globals.User{userName})
	}

	followers[currUserName] = follows
	fmt.Println(followers)
	Show_users(w, r)
}