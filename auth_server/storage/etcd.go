package storage

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/etcd-io/etcd/clientv3/namespace"
)

func Get(){

	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"localhost:2379"}})
	if err != nil {
		// handle error!
	}

	unprefixedKV := cli.KV
	cli.KV = namespace.NewKV(cli.KV, "my-prefix/")
	cli.Watcher = namespace.NewWatcher(cli.Watcher, "my-prefix/")
	cli.Lease = namespace.NewLease(cli.Lease, "my-prefix/")

	cli.Put(context.TODO(), "abc", "123")
	resp, _ := unprefixedKV.Get(context.TODO(), "my-prefix/abc")
	fmt.Printf("%s\n", resp.Kvs[0].Value)
	// Output: 123

	unprefixedKV.Put(context.TODO(), "my-prefix/abc", "456")
	resp, _ = cli.Get(context.TODO(), "abc")
	fmt.Printf("%s\n", resp.Kvs[0].Value)
	// Output: 456
}
