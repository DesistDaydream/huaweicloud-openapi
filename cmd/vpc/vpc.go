package vpc

import (
	"os"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/config"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/logging"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/vpc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var VpcClient *vpc.VpcClient

func CreateCommand() *cobra.Command {
	VpcCmd := &cobra.Command{
		Use:              "vpc",
		Short:            "控制 VPC 资源",
		PersistentPreRun: vpcPersistentPreRun,
	}

	VpcCmd.AddCommand(
		CreateIPGroupCommand(),
		CreateSecurityGroupCommand(),
	)

	return VpcCmd
}

func vpcPersistentPreRun(cmd *cobra.Command, args []string) {
	LogLevel, _ := cmd.Flags().GetString("log-level")
	LogOutput, _ := cmd.Flags().GetString("log-output")
	LogFormat, _ := cmd.Flags().GetString("log-format")
	if err := logging.LogInit(LogLevel, LogOutput, LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	authFile, _ := cmd.Flags().GetString("auth-file")
	userName, err := cmd.Flags().GetString("username")
	if err != nil {
		logrus.Fatalln("请指定用户名")
	}
	region, _ := cmd.Flags().GetString("region")

	// 检查 clientFlags.AuthFile 文件是否存在
	if _, err := os.Stat(authFile); os.IsNotExist(err) {
		logrus.Fatal("文件不存在")
	}
	// 获取认证信息
	auth := config.NewAuthInfo(authFile)

	// 判断传入的域名是否存在在认证信息中
	if !auth.IsUserExist(userName) {
		logrus.Fatalf("认证信息中不存在 %v 用户, 请检查认证信息文件或命令行参数的值", userName)
	}

	// 初始化账号Client
	client, err := vpc.NewVpcClient(
		auth.AuthList[userName].AccessKeyID,
		auth.AuthList[userName].SecretAccessKey,
		region,
	)
	if err != nil {
		panic(err)
	}

	VpcClient = client
}
