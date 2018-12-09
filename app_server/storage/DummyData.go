package storage

import (
	"github.com/google/uuid"
	"github.com/parekhaagam/twitter/globals"
	"time"
)

func InitGlobals() {
	Followers = make(map[string][]globals.User)
	UsersRecord = make(map[string]string)
	UserTweet = make(map[string][]globals.Tweet)
	TweetIdStored = make(map[string]string)
	AllUsers = insertDummies()
}

func insertDummies() (allUsers[] globals.User){
	UsersRecord["manish.n"] = "admin"
	UsersRecord["dhoni007"] = "admin"
	UsersRecord["srk"] = "admin"
	UsersRecord["chandler"] = "admin"

	allUsers = append(allUsers, globals.User{"manish.n"})
	allUsers = append(allUsers, globals.User{"dhoni007"})
	allUsers = append(allUsers, globals.User{"srk"})
	allUsers = append(allUsers, globals.User{"chandler"})


	tweet1 := globals.Tweet{TID:uuid.New().String(), Content:"Mujhe bhi T20 khelna h", Timestamp:time.Now().Unix(), UserId:"dhoni007"}
	tweet2 := globals.Tweet{TID:uuid.New().String(), Content:"Zero releasing December 2018", Timestamp:time.Now().Unix(), UserId:"srk"}
	tweet3 := globals.Tweet{TID:uuid.New().String(), Content:"Could I be wearing anymore clothes", Timestamp:time.Now().Unix(), UserId:"chandler"}
	tweet4 := globals.Tweet{TID:uuid.New().String(), Content:"SDE at Google", Timestamp:time.Now().Unix(), UserId:"manish.n"}
	tweet5 := globals.Tweet{TID:uuid.New().String(), Content:"Virat ne mujhe nikal diya T20 team se", Timestamp:time.Now().Unix(), UserId:"dhoni007"}

	UserTweet["dhoni007"] = append(UserTweet["dhoni007"], tweet1)
	UserTweet["dhoni007"] = append(UserTweet["dhoni007"], tweet5)
	UserTweet["manish.n"] = append(UserTweet["manish.n"], tweet4)
	UserTweet["chandler"] = append(UserTweet["chandler"], tweet3)
	UserTweet["srk"] = append(UserTweet["srk"], tweet2)

	return allUsers
}
