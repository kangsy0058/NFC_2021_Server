package common

import (
	"fmt"
	"log"
	"net/http"
	"nfc_api/database"
	"time"

	"github.com/gin-gonic/gin"
)

func SubGroupLookup(c *gin.Context) {
	//UUID := "이용자1"
	//Email := "uuid@naver.com"
	//DisplayName := "이용자1"
	//Token := "aaaa.bbbb.cccc"
	//PSN := "12가34나"
	//PSN_img := "C:Users사용자이름PicturesMyImg1.jpg"
	//Is_Admin := 0
	//WearableSN := "wsn1111"
	//GroupCode := "0001"
	//GroupName := "그룹1"
	//Address := "경기도 화성서 17-1"
	//
	//responseSubGroup := UserSubGroupModel{UUID, Email, DisplayName, Token, PSN, PSN_img, Is_Admin, WearableSN}
	//responseGroupList := GroupListModel{GroupCode, GroupName, Address}
	groupCode := c.Query("groupCode")
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select u.UUID,u.email,u.displayname,u.token,u.psn,u.psn_img,u.Is_admin,u.wearable_sn,u.group_code,g.group_name,g.address,g.group_log from user_info as u join group_list as g on u.Group_code=g.Group_code where u.group_code=?", groupCode)

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

}

func CreateGroup(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"rtmsg": "Success",
	})
}

func CreateGroupUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"rtmsg": "Success",
	})
}

func GroupAll(c *gin.Context) {
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select Group_code,address,Group_name from group_list")
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
}

func SubGroupAll(c *gin.Context) {
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select u.UUID,u.email,u.displayname,u.token,u.is_admin,g.Group_code,g.address,g.Group_name from user_info as u join group_list as g on u.group_code=g.group_code")
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
}

func DeviceGroupLookUp(c *gin.Context) {
	groupCode := c.Query("groupCode")
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select *from kiosk_set where group_code=?", groupCode)
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

	//responseMessage := GroupDeviceModel{GroupCode, KioskSN, DetailPosition,BuildingName, Latitude, Longtitude}
	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}
func DeviceGroupLookUp1(c *gin.Context) {
	groupCode := c.Query("groupCode")
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select Group_code,building_name,detail_position,kiosk_sn,latitude,longitude from kiosk_set where group_code=?  ", groupCode)
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

	//responseMessage := GroupDeviceModel{GroupCode, KioskSN, DetailPosition,BuildingName, Latitude, Longtitude}
	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

func DevcieGroupAdd(c *gin.Context) {
	groupCode := c.Query("groupCode")
	kioskSN := c.Query("kioskSN")
	detailPosition := c.Query("detailPosition")
	buildingName := c.Query("buildingName")
	latitude := c.Query("latitude")
	longtitide := c.Query("longtitide")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT into kiosk_set (kiost_SN, Group_code, detail_position, building_name, latitude, longitude, kioskset_log) values (?,?,?,?,?,?,?) ", groupCode, kioskSN, detailPosition, buildingName, latitude, longtitide, time.Now())
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"rtmsg": "Success",
	})
}

func DeviceGroupDel(c *gin.Context) {
	groupCode := c.Query("groupCode")
	kioskSN := c.Query("kioskSN")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("Delete from kiosk_set where kiosk_SN = ? AND Group_code = ?", kioskSN, groupCode)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"rtmsg": "Success",
	})
}

func GroupAuthAdd(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{
		"rtmsg": "Success",
	})
}

func AdminAccountLook(c *gin.Context) {
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select u.Is_admin,u.UUID,u.displayname,u.email,g.Group_name,g.address from user_info as u join group_list as g on u.group_code=g.group_code")
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

	//for rows.Next(){
	//	rows.Scan(&data.UUID, &data.Email, &data.DisplayName, &data.Token, &data.PSN,
	//		&data.PSN_img, &data.Is_Admin, &data.WearableSN)
	//}

	c.JSON(http.StatusOK, gin.H{
		"data": results,
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
	})
	return
}

func AdminAccounthDel(c *gin.Context) {
	uuid := c.Query("uuid")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("Delete from user_info where UUID = ? ", uuid)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"rtmsg": "Success",
	})
}

func AdminAccountPut(c *gin.Context) {
	uuid := c.Query("uuid")
	email := c.Query("email")
	displayname := c.Query("displayname")
	token := c.Query("token")
	PSN := c.Query("PSN")
	PSN_img := c.Query("PSN_img")
	is_admin := c.Query("Is_admin")
	wearable_SN := c.Query("wearable_SN")
	Groupcode := c.Query("Group_code")

	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE user_info SET email = ?, displayname = ?, token = ?, PSN = ?, PSN_img = ?, Is_admin =?, wearable_SN = ?, Group_code = ? WHERE UUID = ?", email, displayname, token, PSN, PSN_img, is_admin, wearable_SN, Groupcode, uuid)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"rtmsg": "Success",
	})
}

