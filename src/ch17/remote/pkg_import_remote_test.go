package remote

import (
	"testing"

	cm "github.com/easierway/concurrent_map" // cm 是起表名，给引入的包起别名。
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
