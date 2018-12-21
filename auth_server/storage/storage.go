package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/namespace"
	"github.com/google/uuid"
	"github.com/parekhaagam/twitter/auth_server/storage/memory"
	"strings"
	"sync"
	"time"
)

type Storage interface {
	GetOrCreateToken()
	ValidateToken()
}
var etcdClient *clientv3.Client
var kvStore clientv3.KV
var etcdErr error
var etcdMutex sync.Mutex
func GetOrCreateToken(userId string) (string,error)  {
	kvStore,etcdMutex,etcdErr := getEtcdClientObjects();
	etcdMutex.Lock()
	defer etcdMutex.Unlock()
	if etcdErr != nil {
		return "",ETCD_ERROR
	}
	var key []string
	key = append(key,AUTH_PREFIX,LOGGED_IN_USER_PREFIX,userId)
	token,_ :=kvStore.Get(context.TODO(),strings.Join(key,""))
	if token !=nil && token.Kvs !=nil {
		tokenDetail := memory.TokenDetails{}
		json.Unmarshal(token.Kvs[0].Value,&tokenDetail)
		return tokenDetail.Token,nil
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
		return token,nil
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
func getEtcdClientObjects() (clientv3.KV,sync.Mutex,error){
	if etcdClient==nil && kvStore == nil{
		etcdClient, etcdErr = clientv3.New(clientv3.Config{Endpoints: []string{RAFT_ENDPOINT1,RAFT_ENDPOINT2, RAFT_ENDPOINT3}})
		if etcdErr != nil {
			return nil,etcdMutex,etcdErr
		}
		kvStore = etcdClient.KV
		etcdClient.KV = namespace.NewKV(etcdClient.KV, AUTH_PREFIX)
		etcdClient.Watcher = namespace.NewWatcher(etcdClient.Watcher, AUTH_PREFIX)
		etcdClient.Lease = namespace.NewLease(etcdClient.Lease, AUTH_PREFIX)
	}
	return  kvStore,etcdMutex,nil
}

func IsTokenValid(token string) (bool,error) {
	KVStore,etcdMutex,etcdErr := getEtcdClientObjects();
	etcdMutex.Lock()
	defer etcdMutex.Unlock()
	if etcdErr != nil {
		return false,ETCD_ERROR
	}
	var key []string
	fmt.Println(token)
	key = append(key,AUTH_PREFIX,TOKEN_MAP_PREFIX,token)
	userId,_ :=KVStore.Get(context.TODO(),strings.Join(key,""))
	if userId !=nil && userId.Kvs!=nil {
		if string(userId.Kvs[0].Value) != "" {
			return true,nil
		}
	}
	return false,INVALID_TOKEN
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