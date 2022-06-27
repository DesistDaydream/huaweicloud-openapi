package ipaddressgroup

import (
	"fmt"

	hwcvpc "github.com/DesistDaydream/huaweicloud-openapi/pkg/vpc"
	"github.com/sirupsen/logrus"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
)

type VpcIPAddressGroup struct {
	VpcHandler       *hwcvpc.VpcHandler
	AddressGroupName string
}

func NewVpcIPADdressGroup(ecsHandler *hwcvpc.VpcHandler) *VpcIPAddressGroup {
	return &VpcIPAddressGroup{
		VpcHandler: ecsHandler,
	}
}

// 列出所有 IP 地址组
func (v *VpcIPAddressGroup) ListAddressGroup() {
	request := &model.ListAddressGroupRequest{}
	response, err := v.VpcHandler.Client.ListAddressGroup(request)
	if err != nil {
		logrus.Infoln(err)
	}
	logrus.Infof("当前共有 %v 个 IP 地址组", response.PageInfo.CurrentCount)
	for _, ag := range *response.AddressGroups {
		logrus.Infof("【%v】地址组地址列表：", ag.Name)
		for _, ip := range ag.IpSet {
			fmt.Println(ip)
		}
	}
}

// 全量更新 IP 地址组(若 IP 地址组中存在 Excel 中不存在的 IP 地址，则删除)
func (v *VpcIPAddressGroup) UpdateAddressGroup(name string, id string, ipset []string, dryRun bool) {
	request := &model.UpdateAddressGroupRequest{}
	request.AddressGroupId = id
	request.Body = &model.UpdateAddressGroupRequestBody{
		AddressGroup: &model.UpdateAddressGroupOption{
			Name:  &name,
			IpSet: &ipset,
		},
		DryRun: &dryRun,
	}

	for _, ip := range ipset {
		logrus.WithField("ip", ip).Infoln("检查将要更新的 IP 地址")
	}

	// 如果 dryRun 为 false，则手动确认是否执行
	if !dryRun {
		logrus.Infof("请确认是否要更新【%v】地址组，输入y/Y确认，输入其他键退出", name)
		var input string
		_, _ = fmt.Scanln(&input)
		if input != "y" && input != "Y" {
			logrus.Infoln("退出")
			return
		}
	}

	// 执行操作，更新地址组
	resp, err := v.VpcHandler.Client.UpdateAddressGroup(request)
	if err != nil {
		logrus.Errorf("更新失败: %v", err)
	} else {
		logrus.Infoln("更新成功，任务ID：", *resp.RequestId)
	}
}
