package app_server

import (
	"fmt"
	"github.com/parekhaagam/twitter/globals"
)

func InitGlobals() {
	/*Followers = make(map[string][]globals.User)
	UsersRecord = make(map[string]string)
	UserTweet = make(map[string][]globals.Tweet)*/
	//TweetIdStored = make(map[string]string)
	insertDummies()
}

func insertDummies() {

	//UsersRecord["manish.n"] = "admin"
	//intentionally not handled the error here, since this is dummy data
	status,_ := InsertUserRecord("manish.n", "admin")
	fmt.Println("status:", status)

	//UsersRecord["dhoni007"] = "admin"
	status,_ = InsertUserRecord("dhoni007", "admin")
	fmt.Println("status:", status)

	//UsersRecord["srk"] = "admin"
	status,_ = InsertUserRecord("srk", "admin")
	fmt.Println("status:", status)

	//UsersRecord["chandler"] = "admin"
	status,_ = InsertUserRecord("chandler", "admin")
	fmt.Println("status:", status)

	/*
	allUsers = append(allUsers, globals.User{"manish.n"})
	allUsers = append(allUsers, globals.User{"dhoni007"})
	allUsers = append(allUsers, globals.User{"srk"})
	allUsers = append(allUsers, globals.User{"chandler"})

	tweet1 := globals.Tweet{TID: uuid.New().String(), Content: "Mujhe bhi T20 khelna h", Timestamp: time.Now().Unix(), UserId: "dhoni007"}
	tweet2 := globals.Tweet{TID: uuid.New().String(), Content: "Zero releasing December 2018", Timestamp: time.Now().Unix(), UserId: "srk"}
	tweet3 := globals.Tweet{TID: uuid.New().String(), Content: "Could I be wearing anymore clothes", Timestamp: time.Now().Unix(), UserId: "chandler"}
	tweet4 := globals.Tweet{TID: uuid.New().String(), Content: "SDE at Google", Timestamp: time.Now().Unix(), UserId: "manish.n"}
	tweet5 := globals.Tweet{TID: uuid.New().String(), Content: "Virat ne mujhe nikal diya T20 team se", Timestamp: time.Now().Unix(), UserId: "dhoni007"}
*/
	tid,_ := StorageInsertTweets(globals.User{"dhoni007"}, "Mujhe bhi T20 khelna h")
	fmt.Println("tid", tid)
	tid,_ = StorageInsertTweets(globals.User{"dhoni007"}, "Virat ne mujhe nikal diya T20 team se")
	fmt.Println("tid", tid)
	tid,_ = StorageInsertTweets(globals.User{"manish.n"}, "SDE at Google")
	fmt.Println("tid", tid)
	tid,_ = StorageInsertTweets(globals.User{"chandler"}, "Could I be wearing anymore clothes")
	fmt.Println("tid", tid)
	tid,_ = StorageInsertTweets(globals.User{"srk"}, "Zero releasing December 2018")
	fmt.Println("tid", tid)

	/*
	UserTweet["dhoni007"] = append(UserTweet["dhoni007"], tweet1)
	UserTweet["dhoni007"] = append(UserTweet["dhoni007"], tweet5)
	UserTweet["manish.n"] = append(UserTweet["manish.n"], tweet4)
	UserTweet["chandler"] = append(UserTweet["chandler"], tweet3)
	UserTweet["srk"] = append(UserTweet["srk"], tweet2)


	status := InsertUserRecord("agamTesting", "admin")
	fmt.Println("insert status:" , status)
	status = CheckUserRecord("agamTesting", "admin")
	fmt.Println("check status:" , status)


	status = InsertUserRecord("yashTesting4", "admin")
	status = InsertUserRecord("manishTesting4", "admin")

	fmt.Println("\n\ntesting followerUser")
	StorageFollowUser(globals.User{"agamTesting"}, []string{"yashTesting"})
	usersFollow := GetAllFollowingUser(globals.User{"agamTesting"})
	if usersFollow[0].UserName == "yashTesting"{
		fmt.Println("followUser passed")
	}else{
		fmt.Println("followUser failed")
	}

	fmt.Println("\n\ntesting tweet insert")
	tid := StorageInsertTweets(globals.User{"agamTesting"}, "Mujhe bhi T20 khelna h")
	fmt.Println("tid", tid)
	tweetsList := StorageGetFollowersTweets([]globals.User{globals.User{"agamTesting"}})
	if tweetsList[0].Content =="Mujhe bhi T20 khelna h"{
		fmt.Println("insert tweet passed")
	}else{
		fmt.Println("insert tweet failed")
	}

	fmt.Println("\n\ntesting getting all users")
	allUsersList := Storage_Get_all_users("agamTesting")
	if  len(allUsersList.List)== 2{
		fmt.Println("passed")
	} else{
		fmt.Println("failed" )
		for _, name := range allUsersList.List {
			fmt.Println(name.UserName)
		}

	}*/
}
