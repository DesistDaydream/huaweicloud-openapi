package waf

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	waf "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1"
	wafregionv1 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1/region"
)

type WafClient struct {
	Client *waf.WafClient
}

// 创建控制 WAF 的客户端
func NewWafClient(ak, sk, region string) (*WafClient, error) {
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := waf.NewWafClient(
		waf.WafClientBuilder().
			WithRegion(wafregionv1.ValueOf(region)).
			WithCredential(auth).
			Build())

	return &WafClient{
		Client: client,
	}, nil
}
