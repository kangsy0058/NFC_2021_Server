package common

import (
	"database/sql"
	"log"
	"net/http"
	"nfc_api/database"
	"os"

	"github.com/gin-gonic/gin"
)

type AppUserInfoModel struct {
	UUID        string `json:"UUID" example:"이용자1"`
	Email       string `json:"Email" example:"uuid@naver.com"`
	DisplayName string `json:"DisplayName" example:"이용자1"`
	PSN         string `json:"PSN" example:"12가34나"`
	WearableSN  string `json:"Wearable_SN" example:"wsn1111"`
	Is_Admin    int    `json:"Is_Admin" example:"0"`
}
type UserInfoModel struct {
	level int    `json:"level" example:"12가34나" `
	name  string `json:"name" example:"홍길동"`
	Token string `json:"Token" example:"aaaa.bbbb.cccc"`
	Email string `json:"Email" example:"uuid@naver.com"`
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

func AppUserInfo(c *gin.Context) {
	UUID := c.Query("UUID")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("Select UUID, email, displayname, PSN, wearable_SN, Is_admin FROM user_info where UUID = ?", UUID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var responseMessage AppUserInfoModel
	var tmp_PSN, tmp_WSN sql.NullString
	for rows.Next() {
		err := rows.Scan(&responseMessage.UUID, &responseMessage.Email, &responseMessage.DisplayName, &tmp_PSN, &tmp_WSN, &responseMessage.Is_Admin)
		if err != nil {
			log.Fatal(err)
		}
	}
	responseMessage.PSN = tmp_PSN.String
	responseMessage.WearableSN = tmp_WSN.String
	if !tmp_PSN.Valid {
		responseMessage.PSN = "정보없음"
	}
	if !tmp_WSN.Valid {
		responseMessage.WearableSN = "정보없음"
	}

	c.JSON(http.StatusOK, gin.H{
		"User_log": responseMessage,
	})
}

func CommonUserInfo(c *gin.Context) {
	Type := c.Query("type")

	//Level := 0
	//Name := "홍길동"
	//Token := "aaaa.bbbb.cccc"
	//Email := "uuid@naver.com"
	//responseMessage := UserInfoModel{Level, Name, Token, Email}

	if Type == "1" {
		c.JSON(http.StatusOK, gin.H{
			"rt": http.StatusOK,
			"data": gin.H{
				"level":     1,
				"name":      "홍길동1",
				"token":     "aaaa.bbbb.cccc",
				"approve":   1,
				"email":     "uuid1@naver.com",
				"groupCode": "0001",
				"address":   "경기도 화성시 17-1",
				"groupName": "그룹1",
			}})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"rt": http.StatusOK,
			"data": gin.H{
				"level": 0,
				"name":  "홍길동1",
				"token": "aaaa.bbbb.cccc",
			}})

	}
}

//if err != nil {
//	//cache에 데이터가 없는경우 db 확인
//	db, err := database.Mariadb()
//	if err != nil {
//		// can't connect database return status code 500
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return
//	}
//	defer db.Close()
//	c.JSON(http.StatusOK, gin.H{
//		"rt":   http.StatusOK,
//		"data": responseMessage,
//	})
//}

func UserLogin(c *gin.Context) {
	groupCode := "0001"
	Is_Admin := 1
	responseMessage := UserLoginModel{groupCode, Is_Admin}

	c.JSON(http.StatusOK, gin.H{
		"rt":   http.StatusOK,
		"data": responseMessage})
	c.Abort()
}

func CreateUserDevice(c *gin.Context) {
	Wearable_SN := c.Query("Wearable_SN")
	UUID := c.Query("UUID")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	//rtmsg := "Success"

	_, err = db.Exec("UPDATE user_info SET wearable_SN = ? WHERE UUID = ?", Wearable_SN, UUID)
	if err != nil {
		//	rtmsg = "Failed"
		log.Fatal("insert error: ", err)
	}

	c.JSON(http.StatusCreated, gin.H{
		// "rtmsg": rtmsg,
		"rtmsg": "Success",
	})
}

func DeleteUserDevice(c *gin.Context) {
	Wearable_SN := c.Query("Wearable_SN")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	// rtmsg := "Success"
	_, err = db.Exec("UPDATE user_info SET wearable_SN = NULL WHERE wearable_SN = ?", Wearable_SN)
	if err != nil {
		// 	rtmsg = "Failed"
		log.Fatal("delete error: ", err)
	}
	c.JSON(http.StatusAccepted, gin.H{
		// "rtmsg": rtmsg,
		"rtmsg": "Success",
	})
}

func CreateUserPsersonalNumber(c *gin.Context) {
	PSN := c.Query("PSN")
	PSN_img := c.Query("PSN_img")
	UUID := c.Query("UUID")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	//rtmsg := "Success"
	_, err = db.Exec("UPDATE user_info SET PSN = ?, PSN_img = ? WHERE UUID = ?", PSN, PSN_img, UUID)
	if err != nil {
		//	rtmsg = "Failed"
		log.Fatal("insert into users error: ", err)
	}
	c.JSON(http.StatusCreated, gin.H{
		// "rtmsg": rtmsg,
		"rtmsg": "Success",
	})
}

func DeleteUserPsersonalNumber(c *gin.Context) {
	PSN := c.Query("PSN")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	// rtmsg := "Success"
	_, err = db.Exec("UPDATE user_info SET PSN = NULL, PSN_img = NULL WHERE PSN = ?", PSN)
	if err != nil {
		// 	rtmsg = "Failed"
		log.Fatal("delete error: ", err)
	}
	c.JSON(http.StatusAccepted, gin.H{
		// "rtmsg": rtmsg,
		"rtmsg": "Success",
	})
}

func CreateFCMToken(c *gin.Context) {
	UUID := c.Query("UUID")
	token := c.Query("token")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	// rtmsg := "Success"
	_, err = db.Exec("UPDATE user_info SET token = ? WHERE UUID = ?", token, UUID)
	if err != nil {
		// rtmsg = "Failed"
		log.Fatal("insert into users error: ", err)
	}
	c.JSON(http.StatusCreated, gin.H{
		// "rtmsg": rtmsg,
		"rtmsg": "Success",
	})
}

func CreatePushChannel(c *gin.Context) {
	UUID := c.Query("UUID")
	topic := c.Query("topic")

	logFile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Print(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	log.Println(UUID, "adds topic", topic)

	c.JSON(http.StatusCreated, gin.H{
		"rtmsg": "Success",
	})
}

func DeletePushChannel(c *gin.Context) {
	UUID := c.Query("UUID")
	topic := c.Query("topic")

	logFile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Print(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	log.Println(UUID, "deletes topic", topic)
	c.JSON(http.StatusAccepted, gin.H{
		"rtmsg": "Success",
	})
}

func SignUp(c *gin.Context) {
	UUID := c.Query("UUID")
	Email := c.Query("Email")
	displayname := c.Query("displayname")
	PSN := c.Query("PSN")
	PSN_img := c.Query("PSN_img")
	WSN := c.Query("WSN")

	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into user_info(UUID,email,displayname,PSN,PSN_img,wearable_sn) values(?,?,?,?,?,?)", UUID, Email, displayname, PSN, PSN_img, WSN)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"rtmsg": "Success",
	})
}

func ChangeUserName(c *gin.Context) {
	UUID := c.Query("UUID")
	displayname := c.Query("displayname")

	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE user_info SET displayname = ? WHERE UUID = ?", displayname, UUID)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusAccepted, gin.H{
		"rtmsg": "Success",
	})
}
