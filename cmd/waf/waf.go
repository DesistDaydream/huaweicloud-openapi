package waf

import (
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/huaweiclient"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/waf"
	"github.com/spf13/cobra"
)

var WafClient *waf.WafClient

func CreateCommand() *cobra.Command {
	WafCmd := &cobra.Command{
		Use:              "waf",
		Short:            "控制 WAF 资源",
		PersistentPreRun: wafPersistentPreRun,
	}

	WafCmd.AddCommand(
		CreateIPGroupCommand(),
	)

	return WafCmd
}

func wafPersistentPreRun(cmd *cobra.Command, args []string) {
	// 执行根命令的初始化操作
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}

	// 初始化账号Client
	client, err := waf.NewWafClient(
		huaweiclient.Info.AK,
		huaweiclient.Info.SK,
		huaweiclient.Info.Region,
	)
	if err != nil {
		panic(err)
	}

	WafClient = client
}
