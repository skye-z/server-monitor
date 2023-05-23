/*
BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package main

import (
	"server-monitor/cmd"
	"server-monitor/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
