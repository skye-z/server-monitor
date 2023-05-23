/*
子命令 - 停止监控服务

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"fmt"
	"server-monitor/config"
	"time"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop server",
	Long:  "Stop server monitoring and reporting service",
	Run: func(cmd *cobra.Command, args []string) {
		var pid = config.GetInt32("service.pid")
		if pid == 0 {
			fmt.Println("Service not running")
			return
		}
		p, err := process.NewProcess(pid)
		if err != nil {
			fmt.Println("Service not running")
			config.Set("service.pid", "0")
			return
		}
		fmt.Println("Service is stopping")

		ppid, _ := p.Ppid()
		daemon, _ := process.NewProcess(ppid)
		daemon.Kill()
		time.Sleep(time.Second * 1)
		p.Kill()

		fmt.Println("Service stopped")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
