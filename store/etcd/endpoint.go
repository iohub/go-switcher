package etcd

import (
	"github.com/juju/errors"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
)

var (
	cli     *clientv3.Client
	kvStore clientv3.KV
	config  clientv3.Config
)

//	config = clientv3.Config{
//		Endpoints:   []string{"127.0.0.1:2379"},
//		DialTimeout: time.Duration(5) * time.Millisecond,
//	}

func Init(cfg clientv3.Config) error {
	var err error
	if cli, err = clientv3.New(cfg); err != nil {
		return errors.Annotatef(err, "etcd.Init() -> clientv3.New(), config:%+v", cfg)
	}
	kvStore = clientv3.NewKV(cli)
	return nil
}

func getVal(ctx context.Context, key string) (string, error) {
	rsp, err := kvStore.Get(ctx, key)
	if err != nil {
		return "", errors.Annotatef(err, "etcd.getVal() -> Get(), key:%v", key)
	}
	if len(rsp.Kvs) == 0 {
		return "", nil
	}
	return string(rsp.Kvs[0].Value), nil
}
