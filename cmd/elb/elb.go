package elb

import (
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/elb"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/huaweiclient"
	"github.com/spf13/cobra"
)

var (
	elbClient *elb.ElbClient
)

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
	// 执行根命令的初始化操作
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}

	// 初始化账号Client
	client, err := elb.NewElbClient(
		huaweiclient.Info.AK,
		huaweiclient.Info.SK,
		huaweiclient.Info.Region,
	)
	if err != nil {
		panic(err)
	}

	elbClient = client
}
