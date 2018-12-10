package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/google/uuid"
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
	key = append(key,AUTH_PREFIX,LOGGED_IN_USER_PREFIX,userId)
	token,_ :=kvStore.Get(context.TODO(),strings.Join(key,""))
	if token !=nil && token.Kvs !=nil {
		tokenDetail := memory.TokenDetails{}
		json.Unmarshal(token.Kvs[0].Value,&tokenDetail)
		return tokenDetail.Token
	}else {
		//Get()
		token := uuid.New().String()
		var loggedInUserMapKey []string
		loggedInUserMapKey = append(loggedInUserMapKey,LOGGED_IN_USER_PREFIX,userId)
		tokenDetailsJsonObject,_ :=json.Marshal(memory.TokenDetails{UserId: userId,Token: token,TimeGenerated:time.Now()})
		tokenDetailsJsonString := string(tokenDetailsJsonObject)
		response,responseErr :=etcdClient.Put(context.TODO(),strings.Join(loggedInUserMapKey,""),tokenDetailsJsonString)
		fmt.Println(response,responseErr)
		var tokenMapKey []string
		tokenMapKey = append(tokenMapKey,TOKEN_MAP_PREFIX,token)
		etcdClient.Put(context.TODO(),strings.Join(tokenMapKey,""),tokenDetailsJsonString)
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
		etcdClient, etcdErr = clientv3.New(clientv3.Config{Endpoints: []string{RAFT_ENDPOINT}})
		if etcdErr != nil {
			// handle error!
		}
		kvStore = etcdClient.KV
		etcdClient.KV = namespace.NewKV(etcdClient.KV, AUTH_PREFIX)
		etcdClient.Watcher = namespace.NewWatcher(etcdClient.Watcher, AUTH_PREFIX)
		etcdClient.Lease = namespace.NewLease(etcdClient.Lease, AUTH_PREFIX)
	}
	return etcdClient, kvStore,nil
}

func IsTokenValid(token string) (bool) {
	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
	}
	var key []string
	fmt.Println(token)
	key = append(key,AUTH_PREFIX,TOKEN_MAP_PREFIX,token)
	userId,_ :=kvStore.Get(context.TODO(),strings.Join(key,""))
	if userId !=nil && userId.Kvs!=nil {
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