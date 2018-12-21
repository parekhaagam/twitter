#!/usr/bin/env bash
echo "------go getting etcd------"
go get github.com/etcd-io/etcd
echo "------go getting gorman------"
go get github.com/mattn/goreman
echo "------Starting etcd Server------"
echo $GOPATH
comm="$GOPATH/bin/goreman"
commarg="-f ./Procfile start"
comm="$comm $commarg"
nohup $GOPATH/bin/goreman -f ./Procfile start  > raft.log &


echo "------Building Authentication Server------"
go build cmd/main/auth.go > authServer.log

echo "------Building Storage Server------"
go build cmd/main/storage.go > storageServer.log

echo "------Building Web Server------"
go build cmd/main/web.go > webServer.log

echo "------Starting Authentication Server------"
nohup go run cmd/main/auth.go > authServer.log &

echo "------Starting Storage Server------"
nohup go run cmd/main/storage.go > storageServer.log  &

echo "------Starting Web Server------"
nohup go run cmd/main/web.go > webServer.log &