func AdminUserLook(c *gin.Context) {
	wearable_SN := c.Query("wearable_SN")
	db, err := database.Mariadb()
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select latitude, longitude, building_name, date, time, group_code from user_log where wearable_sn=?", wearable_SN)
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

func AdminOtherUserLook(c *gin.Context) {
	wearable_SN := c.Query("wearable_SN")
	db, err := database.Mariadb()
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//wearableSN := "WSN1111"

	rows, err := db.Query("select latitude,longitude,building_name,time from user_log where building_name=(select building_name from user_log where wearable_sn= ? group by building_name)", wearable_SN)

	defer db.Close()
	fmt.Printf("", rows, 123)
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
		"data": gin.H{
			"werable_SN": wearable_SN,
			"log":        results,
		},
	})
	return
}

func AdminDeviceLook(c *gin.Context) {
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select *from user_info where wearable_sn=?", "wsn1111")
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

func DeviceMT(c *gin.Context) {
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT * FROM user_info where UUID=?", "이용자10")
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

}

func DataGraph(c *gin.Context) {
	db, err := database.Mariadb()
	//var data UserSubGroupModel
	if err != nil {
		// can't connect database return status code 500
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select date, avg(temp) from user_log where date >= date(subdate(now(),INTERVAL 8 DAY)) and date <= date(now()) Group by date;")
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

	//

	// rows1, err1 := db.Query("select DATE_FORMAT(date,"%Y-%u") weeks,avg(temp) from user_log where date >=date_add(curdate(),interval -56 day) and date <=date_format(curdate(),'%Y-%m-%d') GROUP BY weeks;")
	// defer db.Close()
	// cols1, err1 := rows1.Columns()
	// if err1 != nil {
	// 	return
	// }
	// data1 := make([]interface{}, len(cols1))

	// for i1, _ := range data1 {
	// 	var d []byte
	// 	data[i] = &d
	// }
	// results1 := make([]map[string]interface{}, 0)
	// for rows1.Next() {
	// 	err1 := rows1.Scan(data...)
	// 	if err1 != nil {
	// 		return
	// 	}
	// 	result1 := make(map[string]interface{})
	// 	for i1, item1 := range data1 {
	// 		result1[cols1[i1]] = string(*(item1.(*[]byte)))
	// 	}
	// 	results1 = append(results1, result1)
	// }

	//

	rows2, err2 := db.Query("select building_name,count(building_name) from user_log group by building_name order by count(building_name) desc LIMIT 6")
	defer db.Close()
	cols2, err2 := rows2.Columns()
	if err2 != nil {
		return
	}
	data2 := make([]interface{}, len(cols2))

	for i2, _ := range data2 {
		var d2 []byte
		data2[i2] = &d2
	}
	results2 := make([]map[string]interface{}, 0)
	for rows2.Next() {
		err2 := rows2.Scan(data2...)
		if err2 != nil {
			return
		}
		result2 := make(map[string]interface{})
		for i2, item2 := range data2 {
			result2[cols2[i2]] = string(*(item2.(*[]byte)))
		}
		results2 = append(results2, result2)
	}
	//fmt.Printf("", rows, 123)

	c.JSON(http.StatusOK, gin.H{
		"temperature": gin.H{
			"day": results,
			//"week": results1,
		},
		"visited_building": results2,
	})
}

func Dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"usedTerminalCount": gin.H{
			"now":    30,
			"change": 40,
		},
		"usedKioskCount": gin.H{
			"now":    20,
			"change": 30,
		},
		"avgTemperature": gin.H{
			"now":    36,
			"change": 37,
		},
		"buildingVisitCount": gin.H{
			"now":    10,
			"change": 20,
		},
	})
}

func CreateWearableDevice(c *gin.Context) {
	UUID := c.Query("target_UUID")
	wearable_SN := c.Query("displayname")

	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	_, err = db.Exec("insert into user_info(UUID,wearable_sn) values(?,?) ", UUID, wearable_SN)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"rtmsg": "Success",
	})
}

func ModifyWearableDevice(c *gin.Context) {
	UUID := c.Query("uuid")
	wearable_SN := c.Query("wearable_SN")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	_, err = db.Exec("update user_info set wearable_sn=? where UUID=?", wearable_SN, UUID)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"rtmsg": "Success",
	})
}

func DeleteWearableDevice(c *gin.Context) {
	wearable_SN := c.Query("wearable_SN")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	_, err = db.Exec("delete from user_info where wearable_sn=?", wearable_SN)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusAccepted, gin.H{
		"rtmsg": "Success",
	})
}

func ModifyUserDevice(c *gin.Context) {
	UUID := c.Query("uuid")
	email := c.Query("email")
	displayname := c.Query("displayname")
	wearable_SN := c.Query("wearable_SN")
	//token := c.Query("token")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	_, err = db.Exec("update user_info set email = ?, displayname = ? ,wearable_sn=? where UUID=?", email, displayname, wearable_SN, UUID)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"rtmsg": "Success",
	})
}

func ModifyUserAccount(c *gin.Context) {
	token := c.Query("token")
	email := c.Query("email")
	displayname := c.Query("displayname")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	_, err = db.Exec("update user_info set email=?, displayname=? where token=?", email, displayname, token)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"rtmsg": "Success",
	})
}

func DeleteUserAccount(c *gin.Context) {
	token := c.Query("token")
	db, err := database.Mariadb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	_, err = db.Exec("delete from user_info where token=?", token)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusAccepted, gin.H{
		"rtmsg": "Success",
	})
}
