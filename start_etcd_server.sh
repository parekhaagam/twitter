#!/usr/bin/env bash
echo "Starting etcd server"

comm="$GOPATH/bin/goreman"
commarg="-f ./Procfile start"

comm="$comm $commarg"
echo $comm
nohup $comm > raft.log &