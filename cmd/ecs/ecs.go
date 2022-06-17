package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/config"
	hwcecs "github.com/DesistDaydream/huaweicloud-openapi/pkg/ecs"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/ecs/lifecycle"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/huaweiclient"
)

// LogInit 日志功能初始化，若指定了 log-output 命令行标志，则将日志写入到文件中
func LogInit(level, file, format string) error {
	switch format {
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05",
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			// FieldMap:          map[logrus.fieldKey]string{},
			// CallerPrettyfier: func(*runtime.Frame) (string, string) {},
			PrettyPrint: false,
		})
	}

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(logLevel)

	if file != "" {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	}

	return nil
}

func main() {
	// operation := pflag.StringP("operation", "o", "", "操作类型: [add, list, batch]")
	logLevel := pflag.String("log-level", "info", "日志级别:[debug, info, warn, error, fatal]")
	logFile := pflag.String("log-output", "", "日志输出位置，不填默认标准输出 stdout")
	logFormat := pflag.String("log-format", "text", "日志输出格式: [text, json]")

	authFile := pflag.StringP("auth-file", "f", "auth.yaml", "认证信息文件")
	userName := pflag.StringP("username", "u", "", "认证信息文件")
	// 添加命令行标志
	pflag.Parse()

	// 初始化日志
	if err := LogInit(*logLevel, *logFile, *logFormat); err != nil {
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
