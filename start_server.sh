#!/usr/bin/env bash
echo "------Starting Authentication Server------"
nohup go run cmd/main/auth.go > authServer.log &
echo "------Starting Storage Server------"
nohup go run cmd/main/storage.go > storageServer.log  &
echo "------Starting Web Server------"
nohup go run cmd/main/web.go > webServer.log &