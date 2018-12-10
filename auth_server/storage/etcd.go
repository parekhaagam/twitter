package storage

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/etcd-io/etcd/clientv3/namespace"
	"strings"
)

func Get(){

	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"localhost:2379"}})
	if err != nil {
		// handle error!
	}

	unprefixedKV := cli.KV
	var loggedInUserMapKey []string
	loggedInUserMapKey = append(loggedInUserMapKey,LOGGED_IN_USER_PREFIX,"123uname")
	cli.KV = namespace.NewKV(cli.KV, AUTH_PREFIX)
	cli.Watcher = namespace.NewWatcher(cli.Watcher, AUTH_PREFIX)
	cli.Lease = namespace.NewLease(cli.Lease, AUTH_PREFIX)

	cli.Put(context.TODO(), strings.Join(loggedInUserMapKey,""), "123")
	resp, _ := unprefixedKV.Get(context.TODO(), strings.Join(loggedInUserMapKey,""))
	fmt.Printf("%s\n", resp.Kvs[0].Value)
	// Output: 123

	unprefixedKV.Put(context.TODO(), "my-prefix/abc", "456")
	resp, _ = cli.Get(context.TODO(), "abc")
	fmt.Printf("%s\n", resp.Kvs[0].Value)
	// Output: 456
}
