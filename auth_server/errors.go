package auth_server

import "errors"

var ETCD_ERROR = errors.New("ETCD servers not available")
var INVALID_TOKEN = errors.New("Invalid Token")
