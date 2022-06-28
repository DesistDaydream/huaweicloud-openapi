package elb

import (
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	elbCmd := &cobra.Command{
		Use:   "elb",
		Short: "控制 ELB 资源",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	elbCmd.AddCommand(
		CreateIPGroupCommand(),
	)

	return elbCmd
}
