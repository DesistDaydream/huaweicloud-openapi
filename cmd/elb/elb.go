package elb

import (
	"os"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/config"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/elb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ElbClient *elb.ElbClient

func CreateCommand() *cobra.Command {
	elbCmd := &cobra.Command{
		Use:              "elb",
		Short:            "控制 ELB 资源",
		PersistentPreRun: elbPersistentPreRun,
	}

	elbCmd.AddCommand(
		CreateIPGroupCommand(),
	)

	return elbCmd
}

func elbPersistentPreRun(cmd *cobra.Command, args []string) {

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
	client, err := elb.NewElbClient(
		auth.AuthList[UserName].AccessKeyID,
		auth.AuthList[UserName].SecretAccessKey,
		Region,
	)
	if err != nil {
		panic(err)
	}

	ElbClient = client
}
