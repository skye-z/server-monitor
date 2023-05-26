/*
子命令 - 开启监控服务

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"log"
	"os"
	"server-monitor/api"
	"server-monitor/config"
	"server-monitor/daemon"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  "Start server monitoring and reporting service",
	Run: func(cmd *cobra.Command, args []string) {
		var pid = config.GetInt32("service.pid")
		if pid != 0 {
			log.Println("Delivery service is running")
			log.Println("Api service is running")
			return
		}

		if config.GetBool("service.delivery") {
			if config.GetBool("service.api") {
				api.RunServer()
			} else {
				log.Println("Api server not enabled")
			}
			runServer()
		} else {
			if config.GetBool("service.api") {
				log.Println("Api server depends on delivery server")
			}
			log.Println("Delivery server not enabled")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func runServer() {
	envIdx, _ := strconv.Atoi(os.Getenv("BETAX_SMD_IDX"))
	if envIdx > 1 {
		log.Println("Delivery server starting...")
	} else {
		log.Println("Daemon server starting...")
	}
	logFile := "smr.log"
	d := daemon.NewDaemon("Delivery", logFile)
	d.MaxCount = config.GetInt("service.max-retry")
	d.Run()
	// 开始定时任务
	rate := config.GetInt32("service.rate")
	ticker := time.NewTicker(time.Duration(rate) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		// 任务代码
	}
	log.Println("Delivery server started")
}
