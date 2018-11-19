package controllers

import (
	"github.com/parekhaagam/twitter/globals"
	"github.com/parekhaagam/twitter/web/auth"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const WEB_HTML_DIR = "web/html"

type loginPage struct {
	Email    string
	Password string
}

var limit = 25

func Signup(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(WEB_HTML_DIR + "/signup.html")
	if err != nil {
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
	t, err := template.ParseFiles(WEB_HTML_DIR + "/users_to_follow.html")
	if err != nil {
		log.Print("500 Iternal Server Error", err)
	}
	var tokenCookie, cookieErr = r.Cookie("token")
	if cookieErr == nil {
		var userIdCookie, cookieErr = r.Cookie("userId")
		if cookieErr == nil {
			loggedInUserId := userIdCookie.Value
			http.SetCookie(w, tokenCookie);
			err = t.Execute(w, Get_all_users(loggedInUserId))
			if err != nil {
				log.Print("error while executing ", err)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Missing UserId"))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Missing Token"))
	}

}

func Login(w http.ResponseWriter, r *http.Request) {

	t:= template.Must(template.ParseFiles(WEB_HTML_DIR+"/login.html"))
	mLoginPage := loginPage{
		Email:    "EmailId",
		Password: "password",
	}
	err := t.Execute(w, mLoginPage)
	if err != nil {
		log.Print("error while executing ", err)
	}

}

func ValidateLogin(destURL string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		emailId := strings.Join(r.Form["EmailId"], "")
		password := strings.Join(r.Form["password"], "")

		if UserExist(emailId, password) {
			fmt.Println("")
			r.Header.Set("Cookie", "")
			var tokenCookie = http.Cookie{Name: "token", Value: auth.GetToken(emailId)}
			var userIdCookie = http.Cookie{Name: "userId", Value: emailId}
			http.SetCookie(w, &tokenCookie)
			http.SetCookie(w, &userIdCookie)
			r.AddCookie(&tokenCookie)
			r.AddCookie(&userIdCookie)
			http.Redirect(w, r, destURL, http.StatusSeeOther)
		} else {
			w.Write([]byte("Invalid UserId"))
		}
	})
}

func ValidateSignup(destURL string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		emailId := strings.Join(r.Form["EmailId"], "")
		password := strings.Join(r.Form["password"], "")

		if InsertUser(emailId, password) {
			r.Header.Set("Cookie", "")
			var tokenCookie = http.Cookie{Name: "token", Value: auth.GetToken(emailId)}
			var userIdCookie = http.Cookie{Name: "userId", Value: emailId}
			http.SetCookie(w, &tokenCookie)
			http.SetCookie(w, &userIdCookie)
			r.AddCookie(&tokenCookie)
			r.AddCookie(&userIdCookie)
			http.Redirect(w, r, destURL, http.StatusSeeOther)
		} else {
			w.Write([]byte("Invalid userid"))
		}

	})
}

func Follow_users(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userIdCookie, cookieErr = r.Cookie("userId")
		if cookieErr == nil {
			currUser := globals.User{userIdCookie.Value}
			r.ParseForm()
			selected := r.Form["follow-chkbx"]
			FollowUser(currUser, selected[0:]...)
			http.Redirect(w, r, "/feed", http.StatusSeeOther)
			//next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Missing UserId"))
		}
	})

}

func Feed(w http.ResponseWriter, r *http.Request) {
	var userIdCookie, cookieErr = r.Cookie("userId")
	if cookieErr == nil {
		loggedInUser := userIdCookie.Value
		currUser := globals.User{loggedInUser} //should come from session @agam

		r.ParseForm()
		tweet_content := r.Form["tweet_box"]
		if len(tweet_content) > 0 {
			fmt.Println(tweet_content[0])
			InsertTweets(currUser, tweet_content[0])
		}

		following := GetAllFollowing(currUser)
		followingCount := len(following)
		following = append(following, currUser)
		tweets := GetFollowersTweets(following)
		fmt.Println(tweets)

		t, err := template.ParseFiles(WEB_HTML_DIR + "/feed.html")
		if err != nil {
			log.Print("500 Iternal Server Error", err)
		}
		c, cookieErr := r.Cookie("token")
		if cookieErr == nil {
			http.SetCookie(w, c)
			type feedObj struct {
				CurrUser        string
				FollowersNumber int
				FollowingNumber int
				Tweets          []globals.Tweet
			}

			feedsObj := feedObj{CurrUser: loggedInUser, FollowingNumber: followingCount, Tweets: tweets}
			err = t.Execute(w, feedsObj)
			if err != nil {
				log.Print("error while executing ", err)
			}
		} else {
			w.WriteHeader(500)
			//w.Write([]byte("Error"))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Missing UserId"))
	}
}
