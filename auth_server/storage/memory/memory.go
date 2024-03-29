package memory

import (
	"sync"
	"time"
)

/*type MemoryAuth struct {

}
*/
type TokenDetails struct {
	UserId string `json:"UserId"`
	TimeGenerated time.Time `json:"TimeGenerated"`
	Token string `json:"Token"`
}
type Authentication struct {
	IsStarted      bool
	M              sync.Mutex
	LogedInUserMap map[string]TokenDetails
	TokenMap       map[string]TokenDetails
}
var AuthObject  = Authentication{IsStarted:false,
	LogedInUserMap: make(map[string]TokenDetails),
	TokenMap:       make(map[string]TokenDetails)}

func (auth *Authentication)StartAuthObject(){
	AuthObject.IsStarted = true;
}

func (auth *Authentication)StopAuthObject(){
	AuthObject.IsStarted = false;
}