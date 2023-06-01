/*
子命令 - 停止主控服务

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"log"
	"monitor-control/config"
	"time"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop server",
	Long:  "Stop control server",
	Run: func(cmd *cobra.Command, args []string) {
		stopServer()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func stopServer() {
	var pid = config.GetInt32("service.pid")
	if pid == 0 {
		log.Println("Control service not running")
		return
	}
	p, err := process.NewProcess(pid)
	if err != nil {
		log.Println("Control service not running")
		config.Set("service.pid", "0")
		return
	}
	log.Println("Control service stopping...")

	ppid, _ := p.Ppid()
	daemon, _ := process.NewProcess(ppid)
	daemon.Kill()
	time.Sleep(time.Second * 1)
	p.Kill()

	log.Println("Control service stopped")
}
