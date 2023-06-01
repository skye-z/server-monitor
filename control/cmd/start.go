/*
子命令 - 开启主控服务

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"log"
	"monitor-control/config"
	"monitor-control/daemon"
	"monitor-control/service"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  "Start control server",
	Run: func(cmd *cobra.Command, args []string) {
		var pid = config.GetInt32("service.pid")
		if pid != 0 {
			log.Println("Control service is running")
			return
		}

		runServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func runServer() {
	envIdx, _ := strconv.Atoi(os.Getenv("BETAX_SCD_IDX"))
	if envIdx > 1 {
		log.Println("Control server starting...")
	} else {
		log.Println("Daemon server starting...")
	}
	logFile := "smr.log"
	d := daemon.NewDaemon("Control", logFile)
	d.MaxCount = config.GetInt("service.max-retry")
	d.Run()

	service.Run()

	log.Println("Control server started")
}
