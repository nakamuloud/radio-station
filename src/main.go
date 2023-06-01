package main

// 入力 : バケット名､番組名, 追加するURL
// 処理内容 : 1. バケット名/番組名/rss.xmlを取得｡なければ新規作成
//           2. rss.xmlをパースして、新規URLを追加
//           3. rss.xmlを上書き保存
// rssをパースする処理
import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

// channel要素の構造体
type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Items   []Item  `xml:"item"`
}

// item要素の構造体
type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

// RSSの構造体
type RSS struct {
	Channel Channel `xml:"channel"`
}

// RSSの構造体を返す関数
func getRss(url string) (RSS, error) {
	// RSSの構造体
	var rss RSS
	// RSSのURLからRSSを取得
	resp, err := http.Get(url)
	if err != nil {
		return rss, err
	}
	// RSSのBodyを取得
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rss, err
	}
	// RSSのBodyをパース
	err = xml.Unmarshal(byteArray, &rss)
	if err != nil {
		return rss, err
	}
	return rss, nil
}

// RSSの構造体から文字列を返す関数
func (rss *RSS) toString() string {
	// RSSの文字列
	var rssString string
	// Channel要素の文字列を作成
	rssString += fmt.Sprintf("Title: %s\n", rss.Channel.Title)
	rssString += fmt.Sprintf("Link: %s\n", rss.Channel.Link)
	rssString += fmt.Sprintf("Items:%s\n","")
	// Item要素の文字列を作成
	for _, item := range rss.Channel.Items {
		rssString += fmt.Sprintf("\tTitle: %s\n", item.Title)
		rssString += fmt.Sprintf("\tLink: %s\n", item.Link)
	}
	return rssString
}

func main() {
	xml_url := flag.String("xml_url", "https://feeds.megaphone.fm/TBS4550274867", "RSS xml url")
	add_file_url :=flag.String("add_file_url","","add file url")
	flag.Parse()
	// RSSのURL
	rssUrl := "https://feeds.megaphone.fm/TBS4550274867"
	// RSSをパース
	rss, err := getRss(rssUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(rss.toString())
	//GCSにアップロードする処理
	//GCSからRSSを取得する処理

}
func uploadToGcs(){
	credentialFilePath := "~/Downloads/key.json"

}
