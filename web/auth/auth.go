package auth

import (
	"github.com/parekhaagam/twitter/web/auth/storage/memory"
	"github.com/google/uuid"
	"net/http"
	"time"
)

/*type Auth struct {
	storage storage.Storage
}*/

func GetToken(userId string) (string){
	memory.AuthObject.M.Lock()
	defer memory.AuthObject.M.Unlock()
	var val,ok = memory.AuthObject.LogedInUserMap[userId]
	if ok {
		//memory.AuthObject.M.Unlock()
			return val.Token
	}else {
		var token = uuid.New().String()
		var tokenDetailsObject = memory.TokenDetails{UserId: userId,Token: token,TimeGenerated:time.Now()}
		memory.AuthObject.LogedInUserMap[userId] = tokenDetailsObject
		memory.AuthObject.TokenMap[token] = tokenDetailsObject;
		return token
	}
}
func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token,err = r.Cookie("token")
		if err==nil {
			if IsTokenValid(token.Value){
				next.ServeHTTP(w, r)
			}else {
				w.Write([]byte("Invalid Token"))
			}
		}else {
			w.Write([]byte("Invalid Token"))
		}

	})
}
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
