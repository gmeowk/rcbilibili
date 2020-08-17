package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"rebilibili/src/models"
	"rebilibili/src/utils"
)

//Checklive 判断是否开播
func Checklive(roomid string) bool {
	res, err := utils.Get(baseaddress+"room/v1/Room/get_info?id="+roomid, nil, nil)
	if err != nil {
		log.Println(err)
		return Checklive(roomid)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var roomInfo models.RoomInfo
	err1 := json.Unmarshal(body, &roomInfo)
	if err1 != nil {
		log.Println(err)
	}
	if roomInfo.Data.LiveStatus == 1 {
		return true
	}
	return false
}
