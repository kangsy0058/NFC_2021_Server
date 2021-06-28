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
		// Kiosk API
		kiosk_router := v1.Group("/kiosk")
		{
			kiosk_router.GET("/welcome/:name", kiosk.WelcomeApi)
			kiosk_router.GET("/checksn/:sn", kiosk.CheckWearableSN)
		}
		// Web API
		web_router := v1.Group("/web")
		{
			web_admin_router := web_router.Group("/admin")
			{
				web_admin_router.GET("/test", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"hello": "admin_test"})
				})
			}
		}
		// App API
		app_router := v1.Group("/app")
		{
			app_admin_router := app_router.Group("/admin")
			{
				
				app_admin_router.GET("/test", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"hello": "app_test"})
				})
			}

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
