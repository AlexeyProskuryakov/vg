package ytvd

import (
	. "github.com/kkdai/youtube"
	"log"
	"fmt"
)
const DEFAULT_CUR_DIR = "/Users/alesha/tmp/ytvd"
type VideoFile struct {
	Info VideoFileInfo
	
}
type VideoFileInfo struct {
	ContentFileName string
	Url             string
}

var currentDir string

func init() {
	
	log.Println("video will downloads to dir=", currentDir)
}

func DownloadVideo(videoId string) (*VideoFileInfo, error) {
	y := NewYoutube()
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoId)
	y.DecodeURL(url)
	
	fileName := fmt.Sprintf("%s/%s.3gp",DEFAULT_CUR_DIR, videoId)
	err := y.StartDownload(fileName, 1)
	if err == nil{
		result := VideoFileInfo{}
		result.ContentFileName = fileName
		result.Url = url
		return &result, nil
	}
	return nil, err
}
