package vpc

import (
	"os"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func IPGroupListCommand() *cobra.Command {
	var addressGroupListCmd = &cobra.Command{
		Use:   "list",
		Short: "列出 VPC 的所有 IP 地址组",
		Run:   runAddressGroupList,
		Args:  cobra.NoArgs,
	}

	return addressGroupListCmd
}

// 列出所有 IP 地址组
func runAddressGroupList(cmd *cobra.Command, args []string) {
	resp, err := vpcClient.Client.ListAddressGroup(&model.ListAddressGroupRequest{})
	if err != nil {
		logrus.Errorln(err)
	}

	logrus.Infof("当前共有 %v 个 IP 地址组", resp.PageInfo.CurrentCount)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"SG-ID", "名称", "描述"})

	for _, ag := range *resp.AddressGroups {
		table.Append([]string{ag.Id, ag.Name, ag.Description})
	}

	table.Render()
}
