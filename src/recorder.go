package main

import (
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/xfrr/goffmpeg/transcoder"
)

//Recorder 开始录制
func Recorder(url, roomid string) {
	res, err := Get(url, nil, nil)
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
			go Transcode(filefullpath)
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

//Transcode 转码
func Transcode(filepath string) {
	trans := new(transcoder.Transcoder)
	err := trans.Initialize(filepath, path.Join(path.Dir(filepath), strings.TrimSuffix(path.Base(filepath), path.Ext(filepath))+".mp4"))
	if err != nil {
		log.Println(err)
	}
	trans.MediaFile().SetThreads(runtime.NumCPU())
	done := trans.Run(false)
	err = <-done
	os.Remove(filepath)
}
