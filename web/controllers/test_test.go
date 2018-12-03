package controllers

import (
	"github.com/parekhaagam/twitter/globals"
	"fmt"
	"testing"
)
func TestLogin(t *testing.T){
	globals.InitGlobals()
	var status = UserExist("manish.n")
	pass := globals.UsersRecord["manish.n"]
	status = status &&  pass == "admin"
	if status {
		fmt.Println("Passed : ","TestLogin")
	}else {
		t.Fatal("Error : login fails")
	}
}

func TestSignUp(t *testing.T){
	globals.InitGlobals()
	var status = InsertUser("testUser", "admin")
	pass := globals.UsersRecord["testUser"]
	status = status &&  pass == "admin"
	if status{
		fmt.Print("Passed : ","TestSignUp")
	}else{
		t.Fatal("Error : unable to create new user")
	}
}

func TestTweetPost(t *testing.T){
	globals.InitGlobals()
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

	print(exists)
	if exists && tweetFound{
		fmt.Println("Passed : ", "TestTweetPost")
	}else{
		t.Fatal("Error in testTweetPost")
	}
}


func TestFollowAllUser(t *testing.T) {
	globals.InitGlobals()
	currUser := globals.User{"manish.n"}
	fmt.Println(globals.AllUsers)
	count := len(globals.AllUsers)
	list := make([]string,0)
	for _, user := range globals.AllUsers {
		if user.UserName != currUser.UserName {
			fmt.Println("Following ", user.UserName)
			list = append(list, user.UserName)
		}
	}
	FollowUser(currUser, list[0:])
	follows := globals.Followers[currUser.UserName]
	followCount := len(follows)
	fmt.Println(follows)
	if followCount == count-1 {
		fmt.Println("Passed : ", "TestFollowNewUser")
	} else {
		t.Fatal("Error in TestFollowNewUser")
	}
}

func TestFollowersTweet(t *testing.T) {
	globals.InitGlobals()
	currUser := globals.User{"manish.n"}

	selectedUserNames := []string{"dhoni007", "srk", "chandler"}
	FollowUser(currUser, selectedUserNames)

	following := globals.Followers[currUser.UserName]
	tweets := GetFollowersTweets(following)
	for _, tweet := range tweets {
		userFound := false
		for _, user := range following {
			if tweet.UserId == user.UserName {
				userFound = true
				break
			}
		}
		if !userFound {
			t.Fatal("Fail: TestFollowersTweet")
			break
		}
	}
}


