/*
主命令

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"fmt"
	"monitor-control/config"
	"os"

	"github.com/spf13/cobra"
)

var headLogo = `   __  ___         __            _____          __           __
  /  |/  /__ ____ / /____ ____  / ___/__  ___  / /________  / /
 / /|_/ / _ '(_-</ __/ -_) __/ / /__/ _ \/ _ \/ __/ __/ _ \/ / 
/_/  /_/\_,_/___/\__/\__/_/    \___/\___/_//_/\__/_/  \___/_/  

==================== Developed by SkyeZhang ====================`

// 没有任何子命令时调用的基本命令
var rootCmd = &cobra.Command{
	Use:     "monitor-control",
	Short:   "A server monitoring master control",
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
