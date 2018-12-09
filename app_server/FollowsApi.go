package app_server

import (
	"github.com/parekhaagam/twitter/app_server/storage"
	"github.com/parekhaagam/twitter/globals"
)
func FollowUser(follower globals.User, selectedUserNames []string) {
	follows := make([]globals.User, 0)

	for _, userName := range selectedUserNames {
		follows = append(follows, globals.User{userName})
	}

	allFollowers := storage.Followers
	allFollowers[follower.UserName] = follows
}