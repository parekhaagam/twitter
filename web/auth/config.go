package auth

type Configuration struct {
	timeout int
}

var Config = Configuration{10}