package auth

import (
	"../auth/storage/memory"
	"net/http"
	"time"
)

/*type Auth struct {
	storage storage.Storage
}*/

func GetToken(userId string) (string){
	memory.AuthObject.M.Lock()
	defer memory.AuthObject.M.Unlock()
	var val,ok = memory.AuthObject.UserMap[userId]
	if ok {
		memory.AuthObject.M.Unlock()
			return val.Token
	}else {
		var tokenDetailsObject = memory.TokenDetails{UserId: userId,Token: userId,TimeGenerated:time.Now()}
		memory.AuthObject.UserMap[userId] = tokenDetailsObject
		var token = tokenDetailsObject.Token
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
