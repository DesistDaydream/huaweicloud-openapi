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
		Use:   "elb",
		Short: "控制 ELB 资源",
	}

	cobra.OnInitialize(initELB)

	elbCmd.AddCommand(
		CreateIPGroupCommand(),
	)

	return elbCmd
}

func initELB() {
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
