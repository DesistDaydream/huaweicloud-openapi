package vpc

import (
	"github.com/spf13/cobra"
)

func CreateSecurityGroupRulesUpdateCmd() *cobra.Command {
	long := ``

	var securityGroupRulesUpdateCmd = &cobra.Command{
		Use:   "rulesUpdate",
		Short: "更新安全组的规则",
		Long:  long,
		Run:   runSecurityGroupRulesUpdate,
	}

	return securityGroupRulesUpdateCmd
}

func runSecurityGroupRulesUpdate(cmd *cobra.Command, args []string) {
	// TODO: 更新安全组规则
}
