package ecs

import (
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/ecs/lifecycle"
	"github.com/spf13/cobra"
)

func CreateLifecycleCommand() *cobra.Command {
	LifecycleCmd := &cobra.Command{
		Use:   "lifecycle",
		Short: "控制 ECS 资源",
		Run:   runLifecycle,
	}

	return LifecycleCmd
}

func runLifecycle(cmd *cobra.Command, args []string) {
	// 执行操作
	id := "cd2d5bf4-ba5a-4abc-a62f-cd265c754e87"
	l := lifecycle.NewEcsLifecycle(EcsClient)
	l.ShowServer(id)
}
