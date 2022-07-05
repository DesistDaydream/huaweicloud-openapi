/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/DesistDaydream/huaweicloud-openapi/cmd/ecs"
	"github.com/DesistDaydream/huaweicloud-openapi/cmd/elb"
	"github.com/DesistDaydream/huaweicloud-openapi/cmd/vpc"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/config"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/huaweiclient"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// logLevel := pflag.String("log-level", "info", "The logging level:[debug, info, warn, error, fatal]")
	// logFile := pflag.String("log-output", "", "the file which log to, default stdout")
	// logFormat := pflag.String("log-format", "text", "日志输出格式，可选值：text, json")

	// // 初始化日志
	// if err := logging.LogInit(*logLevel, *logFile, *logFormat); err != nil {
	// 	logrus.Fatal("set log level error")
	// }

	app := newApp()
	err := app.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newApp() *cobra.Command {
	long := `对 huaweicloud-openapi 工具的长描述，包含用例等，比如:
https://developer.huaweicloud.com/openapilist
API 在线调试：https://apiexplorer.developer.huaweicloud.com/apiexplorer/overview`

	var RootCmd = &cobra.Command{
		Use:              "huaweicloud-openapi",
		Short:            "通过华为云 OpenAPI 管理资源的工具",
		Long:             long,
		PersistentPreRun: rootPersistentPreRun,
	}

	RootCmd.PersistentFlags().StringP("log-level", "", "info", "日志级别:[debug, info, warn, error, fatal]")
	RootCmd.PersistentFlags().StringP("log-output", "", "", "日志输出位置，不填默认标准输出 stdout")
	RootCmd.PersistentFlags().StringP("log-format", "", "text", "日志输出格式: [text, json]")

	RootCmd.PersistentFlags().StringP("auth-file", "f", "auth.yaml", "认证信息文件")
	RootCmd.PersistentFlags().StringP("username", "u", "", "用户名")
	RootCmd.PersistentFlags().StringP("region", "r", "cn-southwest-2", "地域")

	// 添加子命令
	RootCmd.AddCommand(
		vpc.CreateCommand(),
		elb.CreateCommand(),
		ecs.CreateCommand(),
	)

	return RootCmd
}

// 执行每个 root 下的子命令时，都需要执行的函数
func rootPersistentPreRun(cmd *cobra.Command, args []string) {
	// 初始化日志
	LogLevel, _ := cmd.Flags().GetString("log-level")
	LogOutput, _ := cmd.Flags().GetString("log-output")
	LogFormat, _ := cmd.Flags().GetString("log-format")
	if err := logging.LogInit(LogLevel, LogOutput, LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	// 认证信息文件处理的相关逻辑
	AuthFile, _ := cmd.Flags().GetString("auth-file")
	UserName, err := cmd.Flags().GetString("username")
	if err != nil {
		logrus.Fatalln("请指定用户名")
	}
	region, _ := cmd.Flags().GetString("region")

	// 检查 clientFlags.AuthFile 文件是否存在
	if _, err := os.Stat(AuthFile); os.IsNotExist(err) {
		logrus.Fatal("文件不存在")
	}
	// 获取认证信息
	auth := config.NewAuthInfo(AuthFile)

	// 判断传入的域名是否存在在认证信息中
	if !auth.IsUserExist(UserName) {
		logrus.Fatalf("认证信息中不存在 %v 用户, 请检查认证信息文件或命令行参数的值", UserName)
	}

	huaweiclient.Info = &huaweiclient.ClientInfo{
		AK:     auth.AuthList[UserName].AccessKeyID,
		SK:     auth.AuthList[UserName].SecretAccessKey,
		Region: region,
	}
}
