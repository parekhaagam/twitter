package globals

//Global Data structure for followers
//This map will be used through out the application context
//It is initialised once when the application starts
//Initialization function: initGlobals()
//Initialization function callee: New()

//Data Object for User without Password
type User struct {
	UserName string
}

type Tweet struct {
	UserId      string
	TID         string
	Timestamp   int64
	Content     string
	TimeMessage string
}


type UserFollowed struct{
	UserName string
	Isfollowed bool
}

//Used for "Search Users" Page
type UserList struct{
	List []UserFollowed
	NextPage bool
}
