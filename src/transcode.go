package main

import (
	"log"
	"os"
	"path"
	"rebilibili/src/models"
	"runtime"
	"strings"

	"github.com/xfrr/goffmpeg/transcoder"
)

//Worker ..
func Worker(jobChan <-chan models.TranscodeJob) {
	for job := range jobChan {
		Transcode(job)
	}
}

//Transcode 转码
func Transcode(job models.TranscodeJob) {
	trans := new(transcoder.Transcoder)
	outputPath := path.Join(path.Dir(job.InputPath), strings.TrimSuffix(path.Base(job.InputPath), path.Ext(job.InputPath))+".mp4")
	err := trans.Initialize(job.InputPath, outputPath)
	if err != nil {
		log.Println(err)
	}
	trans.MediaFile().SetThreads(runtime.NumCPU())
	trans.MediaFile().SetPreset("ultrafast")
	done := trans.Run(false)
	err = <-done
	os.Remove(job.InputPath)
}
