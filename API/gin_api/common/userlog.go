package common

import (
	"log"
	"net/http"
	"nfc_api/database"

	"github.com/gin-gonic/gin"
)

type UserLogModel struct {
	IDX          int    `json:"IDX" example:"1" format:"int64"`
	Time         string `json:"Time" example:"15:30:30"`
	Date         string `json:"Date" example:"2021-06-30"`
	Temp         string `json:"temp" example:"36.5"`
	BuildingName string `json:"building_name" exmaple:"Hoseo"`
}

type SpecificUserModel struct {
	lat          int    `json:"lat" example:"36736" format:"int64" `
	lng          int    `json:"lng" example:"127074" format:"int64"`
	buildingName string `json:"building_name" example:"1공학관"`
	datetime     string `json:"datetime" example:"15:30:30"`
	groupCode    int    `json:"group_code" example:"0001" format:"int64"`
}

// Welcome godoc
// @Summary app working test
// @Description 테스트용 작성후 삭제예정
// @Tag common
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
	log.Println(start_date)
	log.Println(end_date)
	var responseMessage []UserLogModel
	for rows.Next() {
		var tmp UserLogModel
		err = rows.Scan(&tmp.IDX, &tmp.Time, &tmp.Date, &tmp.Temp, &tmp.BuildingName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(tmp.Date)
		tmp.Date = tmp.Date[:10]
		responseMessage = append(responseMessage, tmp)
	}
	c.JSON(http.StatusOK, gin.H{
		"User_log": responseMessage})
}

//특정 사용자 동선 조회
//작업중
//func AdminDeviceSK(c *gin.Context) {
//	uuid := "www1"
//	lat :=36736
//	lng :=127074
//	buildingName :="1공학관"
//	datetime :="15:30:30"
//	groupCode := 0001
//
//
//	c.JSON(http.StatusOK,gin.H{
//		"UUID:":uuid,
//	})
//}

func AdminNFClog(c *gin.Context) {
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select *from user_log")
	defer db.Close()
	cols, err := rows.Columns()
	if err != nil {
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
		"data": results,
	})
	return
}

func DeivceLog(c *gin.Context) {
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select *from user_log where Group_code=?", "041-31499-g1")
	defer db.Close()
	cols, err := rows.Columns()
	if err != nil {
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
		"data": results,
	})
	return
}
