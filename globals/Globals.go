package globals

//Global Data structure for followers
//This map will be used through out the application context
//It is initialised once when the application starts
//Initialization function: initGlobals()
//Initialization function callee: New()
var Followers map[string][]User

var AllUsers []User

//Data Object for User without Password
type User struct {
	UserName string
}

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

//Stores User Credentials
//Key: UserName
//Value: Password
var UsersRecord map[string]string
