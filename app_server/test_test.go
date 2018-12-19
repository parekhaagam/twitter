package app_server

import (
	//"github.com/parekhaagam/twitter/globals"
	"fmt"
	"github.com/parekhaagam/twitter/globals"
	"testing"
)
func TestLogin(t *testing.T){

	InitGlobals()
	status := CheckUserRecord("manish.n", "admin")

	if status {
		fmt.Println("Passed : ","TestLogin")
	}else {
		t.Fatal("Error : login fails")
	}
}

func TestSignUp(t *testing.T){

	InitGlobals()
	status := InsertUserRecord("agamTesting", "admin")
	fmt.Println("insert status:" , status)
	status = CheckUserRecord("agamTesting", "admin")
	fmt.Println("check status:" , status)

	if status {
		fmt.Println("Passed : ","TestSignUp")
	}else {
		t.Fatal("Error : SignUp fails")
	}
}

func TestTweetPost(t *testing.T){

	InitGlobals()
	tid := StorageInsertTweets(globals.User{"manish.n"}, "Testing tweet")
	fmt.Println("tid", tid)
	tweetsList := StorageGetFollowersTweets([]globals.User{globals.User{"manish.n"}})
	if tweetsList[0].Content =="Testing tweet"{
		fmt.Println("insert tweet passed")
	}else{
		fmt.Println("insert tweet failed")
	}
}


func TestFollowAllUser(t *testing.T) {

	InitGlobals()

	StorageFollowUser(globals.User{"manish.n"}, []string{"dhoni007"})
	usersFollow := GetAllFollowingUser(globals.User{"manish.n"})
	if usersFollow[0].UserName == "dhoni007"{
		fmt.Println("followUser passed")
	}else{
		fmt.Println("followUser failed")
	}

}

func TestFollowersTweet(t *testing.T) {
	InitGlobals()
	currUser := globals.User{"manish.n"}

	selectedUserNames := []string{"dhoni007", "srk", "chandler"}
	StorageFollowUser(currUser, selectedUserNames)

	following := GetAllFollowingUser(currUser)
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


