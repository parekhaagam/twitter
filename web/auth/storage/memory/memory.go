package memory

import (
	"sync"
	"time"
)

type TokenDetails struct {
	UserId string
	TimeGenerated time.Time
	Token string
}
type Authentication struct {
	IsStarted bool
	M sync.Mutex
	UserMap map[string]TokenDetails
	TokenMap map[string]TokenDetails
}
var AuthObject  = Authentication{IsStarted:false,
	UserMap:make(map[string]TokenDetails),
	TokenMap: make(map[string]TokenDetails)}

func (auth *Authentication)StartAuthObject(){
	AuthObject.IsStarted = true;
}

func (auth *Authentication)StopAuthObject(){
	AuthObject.IsStarted = false;
}