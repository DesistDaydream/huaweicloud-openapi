package main

import (
	"os"

	"fmt"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/config"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/huaweiclient"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	waf "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1/region"
	"github.com/sirupsen/logrus"
)

func init() {
	// 认证信息文件处理的相关逻辑
	AuthFile := "huawei.yaml"
	UserName := os.Args[1]
	region := "cn-southwest-2"

	// 检查 clientFlags.AuthFile 文件是否存在
	if _, err := os.Stat(AuthFile); os.IsNotExist(err) {
		logrus.Fatal("文件不存在")
	}
	// 获取认证信息
	auth := config.NewAuthInfo(AuthFile)

	// 判断传入的用户是否存在在认证信息中
	if !auth.IsUserExist(UserName) {
		logrus.Fatalf("认证信息中不存在 %v 用户, 请检查认证信息文件或命令行参数的值", UserName)
	}

	huaweiclient.Info = &huaweiclient.ClientInfo{
		AK:     auth.AuthList[UserName].AccessKeyID,
		SK:     auth.AuthList[UserName].SecretAccessKey,
		Region: region,
	}
}

func main() {
	ak := huaweiclient.Info.AK
	sk := huaweiclient.Info.SK

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := waf.NewWafClient(
		waf.WafClientBuilder().
			WithRegion(region.ValueOf("cn-east-3")).
			WithCredential(auth).
			Build())

	request := &model.ListPolicyRequest{}
	response, err := client.ListPolicy(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
