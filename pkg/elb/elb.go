package elb

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	elb "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3"
	elbregionv3 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/region"
)

type ElbClient struct {
	Client *elb.ElbClient
}

// 创建控制 ELB 的客户端
func NewElbClient(ak, sk, region string) (*ElbClient, error) {
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := elb.NewElbClient(elb.ElbClientBuilder().
		WithRegion(elbregionv3.ValueOf(region)).
		WithCredential(auth).
		Build())

	return &ElbClient{
		Client: client,
	}, nil
}
