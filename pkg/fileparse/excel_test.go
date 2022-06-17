package fileparse

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewExcelData(t *testing.T) {
	type args struct {
		file       string
		domainName string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "测试",
			args: args{
				file:       "VPC的IP地址组.xlsx",
				domainName: "测试地址组",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ips, id, err := GetVpcIPaddrGroup(tt.args.file, tt.args.domainName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewExcelData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, r := range ips {
				logrus.Infoln(r)
			}
			logrus.Infoln(id)
		})
	}
}
