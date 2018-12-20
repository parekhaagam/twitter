#!/usr/bin/env bash

echo "--------Stopping Web Server---------"
ps -ef | grep go-build.*/exe/web | grep -v grep | awk '{print $2}' | xargs kill

echo "--------Stopping Storage Server---------"
ps -ef | grep go-build.*/exe/storage | grep -v grep | awk '{print $2}' | xargs kill

echo "--------Stopping Authentication Server---------"
ps -ef | grep go-build.*/exe/auth | grep -v grep | awk '{print $2}' | xargs kill

echo "--------Stopping Raft Server---------"
ps -ef | grep .*bin/etcd | grep -v grep | awk '{print $2}' | xargs kill