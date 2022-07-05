package elb

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/elb/ipaddressgroup"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/fileparse"
)

func CreateIPGroupCommand() *cobra.Command {
	var IPGroupCmd = &cobra.Command{
		Use:   "ipGroup",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples`,
		Run:   runIPGroup,
		Args:  cobra.NoArgs,
	}

	IPGroupCmd.PersistentFlags().StringP("operation", "o", "list", "操作类型: [list, update]")
	IPGroupCmd.PersistentFlags().StringP("excel", "e", "ipaddr_group_for_elb.xlsx", "存有 IP 地址组的文件")
	IPGroupCmd.PersistentFlags().StringP("addr-group-name", "n", "测试地址组", "地址组名称")
	// 功能说明：是否只预检此次请求
	// 取值范围：
	// -true：发送检查请求，不会更新地址组内容。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。
	// -false（默认值）：发送正常请求，并直接更新地址组。
	IPGroupCmd.PersistentFlags().BoolP("dry-run", "d", false, "是否真实执行操作")

	return IPGroupCmd
}

func runIPGroup(cmd *cobra.Command, args []string) {
	// 获取全部命令行标志
	operation, _ := cmd.Flags().GetString("operation")
	ipsFile, _ := cmd.Flags().GetString("excel")
	addrGroupName, _ := cmd.Flags().GetString("addr-group-name")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	e := ipaddressgroup.NewElbIPAddressGroup(elbClient)
	// 执行操作
	switch operation {
	case "list":
		e.ListAddressGroup()
	case "update":
		ipset, id, err := fileparse.GetElbIPaddrGroup(ipsFile, addrGroupName)
		if err != nil {
			logrus.Fatalf("解析文件失败: %v", err)
		}
		e.UpdateAddressGroup(addrGroupName, id, ipset, dryRun)
	default:
		logrus.Fatalln("操作类型不存在，请使用 -o 指定操作类型")
	}
}
