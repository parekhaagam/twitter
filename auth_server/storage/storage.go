package storage

import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	"github.com/google/uuid"
	"github.com/parekhaagam/twitter/auth_server"
	"github.com/coreos/etcd/clientv3/namespace"
	"github.com/parekhaagam/twitter/auth_server/storage/memory"
	"strings"
	"time"
)

type Storage interface {
	GetOrCreateToken()
	ValidateToken()
}
var etcdClient *clientv3.Client
var kvStore clientv3.KV
var etcdErr error
func GetOrCreateToken(userId string) (string)  {
	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
	}
	var key []string
	key = append(key,auth_server.AUTH_PREFIX,auth_server.LOGGED_IN_USER_PREFIX,userId)
	token,_ :=kvStore.Get(context.TODO(),strings.Join(key,""))
	if token !=nil {
		tokenDetail := memory.TokenDetails{}
		json.Unmarshal(token.Kvs[0].Value,&tokenDetail)
		return tokenDetail.Token
	}else {
		token := uuid.New().String()
		var loggedInUserMapKey []string
		loggedInUserMapKey = append(loggedInUserMapKey,auth_server.AUTH_PREFIX,auth_server.LOGGED_IN_USER_PREFIX,userId)
		tokenDetailsJsonObject,_ :=json.Marshal(memory.TokenDetails{UserId: userId,Token: token,TimeGenerated:time.Now()})
		tokenDetailsJsonString := string(tokenDetailsJsonObject)
		etcdClient.Put(context.TODO(),strings.Join(loggedInUserMapKey,""),tokenDetailsJsonString)

		var tokenMapKey []string
		tokenMapKey = append(tokenMapKey,auth_server.AUTH_PREFIX,auth_server.TOKEN_MAP_PREFIX,tokenDetailsJsonString)
		etcdClient.Put(context.TODO(),strings.Join(tokenMapKey,""),userId)
		return token
	}
}
/*
func getToken(userId string) (string)  {
	memory.AuthObject.M.Lock()
	defer memory.AuthObject.M.Unlock()
	var val,ok = memory.AuthObject.LogedInUserMap[userId]
	if ok {
		//memory.AuthObject.M.Unlock()
		return val.Token
	}else {
		var token = uuid. New().String()
		var tokenDetailsObject = memory.TokenDetails{UserId: userId,Token: token,TimeGenerated:time.Now()}
		memory.AuthObject.LogedInUserMap[userId] = tokenDetailsObject
		memory.AuthObject.TokenMap[token] = tokenDetailsObject;
		return token
	}

}
*/
func getEtcdClientObjects() (*clientv3.Client,clientv3.KV,error){
	if etcdClient==nil && kvStore == nil{
		etcdClient, etcdErr = clientv3.New(clientv3.Config{Endpoints: []string{"localhost:2379"}})
		if etcdErr != nil {
			// handle error!
		}
		kvStore = etcdClient.KV
		etcdClient.KV = namespace.NewKV(etcdClient.KV, auth_server.AUTH_PREFIX)
		etcdClient.Watcher = namespace.NewWatcher(etcdClient.Watcher, auth_server.AUTH_PREFIX)
		etcdClient.Lease = namespace.NewLease(etcdClient.Lease, auth_server.AUTH_PREFIX)
	}
	return etcdClient, kvStore,nil
}

func IsTokenValid(token string) (bool) {
	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
	}
	var key []string
	key = append(key,auth_server.AUTH_PREFIX,auth_server.TOKEN_MAP_PREFIX,token)
	userId,_ :=kvStore.Get(context.TODO(),strings.Join(key,""))
	if userId !=nil {
		if string(userId.Kvs[0].Value) != "" {
			return true
		}
	}
	return false
}
/*func IsTokenValid(token string) (bool){
	memory.AuthObject.M.Lock()
	defer memory.AuthObject.M.Unlock()
	var _, ok= memory.AuthObject.TokenMap[token]
	if ok {
		return true
	} else {
		return false
	}
}
*/