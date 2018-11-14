package auth

import (
	"../auth/storage/memory"
	"fmt"
	"net/http"
	"time"
)

func GetToken(userid string) (string){
	memory.AuthObject.M.Lock()
	defer memory.AuthObject.M.Unlock()
	var val,ok = memory.AuthObject.UserMap[userid]
	if ok {
		memory.AuthObject.M.Unlock()
			return val.Token
	}else {
		var tokenDetailsObject = memory.TokenDetails{UserId:userid,Token:userid,TimeGenerated:time.Now()}
		memory.AuthObject.UserMap[userid] = tokenDetailsObject
		var token = tokenDetailsObject.Token
		memory.AuthObject.TokenMap[token] = tokenDetailsObject;
		return token
	}
}
func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		memory.AuthObject.M.Lock()
		defer memory.AuthObject.M.Unlock()
		fmt.Println("auth called")
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
