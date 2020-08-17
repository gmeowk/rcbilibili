package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"rebilibili/src/models"
	"rebilibili/src/utils"
)

//Getliveurl 获取直播地址
func Getliveurl(roomid string) string {
	res, err := utils.Get(baseaddress+"room/v1/Room/playUrl?cid="+roomid+"&quality=4&platform=web", nil, nil)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var liveurl models.Liveurl
	err1 := json.Unmarshal(body, &liveurl)
	if err1 != nil {
		log.Println(err)
	}
	return liveurl.Data.Durl[0].URL
}
