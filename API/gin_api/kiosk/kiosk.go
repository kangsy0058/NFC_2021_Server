package kiosk

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//type welcomeModel struct {
//	ID   int    `json:"id" example:"1" format:"int64"`
//	Name string `json:"name" example:"account  name"`
//}
//
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
	//SN := c.Param("KioskSN")
	user_stat := true
	//if err := c.ShouldBindJSON(&SN); err !=nil{
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"rt": 400,
	//		"response": "Parameter Check",
	//	})
	//
	//	log.Print(err.Error())
	//
	//	return
	//}
	usercheckMessage := user_stat

	c.JSON(http.StatusOK, gin.H{"response": usercheckMessage})

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

	c.JSON(http.StatusCreated, gin.H{"response": true})
}
