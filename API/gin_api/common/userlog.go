package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nfc_api/database"
)

type UserLogModel struct {
	IDX          int     `json:"IDX" example:"1" format:"int64"`
	Time         string `json:"Time" example:"15:30:30"`
	Date         string `json:"Date" example:"2021-06-30"`
	Temp         string `json:"temp" example:"36.5"`
	BuildingName string `json:"building_name" exmaple:"Hoseo"`
}

type SpecificUserModel struct {
	lat 	int `json:"lat" example:"36736" format:"int64" `
	lng		int `json:"lng" example:"127074" format:"int64"`
	buildingName string `json:"building_name" example:"1공학관"`
	datetime string `json:"datetime" example:"15:30:30"`
	groupCode int `json:"group_code" example:"0001" format:"int64"`
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
	IDX := 1
	Time := "15:30:15"
	Date := "2021-06-30"
	Temp := "36.5"
	BuildingName := "1공학관"
	responseMessage := UserLogModel{IDX, Time,Date,Temp,BuildingName}

	c.JSON(http.StatusOK, gin.H{"User_log": responseMessage})
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

func AdminNFClog(c *gin.Context){
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows,err := db.Query("select *from user_log")
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

func DeivceLog(c *gin.Context)  {
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows,err := db.Query("select *from user_log where Group_code=?","041-31499-g1")
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