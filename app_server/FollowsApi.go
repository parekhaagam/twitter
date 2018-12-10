package app_server

import (
	"github.com/parekhaagam/twitter/app_server/storage/memory"
	"github.com/parekhaagam/twitter/globals"
)
func FollowUser(follower globals.User, selectedUserNames []string) {
	follows := make([]globals.User, 0)

	for _, userName := range selectedUserNames {
		follows = append(follows, globals.User{userName})
	}

	allFollowers := memory.Followers
	allFollowers[follower.UserName] = follows
}