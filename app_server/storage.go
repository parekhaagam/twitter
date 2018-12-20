package app_server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/namespace"
	"github.com/google/uuid"
	"github.com/parekhaagam/twitter/globals"

	"sort"
	"strings"
	"time"
	log "github.com/sirupsen/logrus"
)

var etcdClient *clientv3.Client
var kvStore clientv3.KV
var etcdErr error

func getEtcdClientObjects() (*clientv3.Client,clientv3.KV,error){
	fmt.Println("inside getetcd clients")
	if etcdClient==nil && kvStore == nil{
		fmt.Println("inside if")
		etcdClient, etcdErr = clientv3.New(clientv3.Config{Endpoints: []string{RAFT_ENDPOINT1,RAFT_ENDPOINT2, RAFT_ENDPOINT3}})
		if etcdErr != nil {
			fmt.Println("error in get etcd client object", etcdErr)
		}
		kvStore = etcdClient.KV
		etcdClient.KV = namespace.NewKV(etcdClient.KV, APP_PREFIX)
		etcdClient.Watcher = namespace.NewWatcher(etcdClient.Watcher, APP_PREFIX)
		etcdClient.Lease = namespace.NewLease(etcdClient.Lease, APP_PREFIX)
	}

	return etcdClient, kvStore,nil
}

func InsertUserRecord(userId string , password string)(bool,error){
	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		return false,etcdErr
	}

	fmt.Println("InsertUserRecord user id:" , userId)
	isUserExist,err :=CheckUserExist(userId)
	if err!=nil {
		return false,err
	} else if !isUserExist {

		log.Info("InsertUserRecord")
		//log.Debug("userId:%v \t password: %v", userId, password)
		var userRecord []string
		userRecord = append(userRecord, USER_RECORD_MAP_PREFIEX, userId)
		response, responseErr := etcdClient.Put(context.TODO(), strings.Join(userRecord, ""), password)
		if responseErr!=nil {
			return false,responseErr
		}
		fmt.Println("insert single user", response, responseErr)

		//var allUser []string

		var allUserRecord [] string
		allUserRecord = append(allUserRecord, ALL_USER_LIST_PREFIX)

		allUserResponse,error := etcdClient.Get(context.TODO(), strings.Join(allUserRecord, ""))
		if error!=nil {
			return false,error
		}
		var allUsersList []globals.User
		if allUserResponse != nil && len(allUserResponse.Kvs) != 0 {
			unMarshalError := json.Unmarshal(allUserResponse.Kvs[0].Value, &allUsersList)
			if unMarshalError!=nil {
				return false,unMarshalError
			}
		}else{
			allUsersList = []globals.User{}
		}

		allUsersList = append(allUsersList, globals.User{userId})
		jsonUserListObject,marshalError := json.Marshal(allUsersList)
		if marshalError!=nil {
			return false,marshalError
		}
		_, allUserRespoError := etcdClient.Put( context.TODO(), strings.Join(allUserRecord, ""), string(jsonUserListObject))

		if allUserRespoError != nil{
			return false,allUserRespoError
		}

		fmt.Println("insert in all user", response,responseErr)
		} else{
		return false,nil
	}

	return true,nil
}

func CheckUserRecord(userId string, password string ) (bool)  {

	log.Info("CheckUserRecord")
	log.Debug("userId:"  + userId + " \t password: " + password)

	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
	}
	var userRecord []string
	userRecord = append(userRecord,USER_RECORD_MAP_PREFIEX,userId)
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

func CheckUserExist(userId string) (bool,error)  {

	log.Info("CheckUserExist")
	//log.Debug("userId:%v ", userId)

	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		return false,etcdErr
	}

	var userRecord []string
	userRecord = append(userRecord,USER_RECORD_MAP_PREFIEX,userId)
	responsePassword,_ :=etcdClient.Get(context.TODO(),strings.Join(userRecord,""))
	if responsePassword !=nil && responsePassword.Kvs !=nil {
		fmt.Println("response password false")
		return true,nil
	}
	fmt.Println("response password true")
	return false,nil

}

func StorageFollowUser(follower globals.User, selectedUserNames []string) error{

	log.Info("StorageFollowUser")
	//log.Debug("follower userId:%v \tselected User to follow:%v", follower.UserName, selectedUserNames)

	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
		return etcdErr
	}

	follows := make([]globals.User, 0)
	for _, userName := range selectedUserNames {
		follows = append(follows, globals.User{userName})
	}

	var allUser []string
	allUser = append(allUser, FOLLOWERS_MAP_PREFIX, follower.UserName)

	jsonFollowerUserObject,_ := json.Marshal(follows)
	response, responseErr := etcdClient.Put(context.TODO(), strings.Join(allUser, ""), string(jsonFollowerUserObject))
	fmt.Println("StorageFollowUser: ", response, responseErr)
	return nil
}

func GetAllFollowingUser(user globals.User) ([]globals.User,error){

	log.Info("GetAllFollowingUser")
	//log.Debug("userId:%v ", user.UserName)

	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		return nil,etcdErr
	}

	var allUser []string
	allUser = append(allUser, FOLLOWERS_MAP_PREFIX, user.UserName)

	allUserResponse,_ := etcdClient.Get(context.TODO(), strings.Join(allUser, ""))
	userList := []globals.User{}
	if allUserResponse != nil && len(allUserResponse.Kvs) != 0 {
		json.Unmarshal(allUserResponse.Kvs[0].Value, &userList)
	}

	return userList,nil
}

