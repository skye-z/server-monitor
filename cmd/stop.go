/*
子命令 - 停止监控服务

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"log"
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
		stopServer()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func stopServer() {
	if config.GetBool("service.delivery") {
		var pid = config.GetInt32("service.pid")
		if pid == 0 {
			log.Println("Delivery service not running")
			return
		}
		p, err := process.NewProcess(pid)
		if err != nil {
			log.Println("Delivery service not running")
			config.Set("service.pid", "0")
			return
		}
		log.Println("Delivery service stopping...")

		ppid, _ := p.Ppid()
		daemon, _ := process.NewProcess(ppid)
		daemon.Kill()
		time.Sleep(time.Second * 1)
		p.Kill()

		log.Println("Delivery service stopped")
	}
	if config.GetBool("service.api") {
		log.Println("Api service stopped")
	}
}
