package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"
	"nfc_api/common"
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
			// =========== 수정 필요 =============
			kiosk_router.GET("/sncheck/:werableSN", kiosk.CheckWearableSN)
			kiosk_router.POST("/userlog", kiosk.Userlog)
			//====================================
		}

		//Common API
		//app
		commomn_router := v1.Group("/common")
		//commomn_router.Use(firebaseauth.FirebaseAuthMiddleware())
		{
			// =========== 테스트 필요 ==========

			//====================================

			// =========== 수정 필요 =============
			commomn_router.POST("/user/change", common.CreatePushChannel)
			commomn_router.DELETE("/user/change", common.DeletePushChannel)
			commomn_router.GET("/user/login", common.UserLogin)
			commomn_router.GET("/user/info", common.CommonUserInfo)
			commomn_router.PATCH("/user/displayname", common.ChangeUserName)     //09.16 큰 상관은 없지만 UUID에 존재하지 않은 UUID 입력해도 Success나옵니다
			commomn_router.POST("/user/device", common.CreateUserDevice)         // 09.16 Success는 돌아오는데 DB값이 그대로입니다
			commomn_router.DELETE("/user/device", common.DeleteUserDevice)       ///09.16 Success는 돌아오는데 DB값이 그대로입니다
			commomn_router.POST("/user/pid", common.CreateUserPsersonalNumber)   //09.16 Success는 돌아오는데 DB값이 그대로입니다
			commomn_router.DELETE("/user/pid", common.DeleteUserPsersonalNumber) //09.16 Success는 돌아오는데 DB값이 그대로입니다
			commomn_router.POST("/user/FBToken", common.CreateFCMToken)          // 09.16  Success는 돌아오는데 DB값이 그대로입니다
			commomn_router.POST("/user/datainit", common.SignUp)                 //09.16  Error: socket hang up 나옵니다

			// ===================================

			// =========== 개발완료 =============
			commomn_router.GET("/userlog/visitHistory", common.VisitHistory) //09.16😀
			commomn_router.GET("/user/userinfo", common.AppUserInfo)         //09.16😀
			// ===================================
		}

		//kiosk_Admin API
		// kiosk에 통합 또는 이전 상의
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
			// =========== 테스트 필요 ==========

			// ===================================

			// =========== 수정 필요 =============
			user_admin_router.PUT("/subgroup/user", common.CreateGroupUser)
			user_admin_router.POST("/subgroup/authadd", common.GroupAuthAdd) //상위관리자 권한 부여
			user_admin_router.GET("/dashboard/data-trends", common.Dashboard)
			user_admin_router.GET("/dashboard/data-graph", common.DataGraph)
			user_admin_router.GET("/deviceMGMT", common.DeviceMT)
			user_admin_router.PUT("/accountMGMT", common.ModifyUserAccount)    //자신 게정 수정
			user_admin_router.DELETE("/accountMGMT", common.DeleteUserAccount) // 자신 계정 삭제
			user_admin_router.GET("devicelog/lookup", common.DeivceLog)

			user_admin_router.GET("/subgroup/device/lookup/all", common.DeviceGroupLookUp) ////파람값을 넣었는데 아무것도 안나옴
			user_admin_router.POST("/subgroup/device/add", common.DevcieGroupAdd)          //socket hang up                                   //디바이스 생성
			user_admin_router.DELETE("/subgroup/device/del", common.DeviceGroupDel)        //return은 오지만 데이터베이스에 반영이안됨                              //디바이스 삭제
			user_admin_router.DELETE("/account", common.AdminAccounthDel)                  //return은 오지만 데이터베이스에 반영이안됨                                               //계정 삭제
			user_admin_router.PUT("/subgroup/group", common.CreateGroup)                   //return은 오지만 데이터베이스에 반영이안됨
			user_admin_router.POST("/wearabledevice", common.CreateWearableDevice)         //socket hang up    //웨어러블디바이스 생성
			user_admin_router.PUT("/wearabledevice", common.ModifyWearableDevice)          //return은 오지만 수정이 안됩니다.    //웨어러블디바이스 수정
			user_admin_router.DELETE("/wearabledevice", common.DeleteWearableDevice)       //return은 오지만 삭제가 안됩니다. //웨어러블디바이스 삭제
			user_admin_router.PUT("/deviceMGMT", common.ModifyUserDevice)                  //return은 오지만 수정이 안됩니다.  데이터베이스에 반영이안됨     //하위관리자 디바이스 생성 수정
			user_admin_router.PUT("/account", common.AdminAccountPut)                      //return은 오지만 데이터베이스에 반영이안됨                                     //계정 수정

			// ===================================

			// =========== 개발완료 =============
			user_admin_router.GET("/subgroup/lookup", common.SubGroupLookup)
			user_admin_router.GET("/subgroup/group/all", common.GroupAll)           //그룹 목록 전체 조회
			user_admin_router.GET("/subgroup/group/lookup/all", common.SubGroupAll) // (하위그룹관리)모든 하위관리자 그룹 정보 조회
			user_admin_router.GET("/subgroup/device/lookup/group", common.DeviceGroupLookUp1)
			user_admin_router.GET("/account", common.AdminAccountLook)                                                 //계정관리 조회
			user_admin_router.GET("/wearabledevice/specificuserlook", common.AdminUserLook)                            //특정 사용자와 겹치는 사용자 조회
			user_admin_router.GET("/wearabledevice/specificuserlook/specificuserotheruser", common.AdminOtherUserLook) //특정 사용자 req 일 때 동선 겹치는 사용자 파악
			user_admin_router.GET("/wearabledevice", common.AdminDeviceLook)
			user_admin_router.GET("/nfclog/lookup", common.AdminNFClog) // NFC 태그 기록

			// ===================================
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
