package vpc

import (
	"fmt"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/fileparse"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func IPGroupUpdateCommand() *cobra.Command {
	var ipGroupUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "更新 VPC 中指定的 IP 地址组",
		Run:   runIpGroupUpdate,
		Args:  cobra.NoArgs,
	}

	return ipGroupUpdateCmd
}

// 全量更新 IP 地址组(若 IP 地址组中存在 Excel 中不存在的 IP 地址，则删除)
func runIpGroupUpdate(cmd *cobra.Command, args []string) {
	ipset, id, err := fileparse.GetVpcIPaddrGroup(ipGroupFlags.ipsFile, ipGroupFlags.addrGroupName)
	if err != nil {
		logrus.Fatalf("解析文件失败: %v", err)
	}

	request := &model.UpdateAddressGroupRequest{}
	request.AddressGroupId = id
	request.Body = &model.UpdateAddressGroupRequestBody{
		AddressGroup: &model.UpdateAddressGroupOption{
			Name:  &ipGroupFlags.addrGroupName,
			IpSet: &ipset,
		},
		DryRun: &ipGroupFlags.dryRun,
	}

	for _, ip := range ipset {
		logrus.WithField("ip", ip).Infoln("检查将要更新的 IP 地址")
	}

	// 如果 dryRun 为 false，则手动确认是否执行
	if !ipGroupFlags.dryRun {
		logrus.Infof("请确认是否要更新【%v】地址组，输入y/Y确认，输入其他键退出", ipGroupFlags.addrGroupName)
		var input string
		_, _ = fmt.Scanln(&input)
		if input != "y" && input != "Y" {
			logrus.Infoln("退出")
			return
		}
	}

	// 执行操作，更新地址组
	resp, err := VpcClient.Client.UpdateAddressGroup(request)
	if err != nil {
		logrus.Errorf("更新失败: %v", err)
	} else {
		logrus.Infoln("更新成功，任务ID：", *resp.RequestId)
	}
}
