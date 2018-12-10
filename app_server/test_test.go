package app_server

import (
	//"github.com/parekhaagam/twitter/globals"
	"fmt"
	"github.com/parekhaagam/twitter/app_server/storage/memory"
	"github.com/parekhaagam/twitter/globals"

	"github.com/parekhaagam/twitter/app_server/storage"
	"testing"
)
func TestLogin(t *testing.T){

	storage.InitGlobals()
	var status = UserExist("manish.n")
	pass := memory.UsersRecord["manish.n"]
	status = status &&  pass == "admin"
	if status {
		fmt.Println("Passed : ","TestLogin")
	}else {
		t.Fatal("Error : login fails")
	}
}

func TestSignUp(t *testing.T){
	storage.InitGlobals()
	var status = InsertUser("testUser", "admin")
	pass := memory.UsersRecord["testUser"]
	status = status &&  pass == "admin"
	if status{
		fmt.Print("Passed : ","TestSignUp")
	}else{
		t.Fatal("Error : unable to create new user")
	}
}

func TestTweetPost(t *testing.T){
	storage.InitGlobals()
	currUser := globals.User{"manish.n"}
	tweet_content := "testing tweet"
	TID := InsertTweets(currUser, tweet_content)
	_, exists := memory.TweetIdStored[TID]

	tweetFound := false
	for _,tweet := range memory.UserTweet["manish.n"]{
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
	storage.InitGlobals()
	currUser := globals.User{"manish.n"}
	fmt.Println(memory.AllUsers)
	count := len(memory.AllUsers)
	list := make([]string,0)
	for _, user := range memory.AllUsers {
		if user.UserName != currUser.UserName {
			fmt.Println("Following ", user.UserName)
			list = append(list, user.UserName)
		}
	}
	FollowUser(currUser, list[0:])
	follows := memory.Followers[currUser.UserName]
	followCount := len(follows)
	fmt.Println(follows)
	if followCount == count-1 {
		fmt.Println("Passed : ", "TestFollowNewUser")
	} else {
		t.Fatal("Error in TestFollowNewUser")
	}
}

func TestFollowersTweet(t *testing.T) {
	storage.InitGlobals()
	currUser := globals.User{"manish.n"}

	selectedUserNames := []string{"dhoni007", "srk", "chandler"}
	FollowUser(currUser, selectedUserNames)

	following := memory.Followers[currUser.UserName]
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


