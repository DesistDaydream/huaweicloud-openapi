package vpc

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type AddressGroupFlags struct {
	ipsFile       string
	addrGroupName string
	dryRun        bool
}

var (
	addressGroupFlags AddressGroupFlags
)

func CreateAddressGroupCmd() *cobra.Command {
	var addressGroupCmd = &cobra.Command{
		Use:   "addrGroup",
		Short: "管理 VPC 的 IP 地址组",
	}

	addressGroupCmd.PersistentFlags().StringVarP(&addressGroupFlags.ipsFile, "excel", "e", "ipaddr_group_for_vpc.xlsx", "存有 IP 地址组的文件")
	addressGroupCmd.PersistentFlags().StringVarP(&addressGroupFlags.addrGroupName, "ag-name", "n", "测试地址组", "地址组名称")
	// 功能说明：是否只预检此次请求
	// 取值范围：
	// -true：发送检查请求，不会更新地址组内容。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。
	// -false（默认值）：发送正常请求，并直接更新地址组。
	addressGroupCmd.PersistentFlags().BoolVarP(&addressGroupFlags.dryRun, "dry-run", "d", false, "是否真实执行操作")

	addressGroupCmd.AddCommand(
		IPGroupListCommand(),
		IPGroupUpdateCommand(),
		IPGroupRulesListCommand(),
	)

	return addressGroupCmd
}

func findIpGroupID(sgName string) (string, error) {
	if addressGroupFlags.addrGroupName == "" {
		logrus.Fatalf("请指定地址组名称")
	}

	var agID string

	resp, err := vpcClient.Client.ListAddressGroup(&model.ListAddressGroupRequest{})
	if err != nil {
		logrus.Errorln(err)
	}

	for _, ag := range *resp.AddressGroups {
		if ag.Name == addressGroupFlags.addrGroupName {
			agID = ag.Id
		}
	}

	return agID, nil
}
