package auth

import (
	"../auth/storage/memory"
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
		memory.AuthObject.M.Lock()
		defer memory.AuthObject.M.Unlock()
		//fmt.Println("auth called")
		var token,err = r.Cookie("token")
		if err==nil {
			var _, ok= memory.AuthObject.TokenMap[token.Value]
			if ok {
				next.ServeHTTP(w, r)
			} else {
				w.Write([]byte("Invalid Token"))
			}
		}else {
			w.Write([]byte("Invalid Token"))
		}

	})
}
