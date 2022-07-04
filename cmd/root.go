/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/DesistDaydream/huaweicloud-openapi/cmd/ecs"
	"github.com/DesistDaydream/huaweicloud-openapi/cmd/elb"
	"github.com/DesistDaydream/huaweicloud-openapi/cmd/vpc"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
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

	var rootCmd = &cobra.Command{
		Use:              "huaweicloud-openapi",
		Short:            "通过华为云 OpenAPI 管理资源的工具",
		Long:             long,
		PersistentPreRun: rootPersistentPreRun,
	}

	rootCmd.PersistentFlags().StringP("log-level", "", "info", "日志级别:[debug, info, warn, error, fatal]")
	rootCmd.PersistentFlags().StringP("log-output", "", "", "日志输出位置，不填默认标准输出 stdout")
	rootCmd.PersistentFlags().StringP("log-format", "", "text", "日志输出格式: [text, json]")

	rootCmd.PersistentFlags().StringP("auth-file", "f", "auth.yaml", "认证信息文件")
	rootCmd.PersistentFlags().StringP("username", "u", "", "用户名")
	rootCmd.PersistentFlags().StringP("region", "r", "cn-southwest-2", "地域")

	// 添加子命令
	rootCmd.AddCommand(
		vpc.CreateCommand(),
		elb.CreateCommand(),
		ecs.CreateCommand(),
	)

	return rootCmd
}

func rootPersistentPreRun(cmd *cobra.Command, args []string) {
	// TODO: 这里为什么不执行？Cobra 的 Bug？还是还有别的机制需要触发这里的逻辑
	LogLevel, _ := cmd.Flags().GetString("log-level")
	LogOutput, _ := cmd.Flags().GetString("log-output")
	LogFormat, _ := cmd.Flags().GetString("log-format")
	if err := logging.LogInit(LogLevel, LogOutput, LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
		// TODO: 认证信息文件处理的相关逻辑写在这里
	}
}
