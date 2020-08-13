package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//Liveurl 直播地址
type Liveurl struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		CurrentQuality     int      `json:"current_quality"`
		AcceptQuality      []string `json:"accept_quality"`
		CurrentQn          int      `json:"current_qn"`
		QualityDescription []struct {
			Qn   int    `json:"qn"`
			Desc string `json:"desc"`
		} `json:"quality_description"`
		Durl []struct {
			URL        string `json:"url"`
			Length     int    `json:"length"`
			Order      int    `json:"order"`
			StreamType int    `json:"stream_type"`
			P2PType    int    `json:"p2p_type"`
		} `json:"durl"`
	} `json:"data"`
}

//Getliveurl 获取直播地址
func Getliveurl(roomid string) string {
	res, err := Get("https://api.live.bilibili.com/room/v1/Room/playUrl?cid="+roomid+"&quality=4&platform=web", nil, nil)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var liveurl Liveurl
	err1 := json.Unmarshal(body, &liveurl)
	if err1 != nil {
		log.Println(err)
	}
	return liveurl.Data.Durl[0].URL
}
