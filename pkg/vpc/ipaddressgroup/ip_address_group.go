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

func (v *VpcIPAddressGroup) ListAddressGroup() {
	request := &model.ListAddressGroupRequest{}
	response, err := v.VpcHandler.Client.ListAddressGroup(request)
	if err != nil {
		logrus.Infoln(err)
	}
	logrus.Infoln(response)
}

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
	response, err := v.VpcHandler.Client.UpdateAddressGroup(request)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Infoln(response)
}
