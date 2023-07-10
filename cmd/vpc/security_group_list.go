package vpc

import (
	"os"

	"github.com/olekukonko/tablewriter"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CreateSecurityGroupListCmd() *cobra.Command {
	long := ``

	var securityGroupListCmd = &cobra.Command{
		Use:   "list",
		Short: "列出所有安全组",
		Long:  long,
		Run:   runSecurityGroupList,
	}

	return securityGroupListCmd
}

func runSecurityGroupList(cmd *cobra.Command, args []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "名称", "描述"})

	listSecurityGroupsResponse, err := VpcClient.Client.ListSecurityGroups(&model.ListSecurityGroupsRequest{})
	if err != nil {
		logrus.Fatalf("列出安全组异常，原因: %v", err)
	}

	for _, sg := range *listSecurityGroupsResponse.SecurityGroups {
		table.Append([]string{sg.Id, sg.Name, sg.Description})
	}

	table.Render()
}
