/*
子命令 - 状态监控

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"fmt"
	"monitor-control/config"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Server Status",
	Long:  "A regularly running server status monitoring program",
	Run: func(cmd *cobra.Command, args []string) {
		checkStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func checkStatus() {
	var pid = config.GetInt32("service.pid")
	if pid == 0 {
		deliveryStop()
		return
	}

	p, err := process.NewProcess(pid)
	if err != nil {
		deliveryStop()
		config.Set("service.pid", "0")
		return
	}

	ppid, _ := p.Ppid()
	fmt.Println(`Daemon Service		running | pid	`, ppid)
	fmt.Println(`Control Service		running | pid	`, p.Pid)

	serPort := config.GetInt32("service.port")
	fmt.Println(`Api Service		running | port	`, serPort)
}

func deliveryStop() {
	fmt.Println(`Daemon Service		stop`)
	fmt.Println(`Control service		stop`)
	fmt.Println(`Api Service		stop`)
}
