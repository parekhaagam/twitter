package globals

//Global Data structure for followers
//This map will be used through out the application context
//It is initialised once when the application starts
//Initialization function: initGlobals()
//Initialization function callee: New()

//Data Object for User without Password
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
	NextPage bool `json:"NextPage"`
}
