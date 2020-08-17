package main

import (
	"io"
	"log"
	"os"
	"path"
	"rebilibili/src/models"
	"rebilibili/src/utils"
	"time"
)

//Recorder 开始录制
func Recorder(url, roomid string) {
	res, err := utils.Get(url, nil, nil)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		var filepath string = path.Join(basefilepath, roomid)
		if !exists(filepath) {
			os.Mkdir(filepath, os.ModePerm)
		}
		filefullpath := path.Join(filepath, time.Now().Format("20060102_150405")+".flv")
		out, err := os.Create(filefullpath)
		if err != nil {
			log.Println(err)
		}
		defer out.Close()

		_, err = io.Copy(out, res.Body)
		if tran {
			job := &models.TranscodeJob{}
			job.InputPath = filefullpath
			go func() {
				jobChan <- *job
			}()
		}
		if err != nil {
			log.Println(err)
		}
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
