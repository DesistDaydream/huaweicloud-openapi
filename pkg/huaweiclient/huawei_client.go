package huaweiclient

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	regionv2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3"
	regionv3 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/region"
	"github.com/spf13/pflag"
)

// 创建客户端命令行标志
type ClientFlags struct {
	AuthFile string
	UserName string
}

// 添加命令行标志
func (c *ClientFlags) AddFlags() {
	pflag.StringVarP(&c.AuthFile, "auth-file", "f", "auth.yaml", "认证信息文件")
	pflag.StringVarP(&c.UserName, "username", "u", "", "用户名")
}

// 创建控制 ECS 的客户端
func CreateEcsClient(ak, sk string) (*ecs.EcsClient, error) {
	auth := basic.NewCredentialsBuilder().WithAk(ak).WithSk(sk).Build()

	hcClient := ecs.EcsClientBuilder().WithRegion(regionv2.ValueOf("cn-southwest-2")).WithCredential(auth).Build()

	client := ecs.NewEcsClient(hcClient)

	return client, nil
}

// 创建控制 VPC 的客户端
func CreateVpcClient(ak, sk string) (*vpc.VpcClient, error) {
	auth := basic.NewCredentialsBuilder().WithAk(ak).WithSk(sk).Build()

	client := vpc.NewVpcClient(vpc.VpcClientBuilder().
		WithRegion(regionv3.ValueOf("cn-southwest-2")).
		WithCredential(auth).Build())
	return client, nil
}
