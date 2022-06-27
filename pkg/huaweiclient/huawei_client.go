package huaweiclient

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	ecsregionv2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	elb "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3"
	elbregionv3 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/region"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3"
	vpcregionv3 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/region"
	"github.com/spf13/pflag"
)

// 创建客户端命令行标志
type ClientFlags struct {
	AuthFile string
	UserName string
	Region   string
}

// 添加命令行标志
func (c *ClientFlags) AddFlags() {
	pflag.StringVarP(&c.AuthFile, "auth-file", "f", "auth.yaml", "认证信息文件")
	pflag.StringVarP(&c.UserName, "username", "u", "", "用户名")
	pflag.StringVarP(&c.Region, "region", "r", "cn-southwest-2", "地域")
}

// 创建控制 ECS 的客户端
func CreateEcsClient(ak, sk, region string) (*ecs.EcsClient, error) {
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	hcClient := ecs.EcsClientBuilder().
		WithRegion(ecsregionv2.ValueOf(region)).
		WithCredential(auth).
		Build()

	client := ecs.NewEcsClient(hcClient)

	return client, nil
}

// 创建控制 VPC 的客户端
func CreateVpcClient(ak, sk, region string) (*vpc.VpcClient, error) {
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := vpc.NewVpcClient(vpc.VpcClientBuilder().
		WithRegion(vpcregionv3.ValueOf(region)).
		WithCredential(auth).
		Build())

	return client, nil
}

// 创建控制 ELB 的客户端
func CreateElbClient(ak, sk, region string) (*elb.ElbClient, error) {
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := elb.NewElbClient(elb.ElbClientBuilder().
		WithRegion(elbregionv3.ValueOf(region)).
		WithCredential(auth).
		Build())

	return client, nil
}
