package app_server

//Used for "Search Users" Page


/*func Get_all_users(loggedInUserId string) (ul globals.UserList){
	var users []globals.UserFollowed
	allUsers := memory.AllUsers
	loggedInUser := globals.User{loggedInUserId} //should come from session
	for _,user := range allUsers {
		if user.UserName != loggedInUser.UserName {
			users = append(users, globals.UserFollowed{user.UserName, Follows(loggedInUser, user)})
		}
	}
	fmt.Println(users)
	return globals.UserList{users, false}
}*/