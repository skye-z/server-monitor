/*
接口服务

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/

package api

import (
	"fmt"
	"log"
	"monitor-client/config"
	"net/http"
	"os"
	"strconv"
	"time"
)

var server *http.Server

func RunServer() {
	envIdx, _ := strconv.Atoi(os.Getenv("BETAX_SMD_IDX"))
	if envIdx != 1 {
		return
	}
	log.Println("Api server starting...")
	handler := http.NewServeMux()
	handler.HandleFunc("/api/conf", getConfig)

	serHost := config.GetString("service.host")
	serPort := config.GetInt32("service.port")
	server = &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf("%s:%v", serHost, serPort),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Println("Api server startup success, port is", serPort)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println("Api start error:", err)
		}
	}()
}

func getConfig(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("path", request.URL.Path)
	fmt.Fprintln(writer, "hello world")
}
