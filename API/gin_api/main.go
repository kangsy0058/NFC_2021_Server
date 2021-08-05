package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"nfc_api/common"
	"nfc_api/database"
	_ "nfc_api/docs"
	"nfc_api/firebaseauth"
	"nfc_api/kiosk"
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
			kiosk_router.GET("/sncheck/:werableSN", kiosk.CheckWearableSN)
			kiosk_router.POST("/userlog", kiosk.Userlog)

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
		commomn_router.GET("/user/info", common.UserInfo)
		commomn_router.POST("/user/change", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"rtmsg": "Success",
			})
		})

		//web

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
			user_admin_router.GET("/subgroup/lookup",common.SubGroupLookup) //하위그룹관리 상위관리자에서 하위관리자 정보 조회
			user_admin_router.GET("/subgroup/device/lookup/all",common.DeviceGroupLookUp) //디바이스 전체 조회
			user_admin_router.GET("/subgroup/device/lookup/group",common.DeviceGroupLookUp) // 디바이스 그룹별 조회
			user_admin_router.POST("/subgroup/device/add",common.DevcieGroupAdd) //디바이스 생성
			user_admin_router.DELETE("/subgroup/device/del",common.DeviceGroupDel) //디바이스 삭제
			user_admin_router.POST("/subgroup/authadd",common.GroupAuthAdd) //상위관리자 권한 부여
			user_admin_router.GET("/account/lookup",common.AdminAccountLook) //계정관리 조회
			user_admin_router.DELETE("/account/del",common.AdminAccounthDel) //계정 삭제
			user_admin_router.POST("/account/post",common.AdminAccountPost) //계정 수정
			// 작업중 user_admin_router.GET("/wearabledevice",common.AdminDeviceSK) //특정 사용자 동선 조회(시간순)
			// 작업중 user_admin_router.GET("/wearabledevice") //특정 사용자와 겹치는 사용자 조회
			//user_admin_router.GET("/werabledevicelook",common.) //웨어러블디바이스 조회
			user_admin_router.POST("/wearabledevice", func(c *gin.Context) {
				c.JSON(http.StatusCreated, gin.H{
					"rtmsg": "Success",
				})
			})//웨어러블디바이스 생성
			user_admin_router.PUT("/wearabledevice", func(c *gin.Context) {
				c.JSON(http.StatusCreated, gin.H{
					"rtmsg": "Success",
				})
			}) //웨어러블디바이스 수정
			user_admin_router.DELETE("/wearabledevice", func(c *gin.Context){
				c.JSON(http.StatusAccepted, gin.H{
					"rtmsg": "Success",
				})
			})//웨어러블디바이스 삭제
			// 명세서 TBD user_admin_router.GET("/nfclog") // NFC 태그 기록

			user_admin_router.POST("/deviceMGMT", func(c *gin.Context) {
				c.JSON(http.StatusCreated, gin.H{
					"rtmsg" : "Success",
				})
			})//하위관리자 디바이스 생성
			user_admin_router.PUT("/deviceMGMT", func(c *gin.Context) {
				c.JSON(http.StatusCreated, gin.H{
					"rtmsg" : "Success",
				})
			})//하위관리자 디바이스 생성 수정

			//명세서 TBD user_admin_router.GET("/deviceMGMT")

			user_admin_router.PUT("accountMGMT",func(c *gin.Context){
				c.JSON(http.StatusCreated, gin.H{
					"rtmsg" : "Success",
				})
			})//자신 게정 수정
			user_admin_router.DELETE("accountMGMT", func(c *gin.Context) {
				c.JSON(http.StatusAccepted,gin.H{
					"rtmsg" : "Success",
				})
			})// 자신 계정 삭제

			//명세서 TBD(자신의 디바이스 로그인 조회) user_admin_router.GET("/devicelog")
			user_admin_router.GET("/dashboard", common.Dashboard)
			//명세서 TBD user_admin_router.GET("/sendUser")
			//명세서 TBD user_amdin_router.GET("/sendGroup")
			//명세서 TBD user_admin_router.GET("/sendAll")

		}
	}
	return r
}

func main() {
	//DB Connection Check
	var version string
	db, _ := database.Mariadb()
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
	// Router setup
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
