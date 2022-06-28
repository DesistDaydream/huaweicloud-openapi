package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/config"
	hwcecs "github.com/DesistDaydream/huaweicloud-openapi/pkg/ecs"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/ecs/lifecycle"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/huaweiclient"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/logging"
)

func main() {
	// 添加命令行标志
	logFlags := logging.LoggingFlags{}
	clientFlags := huaweiclient.ClientFlags{}
	logFlags.AddFlags()
	clientFlags.AddFlags()
	pflag.Parse()

	// 初始化日志
	if err := logging.LogInit(logFlags.LogLevel, logFlags.LogOutput, logFlags.LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	// 获取认证信息
	auth := config.NewAuthInfo(clientFlags.AuthFile)
	// 判断传入的域名是否存在在认证信息中
	if !auth.IsUserExist(clientFlags.UserName) {
		logrus.Fatalf("认证信息中不存在 %v 用户, 请检查认证信息文件或命令行参数的值", clientFlags.UserName)
	}

	// 初始化账号Client
	client, err := huaweiclient.CreateEcsClient(
		auth.AuthList[clientFlags.UserName].AccessKeyID,
		auth.AuthList[clientFlags.UserName].SecretAccessKey,
		clientFlags.Region,
	)
	if err != nil {
		panic(err)
	}

	// 实例化各种 API 处理器
	h := hwcecs.NewEcsHandler(client)

	// 执行操作
	id := "cd2d5bf4-ba5a-4abc-a62f-cd265c754e87"
	l := lifecycle.NewEcsLifecycle(h)
	l.ShowServer(id)
}
