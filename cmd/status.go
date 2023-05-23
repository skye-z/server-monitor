/*
子命令 - 状态监控

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"fmt"
	"server-monitor/config"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Server Status Monitoring Service",
	Long:  "A regularly running server status monitoring program",
	Run: func(cmd *cobra.Command, args []string) {
		checkStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkStatus() {
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

	fmt.Println("Service is running, pid is", p.Pid)
}
