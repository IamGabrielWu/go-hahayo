package depend_test

// # Go 未解决的依赖问题
// 1. 同一环境下， 不同项目使用同一包的不同版本
// 2. 无法管理对包的特定版本的依赖
// # vender 路径
// 随着Go 1.5 release 版本的发布， vendor 目录被添加到除了GOPATH和GOROOT之外的依赖项目查找的解决方案。 在Go 1.6之前，你需要手动设置环境变量。
// 查找依赖包路径的解决方案如下：
// 1. 当前包下vendor目录
// 2. 向上级目录查找，知道找到src 下的vendor 目录
// 3. 在GOPATH下面查找依赖包
// 4. 在GOROOT目录查找依赖包
// # 常用的依赖管理工具
// godep https://github.com/tools/godep
// glide https://github.com/Masterminds/glide
// dep https://github.com/golang/dep

import (
	"testing"

	cm "github.com/easierway/concurrent_map" // cm 是起表名，给引入的包起别名。
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
