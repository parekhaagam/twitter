package storage

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/namespace"

	"strings"
)

type Storage interface {
	GetOrCreateToken()
	ValidateToken()
}

var etcdClient *clientv3.Client
var kvStore clientv3.KV
var etcdErr error

func getEtcdClientObjects() (*clientv3.Client,clientv3.KV,error){
	if etcdClient==nil && kvStore == nil{
		etcdClient, etcdErr = clientv3.New(clientv3.Config{Endpoints: []string{RAFT_ENDPOINT}})
		if etcdErr != nil {
			// handle error!
		}
		kvStore = etcdClient.KV
		etcdClient.KV = namespace.NewKV(etcdClient.KV, APP_PREFIX)
		etcdClient.Watcher = namespace.NewWatcher(etcdClient.Watcher, APP_PREFIX)
		etcdClient.Lease = namespace.NewLease(etcdClient.Lease, APP_PREFIX)
	}
	return etcdClient, kvStore,nil
}

func InsertUserRecord(userId string , password string)(bool){
	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
	}

	if !CheckUserExist(userId) {
		var userRecord []string
		userRecord = append(userRecord, APP_PREFIX, USER_RECORD_MAP_PREFIEX, userId)
		response, responseErr := etcdClient.Put(context.TODO(), strings.Join(userRecord, ""), password)
		fmt.Println(response, responseErr)
	} else{
		return false
	}
	/*var allUserRecord []string
	allUserRecord = append(allUserRecord,APP_PREFIX, ALL_USER_LIST_PREFIX)
	etcdClient.Get(context.TODO(),strings.Join(allUserRecord,""))
	currUser := globals.User{"manish.n"}
	*/
	return true
}
func CheckUserRecord(userId string, password string ) (bool)  {
	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
	}
	var userRecord []string
	userRecord = append(userRecord,APP_PREFIX,USER_RECORD_MAP_PREFIEX,userId)
	responsePassword,_ :=etcdClient.Get(context.TODO(),strings.Join(userRecord,""))
	if responsePassword !=nil && responsePassword.Kvs !=nil {
		passwordCheck := string(responsePassword.Kvs[0].Value)
		if passwordCheck == password{
			return true
		}
		return false
	}
	return false
}

func CheckUserExist(userId string) (bool)  {
	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
	}
	var userRecord []string
	userRecord = append(userRecord,APP_PREFIX,USER_RECORD_MAP_PREFIEX,userId)
	responsePassword,_ :=etcdClient.Get(context.TODO(),strings.Join(userRecord,""))
	if responsePassword !=nil && responsePassword.Kvs !=nil {
		passwordCheck := string(responsePassword.Kvs[0].Value)
		if passwordCheck == ""{
			return false
		}
		return true
	}
	return true
}
