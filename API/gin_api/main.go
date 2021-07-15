package main

import (
<<<<<<< HEAD
=======
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
>>>>>>> Haeil
	"net/http"
	"nfc_api/common"
	_ "nfc_api/docs"
	"nfc_api/firebaseauth"
	"nfc_api/kiosk"
<<<<<<< HEAD

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
=======
>>>>>>> Haeil
)

// @title NFC API
// @version 0.0.1
// @description Hoseo NFC 2021 projct API Page

// @host localhost:8080
// @BasePath /

func setupRouter() *gin.Engine {
	r := gin.Default()

	// set gin Context
	r.Use(func(c *gin.Context) {
		// 함수 안에서 사용시
		// firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
		c.Set("firebaseAuth", firebaseauth.SetupFirebase())
	})

	// global middleware setting
	r.Use(CORSMiddleware())

	//Swagger router
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// version 1 router
	v1 := r.Group("/v1")
	{
		// Kiosk API
		kiosk_router := v1.Group("/kiosk")
		{
<<<<<<< HEAD
			kiosk_router.GET("/sncheck/:werableSN", kiosk.CheckWearableSN)
			kiosk_router.POST("/userlog", kiosk.Userlog)
=======
			kiosk_router.GET("/sncheck/", kiosk.CheckWearableSN)
			kiosk_router.PUT("/userlog", kiosk.PutUserlog)
>>>>>>> Haeil

		}

		//Common API
		//app
		commomn_router := v1.Group("/common")
		//commomn_router.Use(firebaseauth.FirebaseAuthMiddleware())
		{
			commomn_router.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})

		}
		commomn_router.POST("/user/device-add", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"rt":    http.StatusCreated,
				"rtmsg": "Success",
			})
		})
		commomn_router.DELETE("/user/device-del", func(c *gin.Context) {
			c.JSON(http.StatusAccepted, gin.H{
				"rtmsg": "Success",
			})
		})
		commomn_router.POST("/user/pid-add", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"rtmsg": "Success",
			})
		})
		commomn_router.DELETE("/user/pid-del", func(c *gin.Context) {
			c.JSON(http.StatusAccepted, gin.H{
				"rtmsg": "Success",
			})
		})
		commomn_router.GET("/userlog/visitHistory", common.VisitHistory)
		commomn_router.POST("/user/FBToken", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"rtmsg": "Success",
			})
		})
		commomn_router.GET("/user/userInfo", common.UserInfo)
		commomn_router.POST("/user/change", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"rtmsg": "Success",
			})
		})

		//web
		commomn_router.GET("/user/login", common.UserLogin)

		//kiosk_Admin API
		kiosk_admin_router := v1.Group("/kioskadmin")
		{
			kiosk_admin_router.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})
		}

		//userlog_Admin API
		user_admin_router := v1.Group("/useradmin")
		{
			user_admin_router.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
					"code":    http.StatusOK,
				})
			})
			//web
<<<<<<< HEAD
			user_admin_router.GET("/subgroup/lookup", common.SubGroupLookup)
			user_admin_router.GET("/subgroup/device/lookup/all", common.DeviceGroupLookUp)
			user_admin_router.GET("/subgroup/device/lookup/group", common.DeviceGroupLookUp)
=======
			user_admin_router.GET("/subgroup/lookup",common.SubGroupLookup)
			user_admin_router.GET("/subgroup/device/lookup/all",common.DeviceGroupLookUp)
			user_admin_router.GET("/subgroup/device/lookup/group",common.DeviceGroupLookUp)
			user_admin_router.POST("/subgroup/device/add",common.DevcieGroupAdd)
			user_admin_router.DELETE("/subgroup/device/del",common.DeviceGroupDel)
			user_admin_router.POST("/subgroup/authadd",common.GroupAuthAdd)
>>>>>>> Haeil
		}
	}
	return r
}

func main() {
<<<<<<< HEAD
	// Router setup
=======



>>>>>>> Haeil
	r := setupRouter()
	//Server start
	r.Run()
}

// CORS 세팅용 middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
