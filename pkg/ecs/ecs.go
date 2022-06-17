package ecs

import (
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
)

type EcsHandler struct {
	Client *ecs.EcsClient
}

func NewEcsHandler(client *ecs.EcsClient) *EcsHandler {
	return &EcsHandler{
		Client: client,
	}
}
