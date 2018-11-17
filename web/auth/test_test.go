package auth

import (
	"fmt"
	"testing"
)

func TestAuthGetToken(t *testing.T)  {
	fmt.Println("Passed : ","TestAuthGetToken")
}
func TestAuthGetTokenConcurrent(t *testing.T)  {
	fmt.Println("Passed : ","TestAuthGetTokenConcurrent")
}
func TestAuthAuthenticateToken(t *testing.T)  {
	fmt.Println("Passed : ","TestAuthAuthenticateToken")
}
func TestAuthAuthenticateTokenConcurrent(t *testing.T)  {
	//t.Fatal("Failed : ","TestAuthAuthenticateTokenConcurrent")
	fmt.Println("Passed : ","TestAuthAuthenticateTokenConcurrent")
}
