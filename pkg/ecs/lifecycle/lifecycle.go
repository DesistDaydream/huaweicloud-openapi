package lifecycle

import (
	"fmt"

	hwcecs "github.com/DesistDaydream/huaweicloud-openapi/pkg/ecs"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/sirupsen/logrus"
)

type EcsLifecycle struct {
	EcsHandler *hwcecs.EcsHandler
}

func NewEcsLifecycle(ecsHandler *hwcecs.EcsHandler) *EcsLifecycle {
	return &EcsLifecycle{
		EcsHandler: ecsHandler,
	}
}

func (d *EcsLifecycle) ShowServer(id string) {
	request := &model.ShowServerRequest{}
	request.ServerId = id
	response, err := d.EcsHandler.Client.ShowServer(request)
	if err == nil {
		fmt.Println(err)
	}

	logrus.Info(response)
}
