package globals

//Global Data structure for followers
//This map will be used through out the application context
//It is initialised once when the application starts
//Initialization function: initGlobals()
//Initialization function callee: New()
var Followers map[string][]User

var AllUsers []User

type User struct {
	Name string
}

