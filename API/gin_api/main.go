package main

import (
	"net/http"
	_ "nfc_api/docs"
	"nfc_api/kiosk"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title NFC API
// @version 0.0.1
// @description Hoseo NFC 2021 projct API Page

// @host localhost:8080
// @BasePath /

func setupRouter() *gin.Engine {
	r := gin.Default()
	//Swagger router
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// version 1 router
	v1 := r.Group("/v1")
	{
		// kiosk api
		kiosk_router := v1.Group("/kiosk")
		{
			kiosk_router.GET("/welcome/:name", kiosk.WelcomeApi)
			kiosk_router.GET("/checksn/:sn", kiosk.CheckWearableSN)
		}

		web_router := v1.Group("/web")
		{

		}

	}

	return r
}

func main() {
	r := setupRouter()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
