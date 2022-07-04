package ecs

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/config"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/ecs"
)

var EcsClient *ecs.EcsClient

func CreateCommand() *cobra.Command {
	ecsCmd := &cobra.Command{
		Use:              "ecs",
		Short:            "控制 ECS 资源",
		PersistentPreRun: ecsPersistentPreRun,
	}
	ecsCmd.AddCommand(
		CreateLifecycleCommand(),
	)

	return ecsCmd
}

func ecsPersistentPreRun(cmd *cobra.Command, args []string) {
	AuthFile, _ := cmd.Flags().GetString("auth-file")
	UserName, err := cmd.Flags().GetString("username")
	if err != nil {
		logrus.Fatalln("请指定用户名")
	}
	Region, _ := cmd.Flags().GetString("region")

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

	// 初始化账号Client
	client, err := ecs.NewEcsClient(
		auth.AuthList[UserName].AccessKeyID,
		auth.AuthList[UserName].SecretAccessKey,
		Region,
	)
	if err != nil {
		panic(err)
	}

	EcsClient = client
}
