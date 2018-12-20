package app_server

import (
	"context"
	"fmt"
	pb "github.com/parekhaagam/twitter/app_server/contract"
	"github.com/parekhaagam/twitter/globals"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"strconv"
)

type StorageServer struct {
	srv *http.Server

	//contentDaemon *contentd.Contentd
}

type StorageServerImpl struct{}

func (a *StorageServerImpl)InsertTweets(ctx context.Context, in *pb.InsertTweetRequest) (*pb.InsertTweetReply, error){

	user := globals.User{in.User.UserName}
	//tid := InsertTweets(user, in.Content)
	tid,error:= StorageInsertTweets(user, in.Content)
	if error!=nil {
		return nil,error
	}
	return &pb.InsertTweetReply{TID:tid},nil
}


func (a *StorageServerImpl)GetFollowersTweets(ctx context.Context, in *pb.GetFollowersRequest)(*pb.GetFollowersReply, error){

	//followings []globals.User :=

	currUser := globals.User{in.User.UserName}

	//following := GetAllFollowing(currUser)
	following,err := GetAllFollowingUser(currUser)
	if err!=nil {
		return nil,err
	}

	followingNumber := len(following)
	following = append(following, currUser)

	//tweets := GetFollowersTweets(following)
	tweets,error := StorageGetFollowersTweets(following)
	if error!=nil {
		return nil,error
	}
	fmt.Println(tweets)

	var tweetArray []*pb.Tweet
	for _,t := range tweets{
		tTemp := pb.Tweet{UserId:t.UserId, TID:t.TID, Timestamp:t.Timestamp, Content:t.Content, TimeMessage:t.TimeMessage}
		tweetArray = append(tweetArray, &tTemp)
	}
	return &pb.GetFollowersReply{Tweets: tweetArray,FollowingNumber:strconv.Itoa(followingNumber)},nil
}

func (a *StorageServerImpl) GetAllUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error){


	//userLists := Get_all_users(in.LoggedInUserId)
	userLists,error := Storage_Get_all_users(in.LoggedInUserId)
	if error!=nil {
		return nil,error
	}
	var users []*pb.UserFollowed
	for _,eachUser := range userLists.List{
		a:= &pb.UserFollowed{UserName:eachUser.UserName, Isfollowed:eachUser.Isfollowed}
		users = append(users, a)
	}

	return &pb.GetUsersResponse{Users:users, NextPage:userLists.NextPage },nil
}


func (a *StorageServerImpl) FollowUser(ctx context.Context, in *pb.FollowUserRequest) (*pb.FollowUserResponse, error) {

	//FollowUser(globals.User{in.UserName}, in.UserNames)
	error :=StorageFollowUser(globals.User{in.UserName}, in.UserNames)
	if error!=nil {
		return nil,error
	}

	return &pb.FollowUserResponse{Status:true},nil
}
func (a *StorageServerImpl) InsertUser(ctx context.Context, in *pb.InsertUserRequest) (*pb.InsertUserResponse, error) {

	fmt.Println("inside insert user")
	status,error := InsertUserRecord(in.UserName, in.Password)
	if error!=nil {
		return nil,error
	}
	fmt.Println("status of Insert user record: ", status)

	/*fmt.Println("status of Insert user record: ", status)
	isSuccess := InsertUser(in.UserName,in.Password)
*/
	return &pb.InsertUserResponse{Success:status},nil
}
func (a *StorageServerImpl) UserExist(ctx context.Context, in *pb.UserExistRequest) (*pb.UserExistResponse, error) {

	//isSuccess := UserExist(in.UserName)
	isSuccess,error := CheckUserExist(in.UserName)
	if error!=nil {
		return nil,error
	}
	return &pb.UserExistResponse{Success:isSuccess},nil
}


func NewStorageServer(cfg *Config) (error) {
	InitGlobals()
	lis, err := net.Listen(TCP, cfg.HTTPAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterStorageServer(server, &StorageServerImpl{})
	// Register reflection service on gRPC server.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}

func (w *StorageServer) Start() error {
	return w.srv.ListenAndServe()
}


func (w *StorageServer) Shutdown(ctx context.Context) error {
	return w.srv.Shutdown(ctx)
}
