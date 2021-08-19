package common

import (
	"database/sql"
	"log"
	"net/http"
	"nfc_api/database"

	"github.com/gin-gonic/gin"
)

var DB_l *sql.DB

type UserLogModel struct {
	IDX          int64   `json:"IDX" example:"1" format:"int64"`
	Time         string  `json:"Time" example:"15:30:30" format:"time"`
	Date         string  `json:"Date" example:"2021-06-30" format:"date"`
	Temp         float64 `json:"temp" example:"36.5" format:"float64"`
	BuildingName string  `json:"building_name" exmaple:"Hoseo"`
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
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT IDX, time, date, temp, building_name FROM user_log WHERE wearable_SN = ? AND date BETWEEN ? AND ?", Wearable_SN, start_date, end_date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var responseMessage []UserLogModel
	for rows.Next() {
		var tmp UserLogModel
		err = rows.Scan(&tmp.IDX, &tmp.Time, &tmp.Date, &tmp.Temp, &tmp.BuildingName)
		if err != nil {
			log.Fatal(err)
		}
		tmp.Date = tmp.Date[:10]
		responseMessage = append(responseMessage, tmp)
	}
	log.Println(responseMessage)
	c.JSON(http.StatusOK, gin.H{
		"User_log": responseMessage})
}
