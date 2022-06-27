package elb

import (
	elb "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3"
)

type ElbHandler struct {
	Client *elb.ElbClient
}

func NewElbHandler(client *elb.ElbClient) *ElbHandler {
	return &ElbHandler{
		Client: client,
	}
}
