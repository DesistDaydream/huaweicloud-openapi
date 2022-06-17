package huaweiclient

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
)

func CreateClient(ak, sk string) (*ecs.EcsClient, error) {
	auth := basic.NewCredentialsBuilder().WithAk(ak).WithSk(sk).Build()

	hcClient := ecs.EcsClientBuilder().WithRegion(region.ValueOf("cn-southwest-2")).WithCredential(auth).Build()

	client := ecs.NewEcsClient(hcClient)

	return client, nil
}
