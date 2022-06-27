package fileparse

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/model"
	"github.com/sirupsen/logrus"
)

func TestGetElbIPaddrGroup(t *testing.T) {
	type args struct {
		file          string
		addrGroupName string
	}
	tests := []struct {
		name       string
		args       args
		wantIpList []model.UpadateIpGroupIpOption
		wantId     string
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "测试",
			args: args{
				file:          "../../ipaddr_group_for_elb.xlsx",
				addrGroupName: "测试地址组",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIpList, gotId, err := GetElbIPaddrGroup(tt.args.file, tt.args.addrGroupName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetElbIPaddrGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, r := range gotIpList {
				logrus.WithFields(logrus.Fields{
					"ip":   r.Ip,
					"desc": *r.Description,
				}).Infof("检查解析出的 IP 信息")
			}
			logrus.Infof("检查解析出的 IP 地址组 ID: %v", gotId)
		})
	}
}
