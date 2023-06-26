package elb

import (
	"github.com/spf13/cobra"
)

type IPGroupFlags struct {
	ipsFile       string
	addrGroupName string
	dryRun        bool
}

var (
	ipGroupFlags IPGroupFlags
)

func CreateIPGroupCommand() *cobra.Command {
	var ipGroupCmd = &cobra.Command{
		Use:   "ipGroup",
		Short: "管理 ELB 的 IP 地址组",
		Args:  cobra.NoArgs,
	}

	ipGroupCmd.PersistentFlags().StringVarP(&ipGroupFlags.ipsFile, "excel", "e", "ipaddr_group_for_elb.xlsx", "存有 IP 地址组的文件")
	ipGroupCmd.PersistentFlags().StringVarP(&ipGroupFlags.addrGroupName, "addr-group-name", "n", "测试地址组", "地址组名称")
	// 功能说明：是否只预检此次请求
	// 取值范围：
	// -true：发送检查请求，不会更新地址组内容。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。
	// -false（默认值）：发送正常请求，并直接更新地址组。
	ipGroupCmd.PersistentFlags().BoolVarP(&ipGroupFlags.dryRun, "dry-run", "d", false, "是否真实执行操作")

	ipGroupCmd.AddCommand(
		IPGroupListCommand(),
		IPGroupUpdateCommand(),
	)

	return ipGroupCmd
}
