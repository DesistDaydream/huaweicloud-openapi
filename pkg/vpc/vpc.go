package vpc

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3"
	vpcregionv3 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/region"
)

type VpcClient struct {
	Client *vpc.VpcClient
}

// 创建控制 VPC 的客户端
func NewVpcClient(ak, sk, region string) (*VpcClient, error) {
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := vpc.NewVpcClient(vpc.VpcClientBuilder().
		WithRegion(vpcregionv3.ValueOf(region)).
		WithCredential(auth).
		Build())

	return &VpcClient{
		Client: client,
	}, nil
}
