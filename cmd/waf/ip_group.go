package waf

import "github.com/spf13/cobra"

func CreateIPGroupCommand() *cobra.Command {
	var ipGroupCmd = &cobra.Command{
		Use:   "ipGroup",
		Short: "控制 IPGroup 资源",
		Run:   runIPGroup,
	}
	return ipGroupCmd
}

func runIPGroup(cmd *cobra.Command, args []string) {
	// TODO
}
