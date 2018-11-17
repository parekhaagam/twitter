package controllers

import (
	"../../globals"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"../auth"
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
	var c,cookieErr = r.Cookie("token")
	if cookieErr==nil{
		http.SetCookie(w,c);
		err = t.Execute(w, Get_all_users())
		if err != nil {
			log.Print("error while executing ", err)
		}
	}else {
		w.WriteHeader(500)
		//w.Write([]byte("Error"))
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


func ValidateLogin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		emailId := strings.Join(r.Form["EmailId"], "")
		password := strings.Join(r.Form["password"], "")

		if UserExist(emailId, password){
			fmt.Println("")
			//w.Write([]byte("valid userid"))
			var c =http.Cookie{Name:"token",Value:auth.GetToken(emailId)}
			r.Header.Set("Cookie","")
			http.SetCookie(w,&c)
			r.AddCookie(&c)
			next.ServeHTTP(w,r)
		}else{
			w.Write([]byte("Invalid UserId"))
		}
	})

}

func ValidateSignup(next http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		emailId := strings.Join(r.Form["EmailId"], "")
		password := strings.Join(r.Form["password"], "")

		if InsertUser(emailId, password) {
			var c =http.Cookie{Name:"token",Value:auth.GetToken(emailId)}
			http.SetCookie(w,&c)
			r.Header.Set("Cookie","")
			r.AddCookie(&c)
			//w.Write([]byte("valid userid"))
			next.ServeHTTP(w,r)
		} else {
			w.Write([]byte("Invalid userid"))
		}

	})
}

func Follow_users(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	selected := r.Form["follow-chkbx"]

	followers := globals.Followers
	currUser := globals.User{"manish.n"} //should come from session @agam
	follows := GetAllFollowing(currUser)

	unfollowed := getMissing(follows, selected)
	for user, index := range unfollowed {
		fmt.Println("Unfollowed", user.UserName)
		follows = append(follows[:index],follows[index+1:]...)
	}

	for _, userName := range selected {
		follows = append(follows, globals.User{userName})
	}

	followers[currUser.UserName] = follows
	fmt.Println(followers)
	Feed(w, r)
}

func getMissing(follows []globals.User, selected []string) (map[globals.User]int){
	map1 := make(map[string]int)
	unfollowed := make(map[globals.User]int)
	for _, userName := range selected {
		map1[userName] = 1
	}

	for i, user := range follows {
		_, ok := map1[user.UserName]
		if !ok {
			unfollowed[user] = i
		}
	}
	return unfollowed
}

func Feed(w http.ResponseWriter, r *http.Request) {
	currUser := globals.User{"manish.n"} //should come from session @agam
	following := GetAllFollowing(currUser)
	tweets := GetFollowersTweets(following)
	t, err := template.ParseFiles(WEB_HTML_DIR+"/feed.html")
	if err != nil{
		log.Print("500 Iternal Server Error", err)
	}
	var c,cookieErr = r.Cookie("token")
	if cookieErr==nil {
		http.SetCookie(w,c)
		type feedObj struct {
			CurrUser string
			FollowersNumber int
			FollowingNumber int
			Tweets []globals.Tweet
		}

		feedsObj := feedObj{CurrUser:"manish.n", FollowersNumber:10, FollowingNumber:10, Tweets:tweets}
		err = t.Execute(w, feedsObj)
		if err != nil {
			log.Print("error while executing ", err)
		}
	}else {
		w.WriteHeader(500)
		//w.Write([]byte("Error"))
	}
}