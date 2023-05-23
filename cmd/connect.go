/*
子命令 - 连接主控

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type ConnectConfig struct {
	ssl  bool
	addr string
	port int
	key  string
}

var conf = new(ConnectConfig)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the monitoring system master control",
	Long:  "A master control system connection tool that can provide central control binding and regular data reporting services.",
	Run: func(cmd *cobra.Command, args []string) {
		var state = check()
		if state {
			var prefix = "http"
			if conf.ssl {
				prefix = "https"
			}
			url := fmt.Sprintf("%s://%s:%v?key=%s", prefix, conf.addr, conf.port, conf.key)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(fmt.Sprintf("connect %s error", url), err)
			} else {
				fmt.Println(fmt.Sprintf("connect %s", url), resp)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().BoolVarP(&conf.ssl, "ssl", "s", false, "ssl connect")
	connectCmd.Flags().StringVarP(&conf.addr, "addr", "a", "", "master service host")
	connectCmd.Flags().IntVarP(&conf.port, "port", "p", 0, "master service port")
	connectCmd.Flags().StringVarP(&conf.key, "key", "k", "", "master connect key")
}

func check() bool {
	var state = true
	if conf.addr == "" {
		fmt.Println("Please use -h to set master service address")
		state = false
	}
	if conf.port == 0 {
		fmt.Println("Please use -p to set master service port")
		state = false
	}
	if conf.key == "" {
		fmt.Println("Please use -k to set master connect key")
		state = false
	}
	return state
}
