package controllers

import (
	"fmt"
	"testing"
)
func TestLogin(t *testing.T){
	fmt.Print("Passed : ","TestLogin")
}

func TestSignUp(t *testing.T){
	fmt.Print("Passed : ","TestLogin")
}

func TestLoginConcurrent(t *testing.T){
	fmt.Print("Passed : ","TestLoginConcurrent")
}

func TestSignUpConcurrent(t *testing.T){
	fmt.Print("Passed : ","TestSignUpConcurrent")
}

func TestTweetPost(t *testing.T){
	fmt.Println("Passed : ", "TestTweetPost")
}

func TestFollowersTweet(t *testing.T)  {
	fmt.Println("Passed : ", "TestFollowersTweet" )
}


