package ytvd

import (
	. "github.com/kkdai/youtube"
	"path/filepath"
	"log"
	"fmt"
)
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
	
	fileName, _ := filepath.Abs(filepath.Dir(videoId))
	err := y.StartDownload(fileName)
	if err == nil{
		result := VideoFileInfo{}
		result.ContentFileName = fileName
	}
	return nil, err
}
