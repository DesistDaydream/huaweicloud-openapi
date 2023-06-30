package vpc

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func IPGroupListCommand() *cobra.Command {
	var ipGroupListCmd = &cobra.Command{
		Use:   "list",
		Short: "列出 VPC 的所有 IP 地址组",
		Run:   runIpGroupList,
		Args:  cobra.NoArgs,
	}

	return ipGroupListCmd
}

// 列出所有 IP 地址组
func runIpGroupList(cmd *cobra.Command, args []string) {
	request := &model.ListAddressGroupRequest{}
	resp, err := VpcClient.Client.ListAddressGroup(request)
	if err != nil {
		logrus.Errorln(err)
	}
	logrus.Infof("当前共有 %v 个 IP 地址组", resp.PageInfo.CurrentCount)
	for _, ag := range *resp.AddressGroups {
		logrus.Infof("【%v】地址组地址列表：", ag.Name)
		for _, ip := range ag.IpSet {
			fmt.Println(ip)
		}
	}
}
