package elb

import (
	"os"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/model"
	"github.com/olekukonko/tablewriter"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func IPGroupListCommand() *cobra.Command {
	var ipGroupListCmd = &cobra.Command{
		Use:   "list",
		Short: "列出 ELB 的 IP 地址组",
		Run:   runIpGroupList,
		Args:  cobra.NoArgs,
	}

	return ipGroupListCmd
}

func runIpGroupList(cmd *cobra.Command, args []string) {
	request := &model.ListIpGroupsRequest{}
	resp, err := elbClient.Client.ListIpGroups(request)
	if err != nil {
		logrus.Errorln(err)
	}

	logrus.Infof("当前共有 %v 个 IP 地址组", resp.PageInfo.CurrentCount)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"SG-ID", "名称", "描述"})

	for _, g := range *resp.Ipgroups {
		table.Append([]string{g.Id, g.Name, g.Description})
	}

	table.Render()
}
