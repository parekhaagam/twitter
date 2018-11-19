package controllers

import (
	"github.com/parekhaagam/twitter/globals"
	"fmt"
)

//Used for "Search Users" Page
type userFollowed struct{
	UserName string
	Isfollowed bool
}

//Used for "Search Users" Page
type UserList struct{
	List []userFollowed
	NextPage bool
}


func Get_all_users(loggedInUserId string) (ul UserList){
	var users []userFollowed
	allUsers := globals.AllUsers
	loggedInUser := globals.User{loggedInUserId} //should come from session
	for _,user := range allUsers {
		if user.UserName != loggedInUser.UserName {
			users = append(users, userFollowed{user.UserName, Follows(loggedInUser, user)})
		}
	}
	fmt.Println(users)
	return UserList{users, false}
}
