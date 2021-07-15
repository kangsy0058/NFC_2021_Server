package common

import (
	"database/sql"
<<<<<<< HEAD
	"log"
=======
	"fmt"
	"github.com/gin-gonic/gin"
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939
	"net/http"

	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
var DB *sql.DB
=======
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939

type UserInfoModel struct {
	PSN        string `json:"PSN" example:"12가34나" `
	WearableSN string `json:"Wearable_SN" example:"wsn1111"`
	Is_Admin   int    `json:"Is_Admin" example:"1"`
}
type UserLoginModel struct {
	groupCode string `json:"Group_Code" example:"0001"`
	Is_Admin  int    `json:"Is_Admin" example:"1"`
}
type UserSubGroupModel struct {
	UUID        string `json:"UUID" example:"이용자1"`
	Email       string `json:"Email" example:"uuid@naver.com"`
	DisplayName string `json:"DisplayName" exmaple:"이용자1"`
	Token       string `json:"Token" example:"aaaa.bbbb.cccc"`
	PSN         string `json:"PSN" example:"12가34나" `
	PSN_img     string `json:"PSN_Img" example:"C:Users사용자이름PicturesMyImg1.jpg"`
	Is_Admin    int    `json:"Is_Admin" example:"0" format:"int64" `
	WearableSN  string `json:"Wearable_SN" example:"wsn1111"`
}
type GroupListModel struct {
	GroupCode string `json:"Group_Code" example:"0001"`
	GroupName string `json:"Group_Name" example:"그룹1"`
	Address   string `json:"Address" example:"경기도 화성시 17-1"`
}
type GroupDeviceModel struct {
	GroupCode      string  `json:"Group_Code" example:"0001"`
	KioskSN        string  `json:"Kiosk_SN" exmaple:"KSN1111"`
	DetailPosition string  `json:"detail_postision" example:"충청남도 아산시 배방읍"`
	BuildingName   string  `json:"building_name" example:"1공힉관"`
	Latitude       float64 `json:"latitude" example:"36.738"`
	Longtitude     float64 `json:"Longtitude" example:"127.074"`
}

<<<<<<< HEAD
=======
func db()  {
	db, err := sql.Open("mysql", "root:hoseolab420@tcp(210.119.104.207:3306)/hoseo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939
func UserInfo(c *gin.Context) {
	UUID := c.Query("UUID")
	rows, err := DB.Query("Select PSN, wearable_SN, Is_admin FROM user_info where UUID = ?", UUID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var PSN, wearable_SN string
	var Is_admin int
	for rows.Next() {
		err := rows.Scan(&PSN, &wearable_SN, &Is_admin)
		if err != nil {
			log.Fatal(err)
		}
	}
	responseMessage := UserInfoModel{PSN, wearable_SN, Is_admin}

	c.JSON(http.StatusOK, gin.H{
		"User_log": responseMessage,
	})
}

func UserLogin(c *gin.Context) {
	groupCode := "0001"
	Is_Admin := 1
	responseMessage := UserLoginModel{groupCode, Is_Admin}

	c.JSON(http.StatusOK, gin.H{
		"rt":   http.StatusOK,
		"data": responseMessage})
	c.Abort()
}

//admin
func SubGroupLookup(c *gin.Context) {
	UUID := "이용자1"
	Email := "uuid@naver.com"
	DisplayName := "이용자1"
	Token := "aaaa.bbbb.cccc"
	PSN := "12가34나"
	PSN_img := "C:Users사용자이름PicturesMyImg1.jpg"
	Is_Admin := 0
	WearableSN := "wsn1111"
	GroupCode := "0001"
	GroupName := "그룹1"
	Address := "경기도 화성서 17-1"

	responseSubGroup := UserSubGroupModel{UUID, Email, DisplayName, Token, PSN, PSN_img, Is_Admin, WearableSN}
	responseGroupList := GroupListModel{GroupCode, GroupName, Address}
	c.JSON(http.StatusOK, gin.H{
		"subGroup":  responseSubGroup,
		"GroupList": responseGroupList,
		//"rt" : http.StatusOK,
	})

}

func DeviceGroupLookUp(c *gin.Context) {
	GroupCode := "0001"
	KioskSN := "KSN1111"
	DetailPosition := "충청남도 아산시 배방읍"
	BuildingName := "1공학관"
	Latitude := 36.738
	Longtitude := 127.074

	responseMessage := GroupDeviceModel{GroupCode, KioskSN, DetailPosition, BuildingName, Latitude, Longtitude}
	c.JSON(http.StatusOK, gin.H{
		"data": responseMessage,
	})
}
<<<<<<< HEAD
=======

func DevcieGroupAdd(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{
		"rtmsg":"Success",
	})
}

func DeviceGroupDel(c *gin.Context)  {
	c.JSON(http.StatusAccepted, gin.H{
		"rtmsg":"Success",
	})
}

func GroupAuthAdd(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{
		"rtmsg":"Success",
	})
}
>>>>>>> 95f215cc4cbba44dbce37a874071ff731cda8939
