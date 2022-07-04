package ipaddressgroup

import (
	"fmt"

	"github.com/sirupsen/logrus"

	hwcelb "github.com/DesistDaydream/huaweicloud-openapi/pkg/elb"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/model"
)

type ElbIPAddressGroup struct {
	ElbHandler       *hwcelb.ElbClient
	AddressGroupName string
}

func NewElbIPAddressGroup(elbHandler *hwcelb.ElbClient) *ElbIPAddressGroup {
	return &ElbIPAddressGroup{
		ElbHandler: elbHandler,
	}
}

// 列出所有 IP 地址组
func (v *ElbIPAddressGroup) ListAddressGroup() {
	request := &model.ListIpGroupsRequest{}
	resp, err := v.ElbHandler.Client.ListIpGroups(request)
	if err != nil {
		logrus.Errorln(err)
	}
	logrus.Infof("当前共有 %v 个 IP 地址组", resp.PageInfo.CurrentCount)
	for _, g := range *resp.Ipgroups {
		logrus.Infof("【%v】地址组地址列表：", g.Name)
		for _, ip := range g.IpList {
			fmt.Println(ip.Ip, ip.Description)
		}
	}
}

// 全量更新 IP 地址组(若 IP 地址组中存在 Excel 中不存在的 IP 地址，则删除)
func (v *ElbIPAddressGroup) UpdateAddressGroup(name string, id string, ipList []model.UpadateIpGroupIpOption, dryRun bool) {
	request := &model.UpdateIpGroupRequest{}
	request.IpgroupId = id
	request.Body = &model.UpdateIpGroupRequestBody{
		Ipgroup: &model.UpdateIpGroupOption{
			IpList: &ipList,
		},
	}

	for _, ip := range ipList {
		logrus.WithFields(logrus.Fields{
			"ip":   ip.Ip,
			"desc": *ip.Description,
		}).Infoln("检查将要更新的 IP 地址")
	}

	// 手动确认是否执行
	if !dryRun {
		logrus.Infof("请确认是否要更新【%v】地址组，输入y/Y确认，输入其他键退出", name)
		var input string
		_, _ = fmt.Scanln(&input)
		if input != "y" && input != "Y" {
			logrus.Infoln("退出")
			return
		}
	}

	resp, err := v.ElbHandler.Client.UpdateIpGroup(request)
	if err != nil {
		logrus.Errorf("更新失败: %v", err)
	} else {
		logrus.Infoln("更新成功，任务ID：", *resp.RequestId)
	}
}
