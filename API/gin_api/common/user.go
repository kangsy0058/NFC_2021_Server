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
		c.AbortWithStatus(http.StatusInternalServerError)
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

type CommonUserInfoRequest struct {
	UUID string `json:"UUID"`
}
type CommonUserInfoResponse struct {
	Name      string `json:"name"`
	Token     string `json:"token"`
	Email     string `json:"email"`
	Approve   bool   `json:"approve"`
	GroupCode string `json:"groupCode"`
	Address   string `json:"address"`
	GroupName string `json:"groupName"`
}

func CommonUserInfo(c *gin.Context) {
	req := &CommonUserInfoRequest{}
	res := &CommonUserInfoResponse{}
	c.BindJSON(req)

	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = db.QueryRow("select u.is_admin,u.displayname,u.token,u.email,a.status,g.group_code,g.address,g.group_name from user_info u join Authority a on u.uuid=a.uuid join group_list g on a.group_code=g.group_code where u.UUID=?", req.UUID).Scan()

	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"rt":   http.StatusOK,
	})

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

type CreateUserDeviceRequest struct {
	WearableSn string `json:"Wearable_SN"`
	Uuid       string `json:"UUID"`
}

func CreateUserDevice(c *gin.Context) {
	req := &CreateUserDeviceRequest{}
	c.BindJSON(req)

	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	//rtmsg := "Success"

	dbResult, err := db.Exec("UPDATE user_info SET wearable_SN = ? WHERE UUID = ?", req.WearableSn, req.Uuid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"rtmsg": err,
		})
		return
	}
	updatedrow, err := dbResult.RowsAffected()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if updatedrow == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"rtmsg": "wearableSN or uuid error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		// "rtmsg": rtmsg,
		"rtmsg": "Success",
	})
}

type DeleteUserDeviceRequest struct {
	WearableSN string `json:"Wearable_SN"`
}

func DeleteUserDevice(c *gin.Context) {
	req := &DeleteUserDeviceRequest{}
	c.BindJSON(req)

	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	// rtmsg := "Success"
	dbResult, err := db.Exec("UPDATE user_info SET wearable_SN = NULL WHERE wearable_SN = ?", req.WearableSN)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"rtmsg": err,
		})
		return
	}
	updatedrow, err := dbResult.RowsAffected()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if updatedrow == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"rtmsg": "wearableSN error",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		// "rtmsg": rtmsg,
		"rtmsg": "Success",
	})
}

type CreateUserPsersonalNumberRequest struct {
	Psn  string `json:"PSN"`
	Uuid string `json:"UUID"`
}

func CreateUserPsersonalNumber(c *gin.Context) {
	req := &CreateUserPsersonalNumberRequest{}
	c.BindJSON(req)

	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"rtmsg": "dbconnection error",
		})
		return
	}
	defer db.Close()
	//rtmsg := "Success"
	dbResult, err := db.Exec("UPDATE user_info SET PSN = ? WHERE UUID = ?", req.Psn, req.Uuid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"rtmsg": err,
		})
		return
	}
	updatedrow, err := dbResult.RowsAffected()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if updatedrow == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"rtmsg": "uuid error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		// "rtmsg": rtmsg,
		"rtmsg": "Success",
	})
}

type DeleteUserPsersonalNumberRequest struct {
	Psn string `json:"PSN"`
}

func DeleteUserPsersonalNumber(c *gin.Context) {
	req := &DeleteUserPsersonalNumberRequest{}
	c.BindJSON(req)

	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	// rtmsg := "Success"
	dbResult, err := db.Exec("UPDATE user_info SET PSN = NULL, PSN_img = NULL WHERE PSN = ?", req.Psn)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"rtmsg": err,
		})
		return
	}
	updatedrow, err := dbResult.RowsAffected()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if updatedrow == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"rtmsg": "psn error",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		// "rtmsg": rtmsg,
		"rtmsg": "Success",
	})
}

type CreateFCMTokenRequest struct {
	UUID  string `json:"UUID"`
	Token string `json:"token"`
}

func CreateFCMToken(c *gin.Context) {

	req := &CreateFCMTokenRequest{}
	c.BindJSON(req)

	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	// rtmsg := "Success"
	dbResult, err := db.Exec("UPDATE user_info SET token = ? WHERE UUID = ?", req.Token, req.UUID)
	if err != nil {
		// rtmsg = "Failed"
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"rtmsg": "db connection error",
		})
		return
	}

	updatedrow, err := dbResult.RowsAffected()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if updatedrow == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"rtmsg": "uuid error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		// "rtmsg": rtmsg,
		"rtmsg": "Success",
	})
}

type CreatePushChannelRequest struct {
	UUID  string `json:"UUID"`
	Topic string `json:"topic"`
}

func CreatePushChannel(c *gin.Context) {
	req := &CreatePushChannelRequest{}
	c.BindJSON(req)

	logFile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Print(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	log.Println(req.UUID, "adds topic", req.Topic)

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

type ChangeUserNameRequest struct {
	UUID        string `json:"UUID"`
	DisplayName string `json:"displayname"`
}

func ChangeUserName(c *gin.Context) {
	req := &ChangeUserNameRequest{}
	c.BindJSON(req)

	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	dbResult, err := db.Exec("UPDATE user_info SET displayname = ? WHERE UUID = ?", req.DisplayName, req.UUID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"rtmsg": err,
		})
		return
	}
	updatedrow, err := dbResult.RowsAffected()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if updatedrow == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"rtmsg": "uuid error",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"rtmsg": "Success",
	})
}
