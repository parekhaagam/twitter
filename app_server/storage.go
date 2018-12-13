package app_server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/namespace"
	"github.com/google/uuid"
	"github.com/parekhaagam/twitter/globals"
	"log"
	"sort"
	"strings"
	"time"
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

func InsertUserRecord(userId string , password string)(bool){
	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
	}

	fmt.Println("InsertUserRecord user id:" , userId)
	if !CheckUserExist(userId) {

		var userRecord []string
		userRecord = append(userRecord, USER_RECORD_MAP_PREFIEX, userId)
		response, responseErr := etcdClient.Put(context.TODO(), strings.Join(userRecord, ""), password)
		fmt.Println("insert single user", response, responseErr)

		//var allUser []string

		var allUserRecord [] string
		allUserRecord = append(allUserRecord, ALL_USER_LIST_PREFIX)

		allUserResponse,_ := etcdClient.Get(context.TODO(), strings.Join(allUserRecord, ""))
		var allUsersList []globals.User
		if allUserResponse != nil && len(allUserResponse.Kvs) != 0 {
			json.Unmarshal(allUserResponse.Kvs[0].Value, &allUsersList)
		}else{
			allUsersList = []globals.User{}
		}

		allUsersList = append(allUsersList, globals.User{userId})
		jsonUserListObject,_ := json.Marshal(allUsersList)
		_, allUserRespoError := etcdClient.Put( context.TODO(), strings.Join(allUserRecord, ""), string(jsonUserListObject))

		if allUserRespoError != nil{
			log.Fatal(allUserResponse)
		}

		fmt.Println("insert in all user", response,responseErr)
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

func CheckUserExist(userId string) (bool)  {
	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
		log.Fatal(etcdErr)
	}

	var userRecord []string
	userRecord = append(userRecord,USER_RECORD_MAP_PREFIEX,userId)
	responsePassword,_ :=etcdClient.Get(context.TODO(),strings.Join(userRecord,""))
	if responsePassword !=nil && responsePassword.Kvs !=nil {
		fmt.Println("response password false")
		return true
	}
	fmt.Println("response password true")
	return false

}

func StorageFollowUser(follower globals.User, selectedUserNames []string){

	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
		log.Fatal(etcdErr)
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
}

func GetAllFollowingUser(user globals.User) ([]globals.User){

	etcdClient,kvStore,etcdErr = getEtcdClientObjects();
	if etcdErr != nil {
		//implement while considering exceptions
		log.Fatal(etcdErr)
	}

	var allUser []string
	allUser = append(allUser, FOLLOWERS_MAP_PREFIX, user.UserName)

	allUserResponse,_ := etcdClient.Get(context.TODO(), strings.Join(allUser, ""))
	userList := []globals.User{}
	if allUserResponse != nil && len(allUserResponse.Kvs) != 0 {
		json.Unmarshal(allUserResponse.Kvs[0].Value, &userList)
	}

	return userList
}

func StorageInsertTweets(user globals.User, content string)string{

	etcdClient,kvStore,etcdErr = getEtcdClientObjects();

	TID := uuid.New().String()
	fmt.Println(TID)

	var tweetTIDStored []string
	tweetTIDStored = append(tweetTIDStored, TWEET_TID_STORED_MAP_PREFIX, TID)

	response, responseErr := etcdClient.Put(context.TODO(), strings.Join(tweetTIDStored, ""), TID)
	fmt.Println("storage insert tweet", response, responseErr )

	tmp := globals.Tweet{
		Content:content,
		Timestamp: time.Now().Unix(),
		TID: TID,
		UserId:user.UserName,
	}

	var tweetStored []string
	tweetStored = append(tweetStored, USER_TWEET_MAP_PREFIX, user.UserName)
	allTweetInsertResponse,_ := etcdClient.Get(context.TODO(), strings.Join(tweetStored, ""))

	var twitterTweet []globals.Tweet
	fmt.Println(allTweetInsertResponse.Kvs)
	if len(allTweetInsertResponse.Kvs) == 0{
		twitterTweet = []globals.Tweet{tmp}
	}else{
		twitterTweet = []globals.Tweet{}
		json.Unmarshal(allTweetInsertResponse.Kvs[0].Value, &twitterTweet)
		twitterTweet = append(twitterTweet, tmp)
	}

	tweetJsonObject, err := json.Marshal(twitterTweet)
	if err != nil{
		log.Fatal(err.Error())
	}
	_, tweetInsertResponseError := etcdClient.Put(context.TODO(), strings.Join(tweetStored, ""), string(tweetJsonObject))

	if tweetInsertResponseError != nil{
		log.Fatal(tweetInsertResponseError.Error())
		return TID
	}

	return ""
}

func StorageGetFollowersTweets(followings []globals.User)[]globals.Tweet{

	etcdClient,kvStore,etcdErr = getEtcdClientObjects();

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

	return followingTweet
}

func Storage_Get_all_users(loggedInUserId string) (ul globals.UserList){

	var users []globals.UserFollowed

	var allUser []string
	allUser = append(allUser, ALL_USER_LIST_PREFIX)
	allUserResponse,_ := etcdClient.Get(context.TODO(), strings.Join(allUser, ""))
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
			users = append(users, globals.UserFollowed{user.UserName, StorageFollows(loggedInUser, user)})
		}
	}
	fmt.Println(users)
	return globals.UserList{users, false}

}

func StorageFollows(user1 globals.User, user2 globals.User) bool{
	followers := GetAllFollowingUser(user1)

	doesFollow := false

		for _,followedUser := range followers {
			if followedUser.UserName == user2.UserName {
				doesFollow = true
				break
			}
		}
	return doesFollow
}
/*
userList := globals.UserList{}
		//fmt.Println(allUserResponse.Kvs)
		if allUserResponse != nil && len(allUserResponse.Kvs) != 0 {
			fmt.Println("appending existing data")
			json.Unmarshal(allUserResponse.Kvs[0].Value, &userList)
		} else{
			userList = globals.UserList{[]globals.UserFollowed{}, false}
		}

		userList.List = append(userList.List, globals.UserFollowed{userId, true})
		jsonUserListObject,_ := json.Marshal(userList)
		response, responseErr = etcdClient.Put(context.TODO(), strings.Join(allUser, ""), string(jsonUserListObject))

 */