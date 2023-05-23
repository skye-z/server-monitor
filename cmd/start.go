/*
子命令 - 开启监控服务

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"log"
	"server-monitor/config"
	"server-monitor/daemon"
	"time"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  "Start server monitoring and reporting service",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func runServer() {
	logFile := "smr.log"
	d := daemon.NewDaemon(logFile)
	d.MaxCount = 5
	d.Run()
	// 开始定时任务
	rate := config.GetInt32("service.rate")
	ticker := time.NewTicker(time.Duration(rate) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		log.Println("working...")
	}
}
