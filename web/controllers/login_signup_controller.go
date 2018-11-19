package controllers

import (
	"../../globals"
)

func UserExist(userName string, password string)bool{

	globals.UserRecordLock.Lock()
	defer globals.UserRecordLock.Unlock()

	_, exists := globals.UsersRecord[userName]
	if exists {
		return true
	}else {
		return false
	}

}

func InsertUser(newUserName string, password string)bool{

	//globals.UserRecordLock.Lock()
	//defer globals.UserRecordLock.Unlock()

	if ! UserExist(newUserName, password) {
		globals.UsersRecord[newUserName] = password
		globals.AllUsers = append(globals.AllUsers, globals.User{newUserName})
		return true
	}else {
		return false
	}
}

//var usersDataCache []globals.UserFollowed = nil

//Returns true if user 1 follows user 2
func Follows(user1 globals.User, user2 globals.User) bool{
	followers := globals.Followers
	follows,ok := followers[user1.UserName]

	doesFollow := false
	if ok {
		for _,followedUser := range follows {
			if followedUser.UserName == user2.UserName {
				doesFollow = true
				break
			}
		}
	}
	return doesFollow
}

func GetAllFollowing(user globals.User) ([]globals.User){
	followers := globals.Followers
	follows,ok := followers[user.UserName]
	if !ok {
		follows = make([]globals.User,0,0)
	}
	return follows
}

/*func getTweets(emailId string)[]{

}*/

