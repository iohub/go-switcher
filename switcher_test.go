package switcher

import (
	"testing"
	"time"

	"github.com/iohub/go-switcher/store/etcd"
	"go.etcd.io/etcd/clientv3"
)

func TestSwitcher(t *testing.T) {
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Duration(5) * time.Millisecond,
	}
	etcd.Init(config)
}
