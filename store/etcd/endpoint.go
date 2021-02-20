package etcd

import (
	"strconv"
	"strings"

	"github.com/juju/errors"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
)

var (
	cli      *clientv3.Client
	kvStore  clientv3.KV
	config   clientv3.Config
	inited   bool
	notFound = errors.New("not found")
)

func Init(cfg clientv3.Config) error {
	var err error
	if cli, err = clientv3.New(cfg); err != nil {
		return errors.Annotatef(err, "etcd.Init() -> clientv3.New(), config:%+v", cfg)
	}
	kvStore = clientv3.NewKV(cli)
	inited = true
	return nil
}

func getVal(ctx context.Context, key string) (string, error) {
	rsp, err := kvStore.Get(ctx, key)
	if err != nil {
		return "", errors.Annotatef(err, "etcd.getVal() -> Get(), key:%v", key)
	}
	if len(rsp.Kvs) == 0 {
		return "", notFound
	}
	return string(rsp.Kvs[0].Value), nil
}

func IsEnable(ctx context.Context, key string, defaultVal bool) bool {
	if !inited {
		panic("etcd endpoint not inited")
	}
	val, err := getVal(ctx, key)
	if err == notFound {
		return defaultVal
	}
	val = strings.ToLower(val)
	b, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}
	return b
}
