#!/usr/bin/env bash
echo "Stopping etcd server"
comm="$GOPATH/bin/goreman run stop"
commarg="etcd"
numOfServers=3

until [ $numOfServers -lt 1 ]; do
    command="$comm $commarg$numOfServers"
    echo $command
    nohup $command > raft_stop.log &
    let numOfServers-=1
done