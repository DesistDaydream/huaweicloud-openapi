package ecs

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	ecsregionv2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
)

type EcsClient struct {
	Client *ecs.EcsClient
}

// 创建控制 ECS 的客户端
func NewEcsClient(ak, sk, region string) (*EcsClient, error) {
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	hcClient := ecs.EcsClientBuilder().
		WithRegion(ecsregionv2.ValueOf(region)).
		WithCredential(auth).
		Build()

	client := ecs.NewEcsClient(hcClient)

	return &EcsClient{
		Client: client,
	}, nil
}
