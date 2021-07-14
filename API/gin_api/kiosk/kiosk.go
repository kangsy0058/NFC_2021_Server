package kiosk

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

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
//}


// Weaable check godoc
// @Summary check Wearable SN
// @Description Wearable SN를 받아 사용하는 유저가 존재하는지 확인하는 기능
// @Tag Kiosk
// @id userCheck
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param sn path string true "Wearable Serial Number"
// @Router /v1/kiosk/checksn/{sn} [get]
// @Success 200 {object} UserCheckModel
func CheckWearableSN(c *gin.Context) {
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
func PutUserlog(c *gin.Context) {
	//name := c.Param("name")
	//message := name + " is very handsome"
	//welcomeMessage := welcomeModel{1, message}
	KioskSN := c.Query("KioskSN")
	WearableSN := c.Query("WearableSN")
	Time := c.GetTime("Time")
	Date := c.GetTime("date")
	Temp := c.GetFloat64("temp")
	response := UserLogModel{KioskSN,WearableSN,Time,Date, Temp}

	c.JSON(http.StatusCreated, gin.H{"response": response})
}


