package huaweiclient

import (
	"github.com/spf13/pflag"
)

type ClientInfo struct {
	AK     string
	SK     string
	Region string
}

// 用于在根命令的 PersistentPreRun 中初始化账号Client
// 以便在子命令中直接引用该变量的值
var Info *ClientInfo

// 创建客户端命令行标志
type ClientFlags struct {
	AuthFile string
	UserName string
	Region   string
}

// 添加命令行标志
func (c *ClientFlags) AddFlags() {
	pflag.StringVarP(&c.AuthFile, "auth-file", "f", "auth.yaml", "认证信息文件")
	pflag.StringVarP(&c.UserName, "username", "u", "", "用户名")
	pflag.StringVarP(&c.Region, "region", "r", "cn-southwest-2", "地域")
}
