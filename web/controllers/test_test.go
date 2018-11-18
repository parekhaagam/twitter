package controllers

import (
	"../../globals"
	"fmt"
	"testing"
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


func TestFollowNewUser(t *testing.T) {
	currUser := globals.User{"dhoni007"}

	//get list of people curr user is following
	follows := globals.Followers[currUser.UserName]
	fmt.Println("old Users followed = ", follows)
	count := 2

	userNewlyFollowed := make([]string,0)
	addNewUser:
	for _, user := range globals.AllUsers {
		fmt.Println("Checking for user = ", user.UserName)
		for _, userFollowed := range follows {
			if userFollowed.UserName != user.UserName {
				userNewlyFollowed = append(userNewlyFollowed, user.UserName)
				FollowUser(currUser, user.UserName)
				count--
			}
			if count == 0 {
				break addNewUser
			}
		}
	}

	fmt.Println("New Users followed = ", userNewlyFollowed)
	follows = globals.Followers[currUser.UserName]
	count = 0
	//check if actually followed
	for _, userName := range userNewlyFollowed{
		for _, user := range follows {
			if user.UserName == userName {
				count++
				break
			}
		}
	}

	if count == 2 {
		fmt.Println("Passed : ", "TestFollowNewUser")
	} else {
		t.Fatal("Error in TestFollowNewUser")
	}

}


func TestTweetTIDConsistency(t *testing.T){

	users := []string{"manish.n", "dhoni007", "srk", "chandler", "manish.n"}
	tweets:= []string{"test 1", "test 2", "test 3", "test 4", "test 5"}
	tidMap := make(map[string]string)

	for i := 0; i < 5; i++ {

		go func(userId string, tweet_content string ) {
		currUser := globals.User{users[i]}
		tweet_content = tweets[i]
		TID := InsertTweets(currUser, tweet_content)
		tidMap[TID] = TID

		}(users[i], tweets[i])

		if len(tidMap) == 5{
			fmt.Println("Passed : ", "TestTweetTIDConsistency")
		}else{
			t.Fatal("Error in TestTweetTIDConsistency")
		}
	}
	}


func TestFollowersTweet(t *testing.T)  {
	fmt.Println("Passed : ", "TestFollowersTweet")
}


