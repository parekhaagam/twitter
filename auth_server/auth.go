package auth_server

import (
	"github.com/parekhaagam/twitter/auth_server/storage/memory"
)

/*type Auth struct {
	contract contract.Storage
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
