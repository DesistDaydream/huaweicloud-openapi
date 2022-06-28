package vpc

import (
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	VpcCmd := &cobra.Command{
		Use:   "vpc",
		Short: "控制 VPC 资源",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	VpcCmd.AddCommand(
		CreateIPGroupCommand(),
		CreateSecurityGroupCommand(),
	)

	return VpcCmd
}
