/*
子命令 - 服务器信息

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Server information query tool",
	Long:  "A tool that queries server hardware and software information",
	Run: func(cmd *cobra.Command, args []string) {
		var cache = "== Server Info ====================================="
		cache += getHostInfo()
		cache += getCpuInfo()
		cache += getMemInfo()
		cache += getDiskInfo()
		cache += getNetInfo()
		cache += "\n============================= " + getNow() + " =="
		fmt.Println(cache)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

// 获取CPU信息
func getCpuInfo() string {
	// 获取CPU的物理核数和逻辑核数
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	infos, _ := cpu.Info()
	cpuInfo := strings.Replace(infos[0].ModelName, " CPU", "", -1)
	cpuInfo = strings.Replace(cpuInfo, " Processor", "", -1)
	cpuInfo = strings.Replace(cpuInfo, " @", "", -1)
	cpuInfo = strings.Replace(cpuInfo, "(R)", "", -1)
	cpuInfo = strings.Replace(cpuInfo, "(TM)", "", -1)
	return fmt.Sprintf(`
CPU Name 	%s
Physical Core	%d core
Logical Core	%d core`, cpuInfo, physicalCnt, logicalCnt)
}

// 获取内存信息
func getMemInfo() string {
	v, _ := mem.VirtualMemory()
	s, _ := mem.SwapMemory()
	virtualTotal := v.Total / 1024 / 1024 / 1024
	swapTotal := s.Total / 1024 / 1024 / 1024
	return "\n" + fmt.Sprintf(`Virtual Memory	%d GB
Swap Memory	%d GB`, virtualTotal, swapTotal)
}

// 获取磁盘信息
func getDiskInfo() string {
	var total uint64 = 0
	// 获取磁盘的分区信息
	info, _ := disk.Partitions(true)
	if info[0].Mountpoint == "/" {
		usage, _ := disk.Usage("/")
		total += usage.Total
	} else {
		for i := 0; i < len(info); i++ {
			usage, _ := disk.Usage(info[i].Mountpoint)
			total += usage.Total
		}
	}
	total = total / 1024 / 1024 / 1024
	return "\n" + fmt.Sprintf("Disk Usage	%v GB", total)
}

// 获取网络信息
func getNetInfo() string {
	list, _ := net.Interfaces()
	var cache = fmt.Sprintf("\nNetwork		%v interfaces", len(list))
	for i := 0; i < len(list); i++ {
		for x := 0; x < len(list[i].Addrs); x++ {
			addr := list[i].Addrs[x].Addr
			if addr == "127.0.0.1/8" || addr == "::1/128" || addr == "fe80::1/64" {
				continue
			}
			cache += "\n" + `		` + list[i].Addrs[x].Addr
		}
	}
	return "\n" + cache
}

// 获取主机信息
func getHostInfo() string {
	info, _ := host.Info()
	bootTime := time.Unix(int64(info.BootTime), 0).Local().Format("2006-01-02 15:04:05")
	return "\n" + fmt.Sprintf(`Server Id	%s
System		%s
System Version	%s
Family		%s
Kernel Version	%s
Kernel Arch	%s
Boot Time	%s
Up Time		%v ms`,
		info.HostID,
		info.Platform,
		info.PlatformFamily,
		info.PlatformVersion,
		info.KernelVersion,
		info.KernelArch,
		bootTime,
		info.Uptime) + "\n"
}

// 获取当前时间
func getNow() string {
	now := time.Now()
	format := "2006-01-02 15:04:05"
	return now.Format(format)
}