func StorageInsertTweets(user globals.User, content string) (string,error){

	log.Info("StorageInsertTweets")
	//log.Debug("userId:%v \t content string:%v ", user.UserName, content)
	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr!=nil {
		return "",ETCD_ERROR
	}
	TID := uuid.New().String()
	fmt.Println(TID)

	var tweetTIDStored []string
	tweetTIDStored = append(tweetTIDStored, TWEET_TID_STORED_MAP_PREFIX, TID)

	response, responseErr := etcdClient.Put(context.TODO(), strings.Join(tweetTIDStored, ""), TID)
	if responseErr!=nil {
		return "",responseErr
	}
	fmt.Println("storage insert tweet", response, responseErr )

	tmp := globals.Tweet{
		Content:content,
		Timestamp: time.Now().Unix(),
		TID: TID,
		UserId:user.UserName,
	}

	var tweetStored []string
	tweetStored = append(tweetStored, USER_TWEET_MAP_PREFIX, user.UserName)
	allTweetInsertResponse,getError := etcdClient.Get(context.TODO(), strings.Join(tweetStored, ""))
	if getError!=nil {
		return "",getError
	}

	var twitterTweet []globals.Tweet
	fmt.Println(allTweetInsertResponse.Kvs)
	if len(allTweetInsertResponse.Kvs) == 0{
		twitterTweet = []globals.Tweet{tmp}
	}else{
		twitterTweet = []globals.Tweet{}
		unMarshalError :=json.Unmarshal(allTweetInsertResponse.Kvs[0].Value, &twitterTweet)
		if unMarshalError!=nil {
			return "",unMarshalError
		}
		twitterTweet = append(twitterTweet, tmp)
	}

	tweetJsonObject, marshalError := json.Marshal(twitterTweet)
	if marshalError != nil{
		return "",marshalError
	}
	_, tweetInsertResponseError := etcdClient.Put(context.TODO(), strings.Join(tweetStored, ""), string(tweetJsonObject))

	if tweetInsertResponseError != nil{
		return "",tweetInsertResponseError
	}

	return TID,nil
}

func StorageGetFollowersTweets(followings []globals.User) ([]globals.Tweet,error){

	log.Info("StorageGetFollowersTweets")
	log.Debug("following Users id  ", followings)

	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr!=nil {
		return nil,ETCD_ERROR
	}
	var followingTweet []globals.Tweet

	twitterTweet := []globals.Tweet{}

	for _,following := range followings{

		var tweetStored []string
		tweetStored = append(tweetStored, USER_TWEET_MAP_PREFIX, following.UserName)
		allTweetInsertResponse,_ := etcdClient.Get(context.TODO(), strings.Join(tweetStored, ""))

		fmt.Println("storage get followers tweets", following.UserName)
		fmt.Println(len(allTweetInsertResponse.Kvs))

		if allTweetInsertResponse != nil && len(allTweetInsertResponse.Kvs) != 0 {
			json.Unmarshal(allTweetInsertResponse.Kvs[0].Value, &twitterTweet)
			fmt.Println("twitterTweet:", twitterTweet[0].Content)
			followingTweet = append(followingTweet, twitterTweet...)
		}
	}

	sort.Slice(followingTweet[:], func(i, j int) bool {
		return followingTweet[i].Timestamp > followingTweet[j].Timestamp
	})

	for index, _ := range followingTweet{
		followingTweet[index].TimeMessage = TimeToString(followingTweet[index].Timestamp)
	}

	return followingTweet,nil
}

func Storage_Get_all_users(loggedInUserId string) (globals.UserList,error){

	log.Info("Storage_Get_all_users")
	log.Debug("user id: ", loggedInUserId)

	var users []globals.UserFollowed

	var allUser []string
	allUser = append(allUser, ALL_USER_LIST_PREFIX)
	allUserResponse,etcdError := etcdClient.Get(context.TODO(), strings.Join(allUser, ""))
	if etcdError!=nil {
		return globals.UserList{},ETCD_ERROR
	}
	var allUsers []globals.User
	if allUserResponse != nil && len(allUserResponse.Kvs) != 0 {
		json.Unmarshal(allUserResponse.Kvs[0].Value, &allUsers)
	}

	for _,eachUser := range allUsers{
		fmt.Print(eachUser.UserName)
	}
	loggedInUser := globals.User{loggedInUserId} //should come from session
	for _,user := range allUsers {
		if user.UserName != loggedInUser.UserName {
			isFollows,err :=StorageFollows(loggedInUser, user)
			if err!=nil {
				return globals.UserList{},err
			}
			users = append(users, globals.UserFollowed{user.UserName, isFollows})
		}
	}
	fmt.Println(users)
	return globals.UserList{users, false},nil

}

func StorageFollows(user1 globals.User, user2 globals.User) (bool,error){

	log.Info("StorageFollows")
	log.Debug("first user id: \t second user id:", user1.UserName, user2.UserName)

	followers,err := GetAllFollowingUser(user1)
	if err!=nil {
		return false,err
	} else {
		doesFollow := false

		for _,followedUser := range followers {
			if followedUser.UserName == user2.UserName {
				doesFollow = true
				break
			}
		}
		return doesFollow,nil
	}

}