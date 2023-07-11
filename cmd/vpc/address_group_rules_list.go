package vpc

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
)

func IPGroupRulesListCommand() *cobra.Command {
	var addressGroupRulesListCmd = &cobra.Command{
		Use:   "rulesList",
		Short: "列出 VPC 的所有 IP 地址组",
		Run:   runAddressGroupRulesList,
		Args:  cobra.NoArgs,
	}

	return addressGroupRulesListCmd
}

// 列出所有 IP 地址组
func runAddressGroupRulesList(cmd *cobra.Command, args []string) {

	agID, err := findIpGroupID(addressGroupFlags.addrGroupName)
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	showAddressGroupResponse, err := vpcClient.Client.ShowAddressGroup(&model.ShowAddressGroupRequest{
		AddressGroupId: agID,
	})
	if err != nil {
		logrus.Errorf("获取安全组规则错误，原因: %v", err)
	}

	for _, ip := range showAddressGroupResponse.AddressGroup.IpSet {
		fmt.Println(ip)
	}

}
