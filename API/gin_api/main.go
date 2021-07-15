package main

import (
<<<<<<< HEAD
	"fmt"
	"log"
=======
<<<<<<< HEAD
=======
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
>>>>>>> Haeil
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939
	"net/http"
	"nfc_api/common"
	_ "nfc_api/docs"
	"nfc_api/firebaseauth"
	"nfc_api/kiosk"
<<<<<<< HEAD
	"time"

	"database/sql"
=======
<<<<<<< HEAD
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
=======
>>>>>>> Haeil
)

var DB *sql.DB

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
<<<<<<< HEAD
		commomn_router.POST("/device-add", func(c *gin.Context) {
			Wearable_SN := c.Query("Wearable_SN")
			UUID := c.Query("UUID")
			_, err := DB.Exec("insert into user_info (wearable_SN, UUID) values (?,?) ", Wearable_SN, UUID)
			if err != nil {
				log.Fatal("insert into users error: ", err)
			}
			c.JSON(http.StatusCreated, gin.H{
				"rtmsg": "Success",
			})
		})
		commomn_router.DELETE("/device-del", func(c *gin.Context) {
			Wearable_SN := c.Query("Wearable_SN")
=======
		commomn_router.POST("/user/device-add", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"rt":    http.StatusCreated,
				"rtmsg": "Success",
			})
		})
		commomn_router.DELETE("/user/device-del", func(c *gin.Context) {
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939
			c.JSON(http.StatusAccepted, gin.H{
				"test1": Wearable_SN,
				"rtmsg": "Success",
			})
		})
		commomn_router.POST("/user/pid-add", func(c *gin.Context) {
<<<<<<< HEAD
			PSN := c.Query("PSN")
			PSN_img := c.Query("PSN_img")
			UUID := c.Query("UUID")
=======
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939
			c.JSON(http.StatusCreated, gin.H{
				"test1": PSN,
				"test2": PSN_img,
				"test3": UUID,
				"rtmsg": "Success",
			})
		})
		commomn_router.DELETE("/user/pid-del", func(c *gin.Context) {
<<<<<<< HEAD
			PSN := c.Query("PSN")
=======
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939
			c.JSON(http.StatusAccepted, gin.H{
				"test1": PSN,
				"rtmsg": "Success",
			})
		})
		commomn_router.GET("/userlog/visitHistory", common.VisitHistory)
		commomn_router.POST("/user/FBToken", func(c *gin.Context) {
<<<<<<< HEAD
			token := c.Query("token")
=======
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939
			c.JSON(http.StatusCreated, gin.H{
				"test1": token,
				"rtmsg": "Success",
			})
		})
		commomn_router.GET("/user/userInfo", common.UserInfo)
		commomn_router.POST("/user/change", func(c *gin.Context) {
<<<<<<< HEAD
			Push := c.Query("Push_info Push")
=======
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939
			c.JSON(http.StatusCreated, gin.H{
				"test1": Push,
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
		}

=======
			user_admin_router.GET("/subgroup/lookup",common.SubGroupLookup)
			user_admin_router.GET("/subgroup/device/lookup/all",common.DeviceGroupLookUp)
			user_admin_router.GET("/subgroup/device/lookup/group",common.DeviceGroupLookUp)
			user_admin_router.POST("/subgroup/device/add",common.DevcieGroupAdd)
			user_admin_router.DELETE("/subgroup/device/del",common.DeviceGroupDel)
			user_admin_router.POST("/subgroup/authadd",common.GroupAuthAdd)
		}
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939
	}
	return r
}

func main() {
<<<<<<< HEAD
	r := setupRouter()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//////////
	db, err := sql.Open("mysql", "root:hoseolab420@tcp(210.119.104.207:3306)/hoseo")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	DB = db
	common.DB = db
	common.DB_l = db

	fmt.Println("connect success", db)
	defer db.Close()
	//////////
	////////////GET TEST

	////////////
=======
	// Router setup
	r := setupRouter()
	//Server start
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939
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
