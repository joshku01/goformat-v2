package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goformat-v2/app/global"
	"goformat-v2/app/route"
	"log"
	"os"
)

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	defer func() {
		if err := recover(); err != nil {
			// 補上將err傳至telegram
			//errorcode.ErrorHandler("UNDEFINED_ERROR", err)
			fmt.Println("[❌ Fatal❌ ]:", err)
		}
	}()

	// 開發時，console視窗不顯示有顏色的log
	gin.DisableConsoleColor()

	// 本機開發需要顯示 Gin Log
	var r *gin.Engine
	if os.Getenv("ENV") == "local" {
		r = gin.Default()
	} else {
		r = gin.New()
		r.Use(gin.Recovery())
	}

	// 載入環境設定(所有動作須在該func後執行)
	global.Start()

	//_, err2 := helper.Atoi()
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, errorcode.Failed(err2))
	//})


	route.SetupRouter(r)
	_ = r.Run(":8080")

}
