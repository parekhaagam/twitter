package globals

//Global Data structure for followers
//This map will be used through out the application context
//It is initialised once when the application starts
//Initialization function: initGlobals()
//Initialization function callee: New()

//Data Object for User without Password
/*type User struct {
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
}*/


type User struct {
	UserName string `json:"UserName"`
}

type Tweet struct {
	UserId      string `json:"UserId"`
	TID         string `json:"TID"`
	Timestamp   int64	`json:"Timestamp"`
	Content     string  `json:"Content"`
	TimeMessage string  `json:"TimeMessage"`
}


type UserFollowed struct{
	UserName string `json:"UserName"`
	Isfollowed bool  `json:"isfollowed"`
}

//Used for "Search Users" Page
type UserList struct{
	List []UserFollowed `json:"UserFollowed"`
	NextPage bool `json:"NextPage"`  //to show user profile page
}
