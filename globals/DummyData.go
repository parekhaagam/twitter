package globals

import (
	"github.com/google/uuid"
	"time"
)

func InitGlobals() {
	Followers = make(map[string][]User)
	UsersRecord = make(map[string]string)
	UserTweet = make(map[string][]Tweet)
	TweetIdStored = make(map[string]string)
	AllUsers = insertDummies()
}

func insertDummies() (allUsers[] User){
	UsersRecord["manish.n"] = "admin"
	UsersRecord["dhoni007"] = "admin"
	UsersRecord["srk"] = "admin"
	UsersRecord["chandler"] = "admin"

	allUsers = append(allUsers, User{"manish.n"})
	allUsers = append(allUsers, User{"dhoni007"})
	allUsers = append(allUsers, User{"srk"})
	allUsers = append(allUsers, User{"chandler"})


	tweet1 := Tweet{TID:uuid.New().String(), Content:"Mujhe bhi T20 khelna h", Timestamp:time.Now().Unix(), UserId:"dhoni007"}
	tweet2 := Tweet{TID:uuid.New().String(), Content:"Zero releasing December 2018", Timestamp:time.Now().Unix(), UserId:"srk"}
	tweet3 := Tweet{TID:uuid.New().String(), Content:"Could I be wearing anymore clothes", Timestamp:time.Now().Unix(), UserId:"chandler"}
	tweet4 := Tweet{TID:uuid.New().String(), Content:"SDE at Google", Timestamp:time.Now().Unix(), UserId:"manish.n"}
	tweet5 := Tweet{TID:uuid.New().String(), Content:"Virat ne mujhe nikal diya T20 team se", Timestamp:time.Now().Unix(), UserId:"dhoni007"}

	UserTweet["dhoni007"] = append(UserTweet["dhoni007"], tweet1)
	UserTweet["dhoni007"] = append(UserTweet["dhoni007"], tweet5)
	UserTweet["manish.n"] = append(UserTweet["manish.n"], tweet4)
	UserTweet["chandler"] = append(UserTweet["chandler"], tweet3)
	UserTweet["srk"] = append(UserTweet["srk"], tweet2)

	return allUsers
}
