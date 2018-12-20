#!/usr/bin/env bash
echo "------Starting etcd Server------"
comm="$GOPATH/bin/goreman"
commarg="-f ./Procfile start"
comm="$comm $commarg"
nohup $comm > raft.log &


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