/*
BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package main

import (
	"monitor-client/cmd"
	"monitor-client/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
