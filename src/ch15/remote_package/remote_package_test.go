package remote_package

/*
1. 通过 go get 来获取远程依赖
   go get -u 强制从网络更新远程依赖
2. 注意代码在github上的组织形式，以适应go get
	直接以代码路径开始，不要有src
*/

import (
	"testing"
	// cm为别名，包名太长
	cm "github.com/easierway/concurrent_map"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
