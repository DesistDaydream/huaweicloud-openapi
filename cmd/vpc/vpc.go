package vpc

import (
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/huaweiclient"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/vpc"
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
	// 执行根命令的初始化操作
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}

	// 初始化账号Client
	client, err := vpc.NewVpcClient(
		huaweiclient.Info.AK,
		huaweiclient.Info.SK,
		huaweiclient.Info.Region,
	)
	if err != nil {
		panic(err)
	}

	VpcClient = client
}
