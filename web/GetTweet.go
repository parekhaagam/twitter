package web

import (
	"../globals"
	"github.com/google/uuid"
	"sort"
	"time"
)

func getTweets(userId string)[]globals.Tweet{
	return globals.UserTweet[userId]
}

func insertTweets(userId string, content string){

	if _ ,ok := globals.UserTweet[userId]; ok{


		TID := uuid.New().String()
		for {
			if _, exists := globals.TweetIdStored[TID]; exists {
				TID = uuid.New().String()
			} else {
				break
			}
		}

		tmp := globals.Tweet{
			Content:content,
			Timestamp: time.Now().Unix(),
			TID: TID,
			UserId:userId,
		}
		globals.UserTweet[userId] = append(globals.UserTweet[userId], tmp)
	}
}


func getFollowersTweets(followings []globals.User)[]globals.Tweet{

	var followingTweet []globals.Tweet
	for _,following := range followings{
		followingTweet = append(followingTweet , globals.UserTweet[following.UserName]...)
	}

	sort.Slice(followingTweet[:], func(i, j int) bool {
		return followingTweet[i].Timestamp > followingTweet[j].Timestamp
	})

	return followingTweet
}