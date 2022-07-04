package vpc

import (
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/vpc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CreateSecurityGroupCommand() *cobra.Command {
	long := `操作 VPC 下的安全组资源：
	list 列出所有安全组及其规则
	show 列出指定安全组的规则`

	var SecurityGroupCmd = &cobra.Command{
		Use:   "securityGroup",
		Short: "操作 VPC 下的安全组资源",
		Long:  long,
		Run:   runSecurityGroup,
	}

	SecurityGroupCmd.Flags().StringP("operation", "o", "list", "操作类型，详见命令描述")
	SecurityGroupCmd.Flags().StringP("excel", "e", "security_group_for_vpc.xlsx", "存有安全组的文件")
	SecurityGroupCmd.Flags().StringP("security-group-name", "n", "default", "安全组名称")
	// 功能说明：是否只预检此次请求
	// 取值范围：
	// -true：发送检查请求，不会更新安全组内容。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。
	// -false（默认值）：发送正常请
	SecurityGroupCmd.Flags().BoolP("pre-check", "p", false, "是否只预检此次请求")

	return SecurityGroupCmd
}

func runSecurityGroup(cmd *cobra.Command, args []string) {
	// 获取全部命令行标志
	operation, _ := cmd.Flags().GetString("operation")
	// ipsFile, _ := cmd.Flags().GetString("excel")
	securityGroupName, _ := cmd.Flags().GetString("security-group-name")
	// dryRun, _ := cmd.Flags().GetBool("dry-run")

	// 实例化 VPC 的 API 处理器
	v := vpc.NewVpcSecurityGroup(VpcClient)

	// 执行操作
	switch operation {
	case "list":
		v.ListAllSecurityGroupAndRules()
	case "show":
		v.ListSecurityGroupRules(securityGroupName)
	case "update":
		// TODO
	default:
		logrus.Fatalln("操作类型不存在，请使用 -o 指定操作类型")
	}
}
