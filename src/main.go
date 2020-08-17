package main

import (
	"flag"
	"os"
	"path/filepath"
	"rebilibili/src/models"
	"strings"
	"sync"
	"time"
)

const baseaddress string = "https://api.live.bilibili.com/"

var wg sync.WaitGroup
var roomid string
var tran bool
var basefilepath string
var jobChan = make(chan models.TranscodeJob, 100)

func main() {
	wd, _ := os.Getwd()
	flag.StringVar(&roomid, "r", "", "房间号，多房间用逗号分隔")
	flag.StringVar(&basefilepath, "v", wd, "文件保存位置")
	flag.BoolVar(&tran, "t", false, "转码MP4")
	flag.Parse()
	if roomid != "" {
		roomids := strings.Split(strings.Replace(roomid, "，", ",", -1), ",")
		for _, v := range roomids {
			if v != "" {
				wg.Add(1)
				if tran {
					go Worker(jobChan)
					pat := filepath.Join(basefilepath, v, "*.flv")
					files, _ := filepath.Glob(pat)
					for _, f := range files {
						job := &models.TranscodeJob{}
						job.InputPath = f
						go func() {
							jobChan <- *job
						}()
					}
				}
				go start(v)
			}
		}
		wg.Wait()
	}
}

func start(roomid string) {
	var islive bool = false
	for !islive {
		islive = Checklive(roomid)
		if !islive {
			sleeptime, _ := time.ParseDuration("1m")
			time.Sleep(sleeptime)
		}
	}
	url := Getliveurl(roomid)
	Recorder(url, roomid)
	start(roomid)
}
