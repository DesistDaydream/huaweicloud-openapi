package vpc

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
	"github.com/sirupsen/logrus"
)

type VpcSecurityGroup struct {
	VpcClient         *VpcClient
	SecurityGroupName string
}

func NewVpcSecurityGroup(vpcClient *VpcClient) *VpcSecurityGroup {
	return &VpcSecurityGroup{
		VpcClient: vpcClient,
	}
}

// 显示执行安全组的名称、ID、描述、规则，返回安全组规则的详细信息
func (v *VpcSecurityGroup) ShowSecurityGroup(id string) (*[]model.SecurityGroupRule, error) {
	request := &model.ShowSecurityGroupRequest{}
	request.SecurityGroupId = id
	resp, err := v.VpcClient.Client.ShowSecurityGroup(request)
	if err != nil {
		return nil, err
	}
	return &resp.SecurityGroup.SecurityGroupRules, nil
}

// 列出所有安全组
func (v *VpcSecurityGroup) ListSecurityGroup() (*[]model.SecurityGroup, error) {
	request := &model.ListSecurityGroupsRequest{}
	resp, err := v.VpcClient.Client.ListSecurityGroups(request)
	if err != nil {
		return nil, err
	}

	return resp.SecurityGroups, nil
}

// 列出所有安全组下的所有规则
func (v *VpcSecurityGroup) ListAllSecurityGroupAndRules() {
	sgs, err := v.ListSecurityGroup()
	if err != nil {
		logrus.Errorf("列出所有安全组错误: %v", err)
	}
	logrus.Infof("当前共有 %v 个安全组", len(*sgs))

	for _, sg := range *sgs {
		rules, err := v.ShowSecurityGroup(sg.Id)
		if err != nil {
			logrus.Errorf("列出安全组 %v 下的规则错误: %v", sg.Name, err)
		}
		logrus.Infof("【%v】安全组共有 %v 个规则", sg.Name, len(*rules))
		for _, r := range *rules {
			logrus.WithFields(logrus.Fields{
				"desc":     r.Description,
				"portocol": r.Protocol,
				"port":     r.Multiport,
			}).Infof("【%v】安全组规则", sg.Name)
		}

	}
}

// 列出指定安全组下的所有规则
func (v *VpcSecurityGroup) ListSecurityGroupRules(securityGroupName string) {
	sgs, err := v.ListSecurityGroup()
	if err != nil {
		logrus.Errorf("列出所有安全组错误: %v", err)
	}

	for _, sg := range *sgs {
		if sg.Name == securityGroupName {
			rules, err := v.ShowSecurityGroup(sg.Id)
			if err != nil {
				logrus.Errorf("列出安全组 %v 下的规则错误: %v", sg.Name, err)
			}
			logrus.Infof("【%v】安全组共有 %v 个规则", sg.Name, len(*rules))
			for _, r := range *rules {
				if r.Protocol == "" {
					r.Protocol = "全部"
				}
				if r.Multiport == "" {
					r.Multiport = "全部"
				}
				logrus.WithFields(logrus.Fields{
					"方向": r.Direction,
					"描述": r.Description,
					"协议": r.Protocol,
					"端口": r.Multiport,
					"ID": r.Id,
				}).Infof("【%v】安全组规则", sg.Name)
			}
		}
	}
}
