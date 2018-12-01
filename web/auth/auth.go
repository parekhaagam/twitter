package auth

import (
	"github.com/parekhaagam/twitter/web/auth/storage/memory"
)

/*type Auth struct {
	storage storage.Storage
}*/

func IsTokenValid(token string) (bool){
	memory.AuthObject.M.Lock()
	defer memory.AuthObject.M.Unlock()
	var _, ok= memory.AuthObject.TokenMap[token]
	if ok {
		return true
	} else {
		return false
	}
}
