package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/config"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/fileparse"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/huaweiclient"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/logging"
	hwcvpc "github.com/DesistDaydream/huaweicloud-openapi/pkg/vpc"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/vpc/ipaddressgroup"
)

func main() {
	// 添加命令行标志
	operation := pflag.StringP("operation", "o", "", "操作类型: [list, update]")
	ipsFile := pflag.StringP("excel", "e", "ipaddr_group.xlsx", "存有 IP 地址组的文件")
	addrGroupName := pflag.StringP("addr-group-name", "n", "测试地址组", "地址组名称")
	// 功能说明：是否只预检此次请求
	// 取值范围：
	// -true：发送检查请求，不会更新地址组内容。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。
	// -false（默认值）：发送正常请求，并直接更新地址组。
	dryRun := pflag.BoolP("dry-run", "r", false, "是否真实执行操作")

	logFlags := logging.LoggingFlags{}
	clientFlags := huaweiclient.ClientFlags{}
	logFlags.AddFlags()
	clientFlags.AddFlags()
	pflag.Parse()

	// 初始化日志
	if err := logging.LogInit(logFlags.LogLevel, logFlags.LogOutput, logFlags.LogFormat); err != nil {
		logrus.Fatal("set log level error")
	}

	// 检查 clientFlags.AuthFile 文件是否存在
	if _, err := os.Stat(clientFlags.AuthFile); os.IsNotExist(err) {
		logrus.Fatal("文件不存在")
	}
	// 获取认证信息
	auth := config.NewAuthInfo(clientFlags.AuthFile)

	// 判断传入的域名是否存在在认证信息中
	if !auth.IsUserExist(clientFlags.UserName) {
		logrus.Fatalf("认证信息中不存在 %v 用户, 请检查认证信息文件或命令行参数的值", clientFlags.UserName)
	}

	// 初始化账号Client
	client, err := huaweiclient.CreateVpcClient(auth.AuthList[clientFlags.UserName].AccessKeyID, auth.AuthList[clientFlags.UserName].SecretAccessKey)
	if err != nil {
		panic(err)
	}

	// 实例化各种 API 处理器
	h := hwcvpc.NewVpcHandler(client)

	v := ipaddressgroup.NewVpcIPADdressGroup(h)
	// 执行操作
	switch *operation {
	case "list":
		v.ListAddressGroup()
	case "update":
		ipset, id, err := fileparse.GetVpcIPaddrGroup(*ipsFile, *addrGroupName)
		if err != nil {
			logrus.Fatalf("解析文件失败: %v", err)
		}
		v.UpdateAddressGroup(*addrGroupName, id, ipset, *dryRun)
	default:
		logrus.Fatalln("操作类型不存在，请使用 -o 指定操作类型")
	}

}
