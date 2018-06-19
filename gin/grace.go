package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
)

func main() {
	wd, _ := os.Getwd()
	pidFile, err := os.OpenFile(filepath.Join(wd, "wechat.pid"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("failed to create pid file: %s", err.Error())
		os.Exit(1)
	}
	pidFile.WriteString(strconv.Itoa(os.Getpid()))
	pidFile.Close()

	router := gin.New()

	router.Static("/web", "./web")

	router.Use(gin.Recovery())

	//测试接口
	registerTestAPI(router)

	err = gracehttp.Serve(
		&http.Server{Addr: fmt.Sprintf(":%d", conf.ServiceConfig.HttpPort), Handler: router},
	)

	if err != nil {
		log.Printf("fatal error: %s", err.Error())
	}
}
