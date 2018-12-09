package auth_server

import (
	"fmt"
	"sync"
	"testing"
)

func TestAuthGetToken(t *testing.T)  {
	var  token  = getToken("abc@gmail.com")
	if token!="" {
		fmt.Println("Passed : ","TestAuthGetToken")
	}else {
		t.Fatal("Error : Token not generated properly!")
	}
}
func TestAuthGetTokenConcurrent(t *testing.T)  {
	set := make(map[string]string)
	var wg sync.WaitGroup
	for i:=1 ; i<6 ; i++ {
		wg.Add(1)
		go func(userId string) {
			defer wg.Done()
			var  token  = getToken(userId)
			set[token] = userId
		}("def@gmail.com")
	}
	wg.Wait()
	fmt.Println(len(set))
	if len(set)!=1{
		t.Fatal("TokenGenerator is generating multiple token for same user when accessed concurrently")
	}else {
		fmt.Println("Passed : ","TestAuthGetTokenConcurrent")
	}
}
func TestAuthAuthenticateToken(t *testing.T)  {
	var  token  = getToken("abc@gmail.com")
	if token!="" {
		if !IsTokenValid(token){
			t.Fatal("Error : Token not authenticated properly!")
		}
	}else {
		t.Fatal("Error : Token not generated properly!")
	}
	fmt.Println("Passed : ","TestAuthAuthenticateToken")
}
func TestAuthAuthenticateTokenConcurrent(t *testing.T)  {
	var  token  = getToken("abc@gmail.com")
	set := make(map[string]int)
	if token!="" {
		var wg sync.WaitGroup
		for i:=1 ; i<6 ; i++ {
			wg.Add(1)
			go func(token string) {
				defer wg.Done()
				if !IsTokenValid(token){
					set[token]=1
				}
			}(token)
		}
		wg.Wait()
		if len(set)!=0{
			t.Fatal("Error : Token not authenticated in ",len(set)," attempt(s)")
		}
	}else {
		t.Fatal("Error : Token not generated properly!")
	}
	fmt.Println("Passed : ","TestAuthAuthenticateTokenConcurrent")
}
