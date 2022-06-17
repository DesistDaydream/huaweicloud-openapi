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
	authFile := pflag.StringP("auth-file", "f", "auth.yaml", "认证信息文件")
	userName := pflag.StringP("username", "u", "", "用户名")
	// 添加命令行标志
	logFlags := logging.LoggingFlags{}
	logFlags.AddFlags()
	pflag.Parse()

	// 初始化日志
	if err := logging.LogInit(logFlags.LogLevel, logFlags.LogOutput, logFlags.LogFormat); err != nil {
		logrus.Fatal("set log level error")
	}

	// 获取认证信息
	auth := config.NewAuthInfo(*authFile)
	// 判断传入的域名是否存在在认证信息中
	if !auth.IsUserExist(*userName) {
		logrus.Fatalf("认证信息中不存在 %v 用户, 请检查认证信息文件或命令行参数的值", *userName)
	}

	// 初始化账号Client
	client, err := huaweiclient.CreateClient(auth.AuthList[*userName].AccessKeyID, auth.AuthList[*userName].SecretAccessKey)
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
