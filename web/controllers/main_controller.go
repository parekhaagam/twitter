package controllers

import (
	"context"
	"fmt"
	pb "github.com/parekhaagam/twitter/contracts/authentication"
	spb "github.com/parekhaagam/twitter/contracts/storage"
	"github.com/parekhaagam/twitter/globals"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
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

			conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := spb.NewStorageClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			getUsersResponse, err := c.GetAllUsers(ctx, &spb.GetUsersRequest{LoggedInUserId:loggedInUserId})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}


			var users []UserFollowed

			for _,eachUserList := range getUsersResponse.Users{
				users = append(users, UserFollowed{eachUserList.UserName, eachUserList.Isfollowed})
			}

			err = t.Execute(w, UserList{users, false})
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
			token:=GetToken(emailId)
			var tokenCookie = http.Cookie{Name: "token", Value: token}
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
			token:=GetToken(emailId)
			var tokenCookie = http.Cookie{Name: "token", Value: token}
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
			//currUser := globals.User{userIdCookie.Value}
			r.ParseForm()
			selected := r.Form["follow-chkbx"]


			conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := spb.NewStorageClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			usersList := selected[0:]
			followUserResponse,err := c.FollowUser(ctx, &spb.FollowUserRequest{UserName:userIdCookie.Value, UserNames:usersList})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}

			fmt.Println(followUserResponse.Status)

			//FollowUser(currUser, selected[0:]...)
			http.Redirect(w, r, "/feed", http.StatusSeeOther)
			//next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Missing UserId"))
		}
	})
}

func Feed(w http.ResponseWriter, httpRequest *http.Request) {
	var userIdCookie, cookieErr = httpRequest.Cookie("userId")
	if cookieErr == nil {
		loggedInUser := userIdCookie.Value
		//currUser := globals.User{loggedInUser} //should come from session @agam

		httpRequest.ParseForm()
		tweet_content := httpRequest.Form["tweet_box"]
		if len(tweet_content) > 0 {
			fmt.Println(tweet_content[0])
			//InsertTweets(currUser, tweet_content[0])

			conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := spb.NewStorageClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			insertTweetResponse, err := c.InsertTweets(ctx, &spb.InsertTweetRequest{User:&spb.User{UserName:loggedInUser}, Content:tweet_content[0]})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			fmt.Println(insertTweetResponse.TID)

		}

		conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := spb.NewStorageClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		users:= &spb.User{UserName:loggedInUser}
		getFollwerResponse, err := client.GetFollowersTweets(ctx, &spb.GetFollowersRequest{User:users})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Println(getFollwerResponse.Tweets)

		var tweets []globals.Tweet

		for _,follwerTweet := range getFollwerResponse.Tweets{
			tweets = append(
				tweets,
				globals.Tweet{
					UserId:follwerTweet.UserId,
					TID:follwerTweet.TID,
					Timestamp:follwerTweet.Timestamp,
					Content:follwerTweet.Content,
					TimeMessage:follwerTweet.TimeMessage})
		}
		followingCount := getFollwerResponse.FollowingNumber
		t, err := template.ParseFiles(WEB_HTML_DIR + "/feed.html")
		if err != nil {
			log.Print("500 Iternal Server Error", err)
		}
		c, cookieErr := httpRequest.Cookie("token")
		if cookieErr == nil {
			http.SetCookie(w, c)
			type feedObj struct {
				CurrUser        string
				FollowersNumber int
				FollowingNumber string
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

func GetToken(userid string)  string{
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.GetToken(ctx, &pb.GetTokenRequest{Userid: userid})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Token;

}