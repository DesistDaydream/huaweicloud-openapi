package vpc

import (
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3"
)

type VpcHandler struct {
	Client *vpc.VpcClient
}

func NewVpcHandler(client *vpc.VpcClient) *VpcHandler {
	return &VpcHandler{
		Client: client,
	}
}
