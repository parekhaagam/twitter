package auth_server

import (
	"context"
	"github.com/parekhaagam/twitter/auth_server/storage"
	pb "github.com/parekhaagam/twitter/web_server/contracts/authentication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	//"github.com/parekhaagam/twitter/web_server/controllers"
	"net/http"
)

type AuthServer struct {
	srv *http.Server

	//contentDaemon *contentd.Contentd
}
type AuthServerImpl struct{}
func (a *AuthServerImpl)GetToken(ctx context.Context, in *pb.GetTokenRequest) (*pb.GetTokenReply, error){
	token,err := storage.GetOrCreateToken(in.Userid)
	if err!=nil {
		return nil,err
	}
	return &pb.GetTokenReply{Token:token},nil
}

func (a *AuthServerImpl) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateReply, error){
	isTokenAvailable,err := authenticate(in.Token)
	if err != nil {
		return nil,err
	}
	return &pb.AuthenticateReply{Success:isTokenAvailable},nil
}
func NewAuthServer(cfg *Config) (error) {
	//globals.InitGlobals()
	lis, err := net.Listen(TCP, cfg.HTTPAddr)
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
/*func getToken(userId string) (string)  {
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
}*/
func authenticate(token string) (bool,error) {
	return storage.IsTokenValid(token)
}