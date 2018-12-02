package controllers

import (
	"context"
	"fmt"
	pb "github.com/parekhaagam/twitter/contracts/storage"
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
	tid := InsertTweets(user, in.Content)
	return &pb.InsertTweetReply{TID:tid},nil
	}


func (a *StorageServerImpl)GetFollowersTweets(ctx context.Context, in *pb.GetFollowersRequest)(*pb.GetFollowersReply, error){

	//followings []globals.User :=

	currUser := globals.User{in.User.UserName}
	following := GetAllFollowing(currUser)
	followingNumber := len(following)
	following = append(following, currUser)
	tweets := GetFollowersTweets(following)
	fmt.Println(tweets)

	var tweetArray []*pb.Tweet
	for _,t := range tweets{
		tTemp := pb.Tweet{UserId:t.UserId, TID:t.TID, Timestamp:t.Timestamp, Content:t.Content, TimeMessage:t.TimeMessage}
		tweetArray = append(tweetArray, &tTemp)
	}
	return &pb.GetFollowersReply{Tweets: tweetArray,FollowingNumber:strconv.Itoa(followingNumber)},nil
}

func (a *StorageServerImpl) GetAllUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error){


	userLists := Get_all_users(in.LoggedInUserId)

	var users []*pb.UserFollowed
	for _,eachUser := range userLists.List{
		a:= &pb.UserFollowed{UserName:eachUser.UserName, Isfollowed:eachUser.Isfollowed}
		users = append(users, a)
	}

	return &pb.GetUsersResponse{Users:users, NextPage:userLists.NextPage },nil
}


func (a *StorageServerImpl) FollowUser(ctx context.Context, in *pb.FollowUserRequest) (*pb.FollowUserResponse, error) {

	FollowUser(globals.User{in.UserName}, in.UserNames)
	return &pb.FollowUserResponse{Status:true},nil
}


func NewStorageServer(cfg *Config) (error) {
	globals.InitGlobals()
	lis, err := net.Listen("tcp", cfg.HTTPAddr)
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

