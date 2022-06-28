package vpc

import (
	"os"

	"github.com/DesistDaydream/huaweicloud-openapi/pkg/config"
	"github.com/DesistDaydream/huaweicloud-openapi/pkg/huaweiclient"
	hwcvpc "github.com/DesistDaydream/huaweicloud-openapi/pkg/vpc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CreateSecurityGroupCommand() *cobra.Command {
	long := `操作 VPC 下的安全组资源：
	list 列出所有安全组及其规则
	show 列出指定安全组的规则`

	var SecurityGroupCmd = &cobra.Command{
		Use:   "securityGroup",
		Short: "操作 VPC 下的安全组资源",
		Long:  long,
		Run: func(cmd *cobra.Command, args []string) {
			runSecurityGroup(cmd, args)
		},
		Args: cobra.NoArgs,
	}

	SecurityGroupCmd.PersistentFlags().StringP("operation", "o", "list", "操作类型，详见命令描述")
	SecurityGroupCmd.PersistentFlags().StringP("excel", "e", "security_group_for_vpc.xlsx", "存有安全组的文件")
	SecurityGroupCmd.PersistentFlags().StringP("security-group-name", "n", "default", "安全组名称")
	// 功能说明：是否只预检此次请求
	// 取值范围：
	// -true：发送检查请求，不会更新安全组内容。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。
	// -false（默认值）：发送正常请
	SecurityGroupCmd.PersistentFlags().BoolP("pre-check", "p", false, "是否只预检此次请求")

	return SecurityGroupCmd
}

func runSecurityGroup(cmd *cobra.Command, args []string) {
	// 获取全部命令行标志
	operation, _ := cmd.Flags().GetString("operation")
	// ipsFile, _ := cmd.Flags().GetString("excel")
	securityGroupName, _ := cmd.Flags().GetString("security-group-name")
	// dryRun, _ := cmd.Flags().GetBool("dry-run")
	authFile, _ := cmd.Flags().GetString("auth-file")
	userName, err := cmd.Flags().GetString("username")
	if err != nil {
		logrus.Fatalln("请指定用户名")
	}
	region, _ := cmd.Flags().GetString("region")

	// 检查 clientFlags.AuthFile 文件是否存在
	if _, err := os.Stat(authFile); os.IsNotExist(err) {
		logrus.Fatal("文件不存在")
	}
	// 获取认证信息
	auth := config.NewAuthInfo(authFile)

	// 判断传入的域名是否存在在认证信息中
	if !auth.IsUserExist(userName) {
		logrus.Fatalf("认证信息中不存在 %v 用户, 请检查认证信息文件或命令行参数的值", userName)
	}

	// 初始化账号Client
	client, err := huaweiclient.CreateVpcClient(
		auth.AuthList[userName].AccessKeyID,
		auth.AuthList[userName].SecretAccessKey,
		region,
	)
	if err != nil {
		panic(err)
	}

	// 实例化 VPC 的 API 处理器
	h := hwcvpc.NewVpcHandler(client)

	v := hwcvpc.NewVpcSecurityGroup(h)

	// 执行操作
	switch operation {
	case "list":
		v.ListAllSecurityGroupAndRules()
	case "show":
		v.ListSecurityGroupRules(securityGroupName)
	case "update":
		// TODO
	default:
		logrus.Fatalln("操作类型不存在，请使用 -o 指定操作类型")
	}
}
