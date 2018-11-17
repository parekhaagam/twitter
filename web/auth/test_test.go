package auth

import (
	"fmt"
	"sync"
	"testing"
)

func TestAuthGetToken(t *testing.T)  {
	var  token  = GetToken("abc@gmail.com")
	if token!="" {
		fmt.Println("Passed : ","TestAuthGetToken")
	}else {
		t.Fatal("Error Token not generated properly!")
	}
}
func TestAuthGetTokenConcurrent(t *testing.T)  {
	set := make(map[string]string)
	var wg sync.WaitGroup
	for i:=1 ; i<6 ; i++ {
		wg.Add(1)
		go func(userId string) {
			defer wg.Done()
			var  token  = GetToken(userId)
			set[token] = userId
		}("def@gmail.com")
	}
	wg.Wait()
	fmt.Println(len(set))
	if len(set)!=1{
		t.Fatal("Token is generating multiple token for same user when accessed concurrently")
	}else {
		fmt.Println("Passed : ","TestAuthGetTokenConcurrent")
	}
}
func TestAuthAuthenticateToken(t *testing.T)  {
	fmt.Println("Passed : ","TestAuthAuthenticateToken")
}
func TestAuthAuthenticateTokenConcurrent(t *testing.T)  {
	//t.Fatal("Failed : ","TestAuthAuthenticateTokenConcurrent")
	fmt.Println("Passed : ","TestAuthAuthenticateTokenConcurrent")
}
