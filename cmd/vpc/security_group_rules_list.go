package vpc

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CreateSecurityGroupRulesListCmd() *cobra.Command {
	long := ``

	var securityGroupRulesListCmd = &cobra.Command{
		Use:   "rulesList",
		Short: "列出指定安全组的规则",
		Long:  long,
		Run:   runSecurityGroupRulesList,
	}

	return securityGroupRulesListCmd
}

func runSecurityGroupRulesList(cmd *cobra.Command, args []string) {
	listSecurityGroupsResponse, err := vpcClient.Client.ListSecurityGroups(&model.ListSecurityGroupsRequest{})
	if err != nil {
		logrus.Fatalf("列出安全组异常，原因: %v", err)
	}

	for _, sg := range *listSecurityGroupsResponse.SecurityGroups {
		if sg.Name == sgFlags.sgName {
			showSecurityGroupResponse, err := vpcClient.Client.ShowSecurityGroup(&model.ShowSecurityGroupRequest{
				SecurityGroupId: sg.Id,
			})
			if err != nil {
				logrus.Errorf("列出安全组的规则错误，原因: %v", err)
			}

			logrus.Infof("【%v】安全组共有 %v 个规则", sg.Name, len(showSecurityGroupResponse.SecurityGroup.SecurityGroupRules))

			for _, r := range showSecurityGroupResponse.SecurityGroup.SecurityGroupRules {
				if r.Protocol == "" {
					r.Protocol = "全部"
				}
				if r.Multiport == "" {
					r.Multiport = "全部"
				}
				logrus.WithFields(logrus.Fields{
					"方向":   r.Direction,
					"描述":   r.Description,
					"协议":   r.Protocol,
					"端口":   r.Multiport,
					"规则ID": r.Id,
				}).Infof("【%v】安全组规则", sg.Name)
			}
		}
	}
}
