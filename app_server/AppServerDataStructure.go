package app_server

import (
	"github.com/parekhaagam/twitter/globals"
	"sync"
)

var Followers map[string][]globals.User

var AllUsers []globals.User

var UsersRecord map[string]string
var UserRecordLock sync.Mutex

// Stores User tweets
// Key: UserName
//Value : Tweet Object
var UserTweet map[string][]globals.Tweet

var TweetIdStored map[string]string
