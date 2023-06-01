package main

import (
	"monitor-control/cmd"
	"monitor-control/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
