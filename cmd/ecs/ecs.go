package ecs

import (
	"github.com/spf13/cobra"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/ecs"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/huaweiclient"
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
	// 执行根命令的初始化操作
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}

	// 初始化账号Client
	client, err := ecs.NewEcsClient(
		huaweiclient.Info.AK,
		huaweiclient.Info.SK,
		huaweiclient.Info.Region,
	)
	if err != nil {
		panic(err)
	}

	EcsClient = client
}
