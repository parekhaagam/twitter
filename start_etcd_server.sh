
#mkdir $ACCESSIBLE_PATH/data_dir
chmod 777 $ACCESSIBLE_PATH/data_dir
 #export ETCD_DATA_DIR=$ACCESSIBLE_PATH/data_dir
#mkdir go_path_sw
chmod 777 $ACCESSIBLE_PATH/go_path_sw
export GOPATH=$ACCESSIBLE_PATH/go_path_sw
echo "goreman"
#go get github.com/mattn/goreman
echo "etcd"
#go get github.com/etcd-io/etcd
echo "starting"
comm="$GOPATH/bin/goreman"
commarg="-f $GOPATH/src/go.etcd.io/etcd/Procfile start"
nohup start-stop-daemon --start --make-pidfile --pidfile abc.pid --exec $comm -- $commarg > raft.log &