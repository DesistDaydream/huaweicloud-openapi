package vpc

import (
	"github.com/spf13/cobra"
)

type SecurityGroupCmdFlags struct {
	sgID     string
	excel    string
	sgName   string
	preCheck bool
}

var sgFlags SecurityGroupCmdFlags

func CreateSecurityGroupCommand() *cobra.Command {
	long := `操作 VPC 下的安全组资源：
	list 列出所有安全组及其规则
	show 列出指定安全组的规则`

	var securityGroupCmd = &cobra.Command{
		Use:   "securityGroup",
		Short: "操作 VPC 下的安全组资源",
		Long:  long,
	}
	securityGroupCmd.PersistentFlags().StringVarP(&sgFlags.sgID, "sg-id", "", "", "安全组 ID")
	securityGroupCmd.PersistentFlags().StringVarP(&sgFlags.excel, "excel", "e", "security_group_for_vpc.xlsx", "存有安全组的文件")
	securityGroupCmd.PersistentFlags().StringVarP(&sgFlags.sgName, "sg-name", "n", "default", "安全组名称")
	// 功能说明：是否只预检此次请求
	// 取值范围：
	// -true：发送检查请求，不会更新安全组内容。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。
	// -false（默认值）：发送正常请
	securityGroupCmd.PersistentFlags().BoolVarP(&sgFlags.preCheck, "pre-check", "p", false, "是否只预检此次请求")

	securityGroupCmd.AddCommand(
		CreateSecurityGroupListCmd(),
		CreateSecurityGroupRulesListCmd(),
		CreateSecurityGroupRulesUpdateCmd(),
	)

	return securityGroupCmd
}
