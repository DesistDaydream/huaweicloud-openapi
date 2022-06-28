/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/DesistDaydream/huaweicloud-openapi/cmd/elb"
	"github.com/DesistDaydream/huaweicloud-openapi/cmd/vpc"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/huaweiclient"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	newApp()
}

func newApp() {
	long := `对 huaweicloud-openapi 工具的长描述，包含用例等，比如:
https://developer.huaweicloud.com/openapilist
API 在线调试：https://apiexplorer.developer.huaweicloud.com/apiexplorer/overview`

	var rootCmd = &cobra.Command{
		Use:   "huaweicloud-openapi",
		Short: "通过华为云 OpenAPI 管理资源的工具",
		Long:  long,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// 获取日志命令行标志
			LogLevel, _ := cmd.Flags().GetString("log-level")
			LogOutput, _ := cmd.Flags().GetString("log-output")
			LogFormat, _ := cmd.Flags().GetString("log-format")
			// 初始化日志
			if err := logging.LogInit(LogLevel, LogOutput, LogFormat); err != nil {
				logrus.Fatal("初始化日志失败", err)
			}
		},
	}

	clientFlags := huaweiclient.ClientFlags{}
	clientFlags.AddFlags()
	logFlags := logging.LoggingFlags{}
	logFlags.AddFlags()

	// 添加子命令
	rootCmd.AddCommand(
		vpc.CreateCommand(),
		elb.CreateCommand(),
		// ecs.CreateCommand(),
	)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
