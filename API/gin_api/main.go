package main

import (
	"log"

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
		commomn_router.POST("/device-add", func(c *gin.Context) {
			Wearable_SN := c.Query("Wearable_SN")
			UUID := c.Query("UUID")
			db, err := database.Mariadb()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			defer db.Close()
			rtmsg := "Success"
			_, err = db.Exec("INSERT into user_info (wearable_SN, UUID) values (?,?) ", Wearable_SN, UUID)
			if err != nil {
				_, err = db.Exec("UPDATE user_info SET wearable_SN = ? WHERE UUID = ?", Wearable_SN, UUID)
				if err != nil {
					rtmsg = "Failed"
					log.Fatal("insert error: ", err)
				}
			}
			c.JSON(http.StatusCreated, gin.H{
				"rtmsg": rtmsg,
			})
		})
		commomn_router.DELETE("/device-del", func(c *gin.Context) {
			Wearable_SN := c.Query("Wearable_SN")
			db, err := database.Mariadb()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			defer db.Close()
			rtmsg := "Success"
			_, err = db.Exec("UPDATE user_info SET wearable_SN = NULL WHERE wearable_SN = ?", Wearable_SN)
			if err != nil {
				rtmsg = "Failed"
				log.Fatal("delete error: ", err)
			}
			c.JSON(http.StatusAccepted, gin.H{
				"rtmsg": rtmsg,
			})
		})
		commomn_router.POST("/user/pid", func(c *gin.Context) {
			PSN := c.Query("PSN")
			PSN_img := c.Query("PSN_img")
			UUID := c.Query("UUID")
			db, err := database.Mariadb()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			defer db.Close()
			rtmsg := "Success"
			_, err = db.Exec("UPDATE user_info SET PSN = ?, PSN_img = ? WHERE UUID = ?", PSN, PSN_img, UUID)
			if err != nil {
				rtmsg = "Failed"
				log.Fatal("insert into users error: ", err)
			}
			c.JSON(http.StatusCreated, gin.H{
				"rtmsg": rtmsg,
			})
		})
		commomn_router.DELETE("/user/pid", func(c *gin.Context) {
			PSN := c.Query("PSN")
			db, err := database.Mariadb()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			defer db.Close()
			rtmsg := "Success"
			_, err = db.Exec("UPDATE user_info SET PSN = NULL, PSN_img = NULL WHERE PSN = ?", PSN)
			if err != nil {
				rtmsg = "Failed"
				log.Fatal("delete error: ", err)
			}
			c.JSON(http.StatusAccepted, gin.H{
				"rtmsg": rtmsg,
			})
		})
		commomn_router.GET("/userlog/visitHistory", common.VisitHistory)
		commomn_router.POST("/user/FBToken", func(c *gin.Context) {
			UUID := c.Query("UUID")
			token := c.Query("token")
			db, err := database.Mariadb()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			defer db.Close()
			rtmsg := "Success"
			_, err = db.Exec("UPDATE user_info SET token = ? WHERE UUID = ?", token, UUID)
			if err != nil {
				rtmsg = "Failed"
				log.Fatal("insert into users error: ", err)
			}
			c.JSON(http.StatusCreated, gin.H{
				"rtmsg": rtmsg,
			})
		})
		commomn_router.GET("/user/userinfo", common.AppUserInfo)
		commomn_router.GET("/user/info",common.CommonUserInfo)
		commomn_router.POST("/user/change", func(c *gin.Context) {

			c.JSON(http.StatusCreated, gin.H{
				"rtmsg": "Success",
			})
		})
		commomn_router.DELETE("/user/change", func(c *gin.Context) {

			c.JSON(http.StatusAccepted, gin.H{
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
			user_admin_router.GET("/subgroup/lookup", common.SubGroupLookup)
			user_admin_router.PUT("/subgroup/group",func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"rtmsg": "Success",
			})
			})
			user_admin_router.PUT("/subgroup/user",func(c *gin.Context) {
				c.JSON(http.StatusCreated, gin.H{
					"rtmsg": "Success",
				})
			})
			user_admin_router.GET("/subgroup/group/all",common.GroupAll)//그룹 목록 전체 조회
			user_admin_router.GET("/subgroup/group/lookup/all",common.SubGroupAll)// (하위그룹관리)모든 하위관리자 그룹 정보 조회
			user_admin_router.GET("/subgroup/device/lookup/all", common.DeviceGroupLookUp)
			user_admin_router.GET("/subgroup/device/lookup/group", common.DeviceGroupLookUp1)
			user_admin_router.POST("/subgroup/device/add",common.DevcieGroupAdd) //디바이스 생성
			user_admin_router.DELETE("/subgroup/device/del",common.DeviceGroupDel) //디바이스 삭제
			user_admin_router.POST("/subgroup/authadd",common.GroupAuthAdd) //상위관리자 권한 부여
			user_admin_router.GET("/account",common.AdminAccountLook) //계정관리 조회
			user_admin_router.DELETE("/account",common.AdminAccounthDel) //계정 삭제
			user_admin_router.PUT("/account",common.AdminAccountPut) //계정 수정
			user_admin_router.GET("/wearabledevice/specificuserlook",common.AdminUserLook) //특정 사용자와 겹치는 사용자 조회
			user_admin_router.GET("/wearabledevice/specificuserlook/specificuserotheruser",common.AdminOtherUserLook) //특정 사용자 req 일 때 동선 겹치는 사용자 파악
			user_admin_router.GET("/wearabledevice",common.AdminDeviceLook) //웨어러블디바이스 조회
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
			user_admin_router.GET("/nfclog/lookup",common.AdminNFClog) // NFC 태그 기록
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
			user_admin_router.GET("/deviceMGMT",common.DeviceMT)

			user_admin_router.PUT("/accountMGMT",func(c *gin.Context){
				c.JSON(http.StatusCreated, gin.H{
					"rtmsg" : "Success",
				})
			})//자신 게정 수정
			user_admin_router.DELETE("/accountMGMT", func(c *gin.Context) {
				c.JSON(http.StatusAccepted,gin.H{
					"rtmsg" : "Success",
				})
			})// 자신 계정 삭제
			user_admin_router.GET("devicelog/lookup",common.DeivceLog)
			user_admin_router.GET("/dashboard/data-graph",common.DataGraph)
			user_admin_router.GET("/dashboard/data-trends",common.Dashboard)

			//작업중 user_admin_router.GET("/sendUser")
			//작업중 user_amdin_router.GET("/sendGroup")
			//작업중 user_admin_router.GET("/sendAll")
		}

	}
	return r
}

func main() {
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
