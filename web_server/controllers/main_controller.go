package controllers

import (
	"context"
	"fmt"
	spb "github.com/parekhaagam/twitter/app_server/contract"
	"github.com/parekhaagam/twitter/globals"
	pb "github.com/parekhaagam/twitter/web_server/contracts/authentication"
	"google.golang.org/grpc"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var authClient pb.AuthClient
type loginPage struct {
	Email    string
	Password string
}

var limit = 25
var storageClient spb.StorageClient
func Signup(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(WebHTMLDir + "/signup.html")
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
	t, err := template.ParseFiles(WebHTMLDir + "/users_to_follow.html")
	if err != nil {
		log.Print("500 Iternal Server Error", err)
	}
	var tokenCookie, cookieErr = r.Cookie("token")
	if cookieErr == nil {
		var userIdCookie, cookieErr = r.Cookie("userId")
		if cookieErr == nil {
			loggedInUserId := userIdCookie.Value
			http.SetCookie(w, tokenCookie);

			tweetClient := getStorageClient()
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			getUsersResponse, err := tweetClient.GetAllUsers(ctx, &spb.GetUsersRequest{LoggedInUserId:loggedInUserId})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}


			var users []globals.UserFollowed

			for _,eachUserList := range getUsersResponse.Users{
				users = append(users, globals.UserFollowed{eachUserList.UserName, eachUserList.Isfollowed})
			}

			err = t.Execute(w, globals.UserList{users, false})
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
	files, err1 := ioutil.ReadDir("./")
	if err1 != nil {
		log.Fatal(err1)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	t:= template.Must(template.ParseFiles(WebHTMLDir+"/login.html"))
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
		//password := strings.Join(r.Form["password"], "")
		isValidUser :=userExist(emailId)
		if isValidUser {
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
		isNewUserInserted := insertUser(emailId,password)
		if isNewUserInserted {
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


			tweetClient := getStorageClient()
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			usersList := selected[0:]
			followUserResponse,err := tweetClient.FollowUser(ctx, &spb.FollowUserRequest{UserName:userIdCookie.Value, UserNames:usersList})
			if err != nil {
				log.Fatalf("Something went wrong --> %v", err)
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

			tweetClient := getStorageClient()
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			insertTweetResponse, err := tweetClient.InsertTweets(ctx, &spb.InsertTweetRequest{User:&spb.User{UserName:loggedInUser}, Content:tweet_content[0]})
			if err != nil {
				log.Fatalf("Something went wrong --> %v", err)
			}
			fmt.Println(insertTweetResponse.TID)

		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		tweetClient := getStorageClient()
		users:= &spb.User{UserName:loggedInUser}
		getFollwerResponse, err := tweetClient.GetFollowersTweets(ctx, &spb.GetFollowersRequest{User:users})
		if err != nil {
			log.Fatalf("Something went wrong --> %v", err)
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
		t, err := template.ParseFiles(WebHTMLDir + "/feed.html")
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
	authenticationClient := getAuthServerConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := authenticationClient.GetToken(ctx, &pb.GetTokenRequest{Userid: userid})
	if err != nil {
		log.Fatalf("Something went wrong --> %v", err)
	}
	return r.Token;

}
func getStorageClient() (spb.StorageClient){
	if storageClient == nil {
		conn, err := grpc.Dial(globals.StorageServerEndpoint, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		storageClient = spb.NewStorageClient(conn)
	}
	return storageClient
}
func getAuthServerConnection() (pb.AuthClient){
	if  authClient == nil{
		conn, err := grpc.Dial(globals.AuthServerEndpoint, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		authClient = pb.NewAuthClient(conn)
	}
	return authClient
}
func insertUser(userId string,password string) (bool){
	userDBClient := getStorageClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, _ := userDBClient.InsertUser(ctx, &spb.InsertUserRequest{UserName: userId,Password:password})
	return r.Success
}
func userExist(userId string) (bool){
	userDBClient := getStorageClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := userDBClient.UserExist(ctx, &spb.UserExistRequest{UserName: userId})
	if err!=nil{
		fmt.Println(err)
		return false
	}
	return r.Success
}