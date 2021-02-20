package etcd

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/etcd/clientv3"
)

func TestEndPoint(t *testing.T) {

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Duration(5) * time.Millisecond,
	}
	Init(config)
	ok := IsEnable(context.TODO(), "test_enable_true", false)
	assert.Equal(t, true, ok)
	ok = IsEnable(context.TODO(), "test_enable_false", false)
	assert.Equal(t, false, ok)
}
