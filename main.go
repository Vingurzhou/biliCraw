package main

import (
	"BiliCraw/model"
	"BiliCraw/util"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Initial_state struct {
}

func downloadSingleVideo(url, outPutDir string) {
	frontUrl := strings.Split(url, "?")[0]
	err := os.MkdirAll(outPutDir, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}

	htmlStr := util.BilibiliCraw(frontUrl, "GET")
	re := regexp.MustCompile(`<script>window.__playinfo__={.*?}</script><script>`)
	playInfo := re.FindAllString(htmlStr, -1)[0]
	playInfo = strings.TrimLeft(playInfo, "<script>window.__playinfo__=")
	playInfo = strings.TrimRight(playInfo, "</script><script>")
	var pi model.PlayInfo
	json.Unmarshal([]byte(playInfo), &pi)
	videoUrl := pi.Data.Dash.Video[0]
	audioUrl := pi.Data.Dash.Audio[0]
	fmt.Println(videoUrl, audioUrl)
}
func main() {
	url := "https://www.bilibili.com/video/BV1JE411g7XF"
	frontUrl := strings.Split(url, "?")[0]
	outPutDir := "./videos"
	err := os.MkdirAll(outPutDir, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}

	htmlStr := util.BilibiliCraw(frontUrl, "GET")
	re := regexp.MustCompile(`<script>window.__INITIAL_STATE__={.*?};`)
	initial_state := re.FindAllString(htmlStr, -1)[0]
	initial_state = strings.TrimLeft(initial_state, "<script>window.__INITIAL_STATE__=")
	initial_state = strings.TrimRight(initial_state, ";")
	var is model.InitialState
	json.Unmarshal([]byte(initial_state), &is)
	pages := is.VideoData.Pages
	for i := range pages {
		video_url := fmt.Sprintf("%s?p=%d", frontUrl, i+1)
		downloadSingleVideo(video_url, outPutDir)
		break
	}
}
