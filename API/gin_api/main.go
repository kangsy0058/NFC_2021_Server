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
			commomn_router.POST("/user/change", common.CreatePushChannel)   //기능회의필요
			commomn_router.DELETE("/user/change", common.DeletePushChannel) //기능회의필요
			commomn_router.GET("/user/login", common.UserLogin)             // app 로그인기능으로사용 admin data없이
			commomn_router.GET("/user/info", common.CommonUserInfo)         // admin level 설정후 다시 확인

			// ===================================

			// =========== 개발완료 =============
			commomn_router.GET("/userlog/visitHistory", common.VisitHistory)                   //09.16😀
			commomn_router.GET("/user/userinfo", common.AppUserInfo)                           //09.16😀
			commomn_router.POST("/user/FBToken", common.CreateFCMToken)                        // 확인완료
			commomn_router.PATCH("/user/displayname", common.ChangeUserName)                   //확인완료
			commomn_router.POST("/user/device", common.CreateUserDevice)                       //확인완료
			commomn_router.DELETE("/user/device", common.DeleteUserDevice)                     //확인완료
			commomn_router.POST("/user/pid", common.CreateUserPsersonalNumber)                 //확인완료
			commomn_router.DELETE("/user/pid", common.DeleteUserPsersonalNumber)               //확인완료
			commomn_router.POST("/user/psersonalimage/:uuid", common.CreateUserPsersonalImage) // image 업로드 폴더 도커에서 따로매핑 해줘야함
			commomn_router.POST("/user/datainit", common.SignUp)                               //확인완료
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
			user_admin_router.PUT("/subgroup/user", common.CreateGroupUser) // 유저생성방식 변경후 작성
			user_admin_router.PUT("/subgroup/group", common.CreateGroup)    // 유저생성방식 변경후 작성

			user_admin_router.GET("/dashboard/data-trends", common.Dashboard) // 안돼
			user_admin_router.GET("/dashboard/data-graph", common.DataGraph)  // 작동하지만 일단위 이상함 그리고 일일 방문자와 불가능했던 사람 말고는 통계가 의미있는지?

			user_admin_router.POST("/wearabledevice", common.CreateWearableDevice)   //sql문 wearable 테이블에 만드는거로 수정    //웨어러블디바이스 생성
			user_admin_router.PUT("/wearabledevice", common.ModifyWearableDevice)    //return은 오지만 수정이 안됩니다.    //웨어러블디바이스 수정
			user_admin_router.DELETE("/wearabledevice", common.DeleteWearableDevice) //return은 오지만 삭제가 안됩니다. //웨어러블디바이스 삭제 // 미쳤나

			// ===================================

			// =========== 개발완료 =============
			user_admin_router.GET("/subgroup/lookup", common.SubGroupLookup)
			user_admin_router.GET("/subgroup/group/all", common.GroupAll)           //그룹 목록 전체 조회
			user_admin_router.GET("/subgroup/group/lookup/all", common.SubGroupAll) // (하위그룹관리)모든 하위관리자 그룹 정보 조회
			user_admin_router.GET("/subgroup/device/lookup/group", common.DeviceGroupLookUp1)
			user_admin_router.GET("/account", common.AdminAccountLook)                                                 //계정관리 조회
			user_admin_router.GET("/wearabledevice/specificuserlook", common.AdminUserLook)                            //특정 사용자와 겹치는 사용자 조회
			user_admin_router.GET("/wearabledevice/specificuserlook/specificuserotheruser", common.AdminOtherUserLook) //특정 사용자 req 일 때 동선 겹치는 사용자 파악
			user_admin_router.GET("/nfclog/lookup", common.AdminNFClog)                                                // NFC 태그 기록

			user_admin_router.POST("/subgroup/authadd", common.GroupAuthAdd)        //상위관리자 권한 부여
			user_admin_router.GET("devicelog/lookup/:groupcode", common.DeivceLog)  //그룹별 로그
			user_admin_router.POST("/subgroup/device/add", common.DevcieGroupAdd)   //디바이스 생성
			user_admin_router.DELETE("/subgroup/device/del", common.DeviceGroupDel) //디바이스 삭제
			user_admin_router.DELETE("/account", common.AdminAccounthDel)           //계정 삭제
			user_admin_router.PUT("/account", common.AdminAccountPut)               //계정 수정
			user_admin_router.GET("/wearabledevice", common.AdminDeviceLook)

			// ===================================
			//작업중 user_admin_router.GET("/sendUser")
			//작업중 user_amdin_router.GET("/sendGroup")
			//작업중 user_admin_router.GET("/sendAll")

			// ===== 삭제 =================
			user_admin_router.GET("/deviceMGMT", common.DeviceMT)
			user_admin_router.PUT("/accountMGMT", common.ModifyUserAccount)                //자신 게정 수정
			user_admin_router.DELETE("/accountMGMT", common.DeleteUserAccount)             // 자신 계정 삭제
			user_admin_router.GET("/subgroup/device/lookup/all", common.DeviceGroupLookUp) ////파람값을 넣었는데 아무것도 안나옴
			user_admin_router.PUT("/deviceMGMT", common.ModifyUserDevice)                  //return은 오지만 수정이 안됩니다.  데이터베이스에 반영이안됨     //하위관리자 디바이스 생성 수정
			//===============================

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
