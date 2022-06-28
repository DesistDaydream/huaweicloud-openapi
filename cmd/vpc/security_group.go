package vpc

import "github.com/spf13/cobra"

func CreateSecurityGroupCommand() *cobra.Command {
	var SecurityGroupCmd = &cobra.Command{
		Use:   "securityGroup",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples`,
		Run: func(cmd *cobra.Command, args []string) {
			runSecurityGroup(cmd, args)
		},
		Args: cobra.NoArgs,
	}

	SecurityGroupCmd.PersistentFlags().StringP("operation", "o", "list", "操作类型: [list, update]")
	SecurityGroupCmd.PersistentFlags().StringP("excel", "e", "security_group_for_vpc.xlsx", "存有安全组的文件")
	SecurityGroupCmd.PersistentFlags().StringP("security-group-name", "n", "测试安全组", "安全组名称")
	// 功能说明：是否只预检此次请求
	// 取值范围：
	// -true：发送检查请求，不会更新安全组内容。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。
	// -false（默认值）：发送正常请
	SecurityGroupCmd.PersistentFlags().BoolP("pre-check", "p", false, "是否只预检此次请求")

	return SecurityGroupCmd
}

func runSecurityGroup(cmd *cobra.Command, args []string) {

}
