package elb

import (
	"fmt"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/fileparse"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func IPGroupUpdateCommand() *cobra.Command {
	var ipGroupUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "更新 ELB 的 IP 地址组",
		Run:   runIpGroupUpdate,
		Args:  cobra.NoArgs,
	}

	return ipGroupUpdateCmd
}

// 全量更新 IP 地址组(若 IP 地址组中存在 Excel 中不存在的 IP 地址，则删除)
func runIpGroupUpdate(cmd *cobra.Command, args []string) {
	ipset, id, err := fileparse.GetElbIPaddrGroup(ipGroupFlags.ipsFile, ipGroupFlags.addrGroupName)
	if err != nil {
		logrus.Fatalf("解析文件失败: %v", err)
	}
	request := &model.UpdateIpGroupRequest{}
	request.IpgroupId = id
	request.Body = &model.UpdateIpGroupRequestBody{
		Ipgroup: &model.UpdateIpGroupOption{
			IpList: &ipset,
		},
	}

	for _, ip := range ipset {
		logrus.WithFields(logrus.Fields{
			"ip":   ip.Ip,
			"desc": *ip.Description,
		}).Infoln("检查将要更新的 IP 地址")
	}

	// 手动确认是否执行
	if !ipGroupFlags.dryRun {
		logrus.Infof("请确认是否要更新【%v】地址组，输入y/Y确认，输入其他键退出", ipGroupFlags.addrGroupName)
		var input string
		_, _ = fmt.Scanln(&input)
		if input != "y" && input != "Y" {
			logrus.Infoln("退出")
			return
		}
	}

	resp, err := elbClient.Client.UpdateIpGroup(request)
	if err != nil {
		logrus.Errorf("更新失败: %v", err)
	} else {
		logrus.Infoln("更新成功，任务ID：", *resp.RequestId)
	}
}
