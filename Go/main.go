package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	//创建目录
	_ = os.Mkdir("./hoyo_emoji", 0777)
}

type Emoticon struct {
	Name string        `json:"name"`
	Num  int           `json:"num"`
	List []interface{} `json:"list"`
}

// GetEmoticon 获取表情包列表数据
func GetEmoticon(Api string) ([]*Emoticon, error) {
	//request
	resp, err := http.Get(Api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//response
	var data []byte
	data, err = ioutil.ReadAll(resp.Body)
	//反序列化为json
	var jsonData map[string]interface{}
	_ = json.Unmarshal(data, &jsonData)
	//检查请求状态
	if jsonData["message"].(string) != "OK" {
		return nil, fmt.Errorf("request failure")
	}
	//获取表情包
	emoticonList := make([]*Emoticon, 0, 20)
	for _, item := range jsonData["data"].(map[string]interface{})["list"].([]interface{}) {
		//去掉无表情的元素
		if item.(map[string]interface{})["num"].(float64) == 0 {
			continue
		}
		//数据绑定到结构体
		var emoticon Emoticon
		temp, _ := json.Marshal(item)
		_ = json.Unmarshal(temp, &emoticon)
		emoticonList = append(emoticonList, &emoticon)
	}
	return emoticonList, err
}

func DownLoad(emoticon *Emoticon) error {
	//创建表情包目录
	dirPath := "./hoyo_emoji/" + emoticon.Name
	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}
	for _, emoji := range emoticon.List {
		//协程计数+1
		wg.Add(1)
		go func(fileName string, dirPath string, url string) {
			defer wg.Done()
			//请求数据
			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close()
			//读取数据
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				return
			}
			//保存路径
			filePath := dirPath + "/" + fileName + url[strings.LastIndex(url, "."):]
			f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				log.Println(err)
				return
			}
			defer f.Close()
			//写入数据
			_, err = f.Write(data)
			if err != nil {
				log.Println(err)
				return
			}
		}(emoji.(map[string]interface{})["name"].(string), dirPath, emoji.(map[string]interface{})["icon"].(string))
	}
	//同步
	wg.Wait()
	log.Printf("%s 表情包下载完成，共%d\n", emoticon.Name, emoticon.Num)
	return nil
}

func main() {
	start := time.Now()
	Api := "https://bbs-api-static.miyoushe.com/misc/api/emoticon_set?gids=1"
	emoticonList, err := GetEmoticon(Api)
	if err != nil {
		log.Println(err)
		return
	}
	for _, emoticon := range emoticonList {
		err = DownLoad(emoticon)
		if err != nil {
			log.Println(err)
			return
		}
	}
	end := time.Now()
	log.Println("处理用时：", end.Sub(start))
}
