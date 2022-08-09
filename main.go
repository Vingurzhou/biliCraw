package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func crawSingleUrl(url, method string) string {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err.Error())
	}

	req.Header.Add("authority", "www.bilibili.com")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("cookie", "_uuid=43D02BAA-AF89-A0EC-1431-ED0EBCEBCF3889828infoc; buvid3=39E18E2C-96EF-48BF-ADD2-D55C1184486A167619infoc; b_nut=1634091990; rpdid=|(k|kY|m~Ylu0J'uYJR|mkJ|); LIVE_BUVID=AUTO8316362008528464; video_page_version=v_old_home; CURRENT_BLACKGAP=0; blackside_state=0; buvid4=6EDBAE22-C7DB-B291-B1C0-8B88A345CF5207846-022012016-R7WDlws6U28Dlc/vTb0ONw%3D%3D; CURRENT_QUALITY=80; nostalgia_conf=-1; i-wanna-go-back=-1; buvid_fp_plain=undefined; SESSDATA=a6ecf4ac%2C1666582698%2Cd4f37%2A41; bili_jct=1ab37acb5554f9fea3cbdb523aa35318; DedeUserID=387457550; DedeUserID__ckMd5=6607490eb35f89b8; sid=b2ni1lcd; buvid_fp=7c34590617755999f365f58aba6b2ba3; b_ut=5; PVID=2; fingerprint3=0675fd3b80a7d60bd57b04f322764717; fingerprint=969215c5ec57bd1fff253932d45863d1; hit-dyn-v2=1; bp_video_offset_387457550=691676175150350359; CURRENT_FNVAL=4048; theme_style=light; b_lsid=641041866_18281638065; b_timer=%7B%22ffp%22%3A%7B%22333.788.fp.risk_39E18E2C%22%3A%221828163820F%22%2C%22333.1193.fp.risk_39E18E2C%22%3A%2218281638E1A%22%7D%7D")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("sec-ch-ua", "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	return string(body)
}
func main() {
	url := "https://www.bilibili.com/video/BV1JE411g7XF"
	err := os.MkdirAll("./videos", os.ModePerm)
	if err != nil {
		panic(err.Error())
	}

	htmlStr := crawSingleUrl(url, "GET")
	re := regexp.MustCompile(`<script>window.__INITIAL_STATE__={.*?};`)
	initial_state := re.FindAllString(htmlStr, -1)[0]
	initial_state = strings.TrimLeft(initial_state, "<script>window.__INITIAL_STATE__=")
	initial_state = strings.TrimRight(initial_state, ";")
	fmt.Println(initial_state)
}
