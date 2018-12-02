package auth

import (
	"context"
	"github.com/google/uuid"
	"github.com/parekhaagam/twitter/globals"
	"github.com/parekhaagam/twitter/web/auth/storage/memory"
	"google.golang.org/grpc"
	pb "github.com/parekhaagam/twitter/contracts/authentication"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"

	//"github.com/parekhaagam/twitter/web/controllers"
	"net/http"
)

type AuthServer struct {
	srv *http.Server

	//contentDaemon *contentd.Contentd
}
type AuthServerImpl struct{}
func (a *AuthServerImpl)GetToken(ctx context.Context, in *pb.GetTokenRequest) (*pb.GetTokenReply, error){
	token := getToken(in.Userid)
	return &pb.GetTokenReply{Token:token},nil
}

func (a *AuthServerImpl) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateReply, error){
	return &pb.AuthenticateReply{Success:true},nil
}
func NewAuthServer(cfg *Config) (error) {
	globals.InitGlobals()
	lis, err := net.Listen("tcp", cfg.HTTPAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAuthServer(server, &AuthServerImpl{})
	// Register reflection service on gRPC server.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}

func (w *AuthServer) Start() error {
	return w.srv.ListenAndServe()
}


func (w *AuthServer) Shutdown(ctx context.Context) error {
	return w.srv.Shutdown(ctx)
}
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
		return val.Token
	}

}