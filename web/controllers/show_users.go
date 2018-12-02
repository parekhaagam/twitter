package controllers

import (
	"github.com/parekhaagam/twitter/globals"
	"fmt"
)

//Used for "Search Users" Page
type UserFollowed struct{
	UserName string
	Isfollowed bool
}

//Used for "Search Users" Page
type UserList struct{
	List []UserFollowed
	NextPage bool
}


func Get_all_users(loggedInUserId string) (ul UserList){
	var users []UserFollowed
	allUsers := globals.AllUsers
	loggedInUser := globals.User{loggedInUserId} //should come from session
	for _,user := range allUsers {
		if user.UserName != loggedInUser.UserName {
			users = append(users, UserFollowed{user.UserName, Follows(loggedInUser, user)})
		}
	}
	fmt.Println(users)
	return UserList{users, false}
}