package kiosk

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type welcomeModel struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account  name"`
}

type kioskUserCheckModel struct {
	WearableSN string `json:"wearableSN" example:"wsn1111"`
	IsUser     bool   `json:"isuser" example:"true"`
}

// Welcome godoc
// @Summary kiosk working test
// @Description 테스트용 작성후 삭제예정
// @id hello
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /v1/kiosk/welcome/{name} [get]
// @Success 200 {object} welcomeModel
func WelcomeApi(c *gin.Context) {
	name := c.Param("name")
	message := name + " is very handsome"
	welcomeMessage := welcomeModel{1, message}

	c.JSON(http.StatusOK, gin.H{"message": welcomeMessage})
}

// Weaable check godoc
// @Summary check Wearable SN
// @Description Wearable SN를 받아 사용하는 유저가 존재하는지 확인하는 기능
// @id userCheck
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param sn path string true "Wearable Serial Number"
// @Router /v1/kiosk/checksn/{sn} [get]
// @Success 200 {object} kioskUserCheckModel
func CheckWearableSN(c *gin.Context) {
	wearable := c.Param("sn")
	user_stat := true

	usercheckMessage := kioskUserCheckModel{wearable, user_stat}

	c.JSON(http.StatusOK, gin.H{"message": usercheckMessage})
}
