package app_server

import (
	"fmt"

	"github.com/parekhaagam/twitter/app_server/storage/memory"
	"github.com/parekhaagam/twitter/globals"
	"github.com/google/uuid"
	"sort"
	"strconv"
	"time"
)

func getTweets(userId string)[]globals.Tweet{
	return memory.UserTweet[userId]
}

func InsertTweets(user globals.User, content string)string {


	TID := uuid.New().String()
	fmt.Println(TID)

		for {
			if _, exists := memory.TweetIdStored[TID]; exists {
				TID = uuid.New().String()

			} else {
				break
			}
		}
	memory.TweetIdStored[TID] = TID
		tmp := globals.Tweet{
			Content:content,
			Timestamp: time.Now().Unix(),
			TID: TID,
			UserId:user.UserName,
		}

		if _, ok := memory.UserTweet[user.UserName]; ok {
			memory.UserTweet[user.UserName] = append(memory.UserTweet[user.UserName], tmp)
		}else{
			twitterTweet := []globals.Tweet{tmp}
			memory.UserTweet[user.UserName] = twitterTweet
		}
		return TID
		}



func GetFollowersTweets(followings []globals.User)[]globals.Tweet{

	var followingTweet []globals.Tweet
	for _,following := range followings{
		followingTweet = append(followingTweet , memory.UserTweet[following.UserName]...)
	}

	sort.Slice(followingTweet[:], func(i, j int) bool {
		return followingTweet[i].Timestamp > followingTweet[j].Timestamp
	})

	for index, _ := range followingTweet{
		followingTweet[index].TimeMessage = TimeToString(followingTweet[index].Timestamp)
	}
	return followingTweet
}

func TimeToString(se int64) string{
	tweetTime := time.Unix(se,0)
	now := time.Now()
	diff := now.Sub(tweetTime)
	message := ""

	seconds := int(diff.Seconds())
	if seconds >= 60{
		minutes := int(diff.Minutes())
		if minutes >= 60{
			hours := int(diff.Hours())
			if hours >= 24{
				days := hours / 24
				message = strconv.Itoa(days) + " day ago"
			}else{
				message = strconv.Itoa(hours) + " hour ago"
			}
		}else{
			message = strconv.Itoa(minutes) + " minute ago"
		}
	}else{
		message = strconv.Itoa(seconds) + " second ago"
	}

	return message
}