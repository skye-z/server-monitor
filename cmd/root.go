/*
主命令

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"server-monitor/config"

	"github.com/spf13/cobra"
)

var headLogo = `   ____                       __  ___          _ __
  / __/__ _____  _____ ____  /  |/  /__  ___  (_) /____  ____
 _\ \/ -_) __/ |/ / -_) __/ / /|_/ / _ \/ _ \/ / __/ _ \/ __/
/___/\__/_/  |___/\__/_/   /_/  /_/\___/_//_/_/\__/\___/_/

================== Developed by SkyeZhang ==================`

// 没有任何子命令时调用的基本命令
var rootCmd = &cobra.Command{
	Use:     "server-monitor",
	Short:   "A server monitoring tool",
	Version: config.VERSION,
	Long:    headLogo,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(headLogo + "\n\nI don't know what I can do? Please use -h to view help :)")
	},
}

// 将所有子命令导入主命令并设置标识
// 此函数由 main 调用,只执行一次
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "prints version")
}
