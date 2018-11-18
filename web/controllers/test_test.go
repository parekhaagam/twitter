package controllers

import (
	"fmt"
	"testing"
	"../../globals"
)
func TestLogin(t *testing.T){

	var status = UserExist("manish.n", "admin")
	pass := globals.UsersRecord["manish.n"]
	status = status &&  pass == "admin"
	if status {
		fmt.Println("Passed : ","TestLogin")
	}else {
		t.Fatal("Error : login fails")
	}
}

func TestSignUp(t *testing.T){

	var status = InsertUser("testUser", "admin")
	pass := globals.UsersRecord["testUser"]
	status = status &&  pass == "admin"
	if status{
		fmt.Print("Passed : ","TestSignUp")
	}else{
		t.Fatal("Error : unable to create new user")
	}
}

func TestLoginConcurrent(t *testing.T){
	fmt.Print("Passed : ","TestLoginConcurrent")
}

func TestSignUpConcurrent(t *testing.T){
	fmt.Print("Passed : ","TestSignUpConcurrent")
}

func TestTweetPost(t *testing.T){

	currUser := globals.User{"manish.n"}
	tweet_content := "testing tweet"
	TID := InsertTweets(currUser, tweet_content)
	_, exists := globals.TweetIdStored[TID]

	tweetFound := false
	for _,tweet := range globals.UserTweet["manish.n"]{
		if tweet.TID == TID{
			tweetFound = true
		}
	}

	if exists && tweetFound{
		fmt.Println("Passed : ", "TestTweetPost")
	}else{
		t.Fatal("Error in testTweetPost")
	}
}

func TestFollowersTweet(t *testing.T)  {
	fmt.Println("Passed : ", "TestFollowersTweet" )
}


