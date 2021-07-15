package common

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB_l *sql.DB

type UserLogModel struct {
	IDX          int    `json:"IDX" example:"1" format:"int64"`
	Time         string `json:"Time" example:"15:30:30"`
	Date         string `json:"Date" example:"2021-06-30"`
	Temp         string `json:"temp" example:"36.5"`
	BuildingName string `json:"building_name" exmaple:"Hoseo"`
}

// Welcome godoc
// @Summary app working test
// @Description 테스트용 작성후 삭제예정
// @Tags common
// @id Wearable_SN
// @Accept  json
// @Produce  json
// @Param name path string true "Wearable_SN"
// @Router /v1/common/visitHistory/{Wearable_SN} [get]
// @Success 200 {object} UserLogModel
func VisitHistory(c *gin.Context) {
	Wearable_SN := c.Query("Wearable_SN")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	IDX := 1
	Time := "15:30:15"
	Date := "2021-06-30"
	Temp := "36.5"
	BuildingName := "HOSEO"
	responseMessage := UserLogModel{IDX, Time, Date, Temp, BuildingName}

	c.JSON(http.StatusOK, gin.H{
		"test1":    Wearable_SN,
		"test2":    start_date,
		"test3":    end_date,
		"User_log": responseMessage})
}
