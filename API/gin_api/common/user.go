package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nfc_api/database"
)


type UserInfoModel struct {
	level          int     `json:"level" example:"12가34나" `
	name   string `json:"name" example:"홍길동"`
	Token string `json:"Token" example:"aaaa.bbbb.cccc"`
	Email string `json:"Email" example:"uuid@naver.com"`

}
type UserLoginModel struct {
	groupCode  string `json:"Group_Code" example:"0001"`
	Is_Admin int	`json:"Is_Admin" example:"1"`
}
type UserSubGroupModel struct {
	UUID string `json:"UUID" example:"이용자1"`
	Email string `json:"Email" example:"uuid@naver.com"`
	DisplayName string `json:"DisplayName" exmaple:"이용자1"`
	Token string `json:"Token" example:"aaaa.bbbb.cccc"`
	PSN   string     `json:"PSN" example:"12가34나" `
	PSN_img string `json:"PSN_Img" example:"C:Users사용자이름PicturesMyImg1.jpg"`
	Is_Admin     int `json:"Is_Admin" example:"0" format:"int64" `
	WearableSN   string `json:"Wearable_SN" example:"wsn1111"`
}
type GroupListModel struct {
	GroupCode string `json:"Group_Code" example:"0001"`
	GroupName string `json:"Group_Name" example:"그룹1"`
	Address string `json:"Address" example:"경기도 화성시 17-1"`
}
type GroupDeviceModel struct {
	GroupCode string `json:"Group_Code" example:"0001"`
	KioskSN string `json:"Kiosk_SN" exmaple:"KSN1111"`
	DetailPosition string `json:"detail_postision" example:"충청남도 아산시 배방읍"`
	BuildingName string `json:"building_name" example:"1공힉관"`
	Latitude float64 `json:"latitude" example:"36.738"`
	Longtitude float64 `json:"Longtitude" example:"127.074"`
}

//대시보드 작업중
//func Dashboard(c *gin.Context) {
//	now := 40
//	change := 10
//	usedTerminalCount :=
//
//	c.JSON(http.StatusOK, gin.H{
//		"data" : "usedTerminalCount:",
//	})
//}

func UserInfo(c *gin.Context) {
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
				"email":     "uuid1@naver.com",
				"groupCode": "0001",
				"address":   "경기도 화성시 17-1",
				"groupName": "그룹1",
			}})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"rt":   http.StatusOK,
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



func UserLogin(c *gin.Context)  {
	groupCode :="0001"
	Is_Admin := 1
	responseMessage := UserLoginModel{groupCode, Is_Admin}

	c.JSON(http.StatusOK, gin.H{
		"rt" : http.StatusOK,
		"data":responseMessage})
	c.Abort()
}


//admin
func SubGroupLookup(c *gin.Context)  {
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
	responseGroupList := GroupListModel{GroupCode,GroupName,Address}
	c.JSON(http.StatusOK, gin.H{
		"data" : gin.H{
			"subGroup" : responseSubGroup,
			"GroupList" : responseGroupList,
		},
		})

}

func DeviceGroupLookUp(c *gin.Context)  {
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows,err := db.Query("SELECT * FROM kiosk_set")
	defer db.Close()
	cols, err := rows.Columns()
	if err != nil{
		return
	}
	data := make([]interface{}, len(cols))

	for i, _ := range data {
		var d []byte
		data[i] = &d
	}
	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		err := rows.Scan(data...)
		if err != nil {
			return
		}
		result := make(map[string]interface{})
		for i, item := range data {
			result[cols[i]] = string(*(item.(*[]byte)))
		}
		results = append(results, result)
	}


	//responseMessage := GroupDeviceModel{GroupCode, KioskSN, DetailPosition,BuildingName, Latitude, Longtitude}
	c.JSON(http.StatusOK,gin.H{
		"data" : results,
	})
}

func DevcieGroupAdd(c *gin.Context){
	//groupCode := c.Query("groupCode")
	//kioskSN := c.Query("kioskSN")
	//detailPosition := c.Query("detailPosition")
	//buildingName := c.Query("buildingName")
	//latitude := c.Query("latitude")
	//longtitide := c.Query("longtitide")


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

func AdminAccountLook(c *gin.Context){
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows,err := db.Query("SELECT * FROM user_info")
	defer db.Close()
	cols, err := rows.Columns()
	if err != nil{
		return
	}
	data := make([]interface{}, len(cols))

	for i, _ := range data {
		var d []byte
		data[i] = &d
	}
	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		err := rows.Scan(data...)
		if err != nil {
			return
		}
		result := make(map[string]interface{})
		for i, item := range data {
			result[cols[i]] = string(*(item.(*[]byte)))
		}
		results = append(results, result)
	}

	//for rows.Next(){
	//	rows.Scan(&data.UUID, &data.Email, &data.DisplayName, &data.Token, &data.PSN,
	//		&data.PSN_img, &data.Is_Admin, &data.WearableSN)
	//}

	c.JSON(http.StatusOK, gin.H{
		"data" : results,
		//"data" : gin.H{
		//		"UUID" : "이용자9",
		//		"email" : "uuid5@gmail.com",
		//		"displayname" : "이용자9",
		//		"token" : "aaaa.bbbb.cccc",
		//		"PSN" : "12가34나",
		//		"PNS_img" : "C:Users사용자이름PicturesMyImg9.jpg",
		//		"Is_admin" : 1,
		//		"werable_SN" : "wsn1119",
		//		},
		"valid" : "true",
	})
	return
}


func AdminAccounthDel(c *gin.Context){
	c.JSON(http.StatusAccepted, gin.H{
		"rtmsg":"Success",
	})
}

func AdminAccountPost(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{
		"rtmsg":"Success",
	})
}

func AdminUserLook(c *gin.Context){
	db, err := database.Mariadb()
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows,err := db.Query("select latitude, longitude, building_name, date, time, group_code from user_log where wearable_sn=?","wsn1111")
	defer db.Close()
	cols, err := rows.Columns()
	if err != nil{
		return
	}
	data := make([]interface{}, len(cols))

	for i, _ := range data {
		var d []byte
		data[i] = &d
	}
	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		err := rows.Scan(data...)
		if err != nil {
			return
		}
		result := make(map[string]interface{})
		for i, item := range data {
			result[cols[i]] = string(*(item.(*[]byte)))
		}
		results = append(results, result)
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : results,
	})
	return
}

func Dashboard(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"usedTerminalCount" : gin.H{
			"now" : 30,
			"change" : 40,
		},
		"usedKioskCount" : gin.H{
			"now" : 20,
			"change" : 30,
		},
		"avgTemperature" : gin.H{
			"now" : 36,
			"change" : 37,
		},
		"buildingVisitCount" : gin.H{
			"now" : 10,
			"change" : 20,
		},
	})
}