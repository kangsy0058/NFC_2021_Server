package kiosk

import (
<<<<<<< HEAD
	"net/http"
	"nfc_api/database"
	"nfc_api/redisinit"

=======
>>>>>>> Haeil
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

<<<<<<< HEAD
//type welcomeModel struct {
//	ID   int    `json:"id" example:"1" format:"int64"`
//	Name string `json:"name" example:"account  name"`
=======
type UserLogModel struct {
	KioskSN    string   `json:"Kiosk_SN" exmaple:"KSN1111"`
	WearableSN string   `json:"Wearable_SN" example:"wsn1111"`
	Time       time.Time `json:"time" example:"03:14:18" foramt:"time"`
	Date       time.Time`json:"date" example:"2021-05-16"`
	Temp       float64  `json:"temp" example:"36.5" format:"float64"`
}

//type UserCheckModel struct {
//	WearableSN string `json:"wearableSN" example:"wsn1111"`
//	IsUser     bool   `json:"isuser" example:"true"`
>>>>>>> Haeil
//}

type UserCheckModel struct {
	WearableSN string `json:"wearableSN" example:"wsn1111"`
	IsUser     bool   `json:"isuser" example:"true"`
}

// Weaable check godoc
// @Summary check Wearable SN
// @Description Wearable SN를 받아 사용하는 유저가 존재하는지 확인하는 기능
// @Tag Kiosk
// @id userCheck
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param sn path string true "Wearable Serial Number"
// @Router /v1/kiosk/checksn/{wearableSN} [get]
// @Success 200 {object} UserCheckModel
func CheckWearableSN(c *gin.Context) {
<<<<<<< HEAD
	SN := c.Param("werableSN")

	// return model에 sn 값 추가
	usercheck := UserCheckModel{WearableSN: SN}

	//find status in cache
	status, err := redisinit.GetUserData(SN)
	if err != nil {
		//cache에 데이터가 없는경우 db 확인
		db, err := database.Mariadb()
		if err != nil {
			// can't connect database return status code 500
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// 결과값용 변수 필요하지만 사용하지않음
		var res string
		// DB에서 데이터 확인후 cache에 저장한후 다음과정 진행
		err = db.QueryRow("SELECT UUID FROM user_info WHERE wearable_SN = ?", SN).Scan(&res)
		if err != nil {
			status = "false"
		} else {
			status = "true"
		}
		// 결과 cache에 저장
		redisinit.SetUserData(SN, status)
	}
=======
	KioskSN := c.Query("KioskSN")
	user_stat := true
	if KioskSN == ""{
		c.JSON(http.StatusBadRequest,gin.H{
			"rt": http.StatusBadRequest,
			"response": "Parameter Check",
		})

		log.Print("Parameter Null")

		return
	}
	response := user_stat

	c.JSON(http.StatusOK, gin.H{"response": response})
>>>>>>> Haeil

	// status check  "true" is user "false" is not user
	if status == "true" {
		usercheck.IsUser = true
	} else if status == "false" {
		usercheck.IsUser = false
	}
	c.JSON(http.StatusOK, gin.H{"res": usercheck})
}

// Welcome godoc
// @Summary kiosk working test
// @Description 테스트용 작성후 삭제예정
// @Tag Kiosk
// @id hello
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /v1/kiosk/welcome/{name} [get]
// @Success 200 {object} welcomeModel
func Userlog(c *gin.Context) {
	//name := c.Param("name")
	//message := name + " is very handsome"
	//welcomeMessage := welcomeModel{1, message}
	KioskSN := c.Query("KioskSN")
	WearableSN := c.Query("WearableSN")
	Time := c.GetTime("Time")
	Date := c.GetTime("date")
	Temp := c.GetFloat64("temp")
	response := UserLogModel{KioskSN,WearableSN,Time,Date, Temp}

<<<<<<< HEAD
	c.JSON(http.StatusCreated, gin.H{"response": true})
=======
	c.JSON(http.StatusCreated, gin.H{"response": response})
>>>>>>> Haeil
}
