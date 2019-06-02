package main

import (
	"./HttpResponse"
	"./Models"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	dbhost     = "127.0.0.1:3306"
	dbusername = "root"
	dbpassword = "root"
	dbname     = "bilibili"
)

/*
  获取sql.DB对象
*/
func GetDB() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", dbusername, dbpassword, dbhost, dbname))
	CheckErr(err)
	return db
}

func CheckErr(err error) {
	println(err)
}

func main() {

	isOk := make(chan bool, 1)
	var minAid, maxAid, chooseType int
	fmt.Println("请输入minAid:")
	fmt.Scanln(&minAid)
	fmt.Println("请输入maxAid:")
	fmt.Scanln(&maxAid)
	//lastRecord := findLastRecord()
	fmt.Printf("当前输入开始aid:%d;结束aid:%d \n", minAid, maxAid)

	chooseType = 0
	for chooseType == 0 {
		fmt.Println("请选择操作方式：【1】->采集视频，【2】->更新标签")
		fmt.Scanln(&chooseType)
		if chooseType == 1 || chooseType == 2 {
			break
		} else {
			chooseType = 0
		}
	}

	//if lastRecord.Aid > minAid {
	//	minAid = lastRecord.Aid + 1
	//	if lastRecord.Aid > maxAid {
	//		maxAid = lastRecord.Aid + 1
	//	}
	//	if minAid > maxAid {
	//		fmt.Printf("数据自动纠正minAid:%d,maxAid:%d \n", minAid, maxAid)
	//	} else {
	//		fmt.Printf("数据纠正后：开始aid:%d;结束aid:%d \n", minAid, maxAid)
	//	}
	//}
	rc := make(chan bool, maxAid-minAid+1)
	for i := minAid; i <= maxAid; i++ {
		if chooseType == 1 {
			go rsyncVideoInfo(i, isOk, maxAid, rc)
		} else {
			go rsyncVideoTag(i, isOk, maxAid, rc)
		}

		<-rc
		//if !<-isOk {
		//	break
		//}
	}
	fmt.Println("操作完成")

}

func rsyncVideoInfo(aid int, isOk chan bool, maxAid int, rc chan bool) {

	record := findByAid(aid)
	if record.Aid == 0 {
		res := getVideoInfo(aid)

		if res.Code == 0 && res.Data.Aid > 0 {
			result := transformVideoInfoToModels(res.Data)
			insert(result)
			rc <- true
			isOk <- true
			fmt.Println(strconv.Itoa(aid) + ":数据导入完成！")
		} else {
			rc <- false
			if res.Code == 0 {
				isOk <- false
				fmt.Println(strconv.Itoa(aid) + ":请求失败，IP被封！")
			} else {
				fmt.Println(strconv.Itoa(aid) + ":" + res.Message)
			}
		}
	} else {
		rc <- true
		isOk <- true
		fmt.Println(strconv.Itoa(record.Aid) + ":【" + record.Title + "】已导入！")
	}

	if aid >= maxAid {
		fmt.Println(strconv.Itoa(aid) + ":范围溢出！")
		isOk <- false
	}

}

func transformVideoInfoToModels(videoInfo HttpResponse.VideoInfoResponseData) Models.Bilibili {

	fmt.Println(strconv.Itoa(videoInfo.Aid) + ":准备解析数据！")

	data := Models.Bilibili{}
	data.Aid = videoInfo.Aid
	data.Title = videoInfo.Title
	data.Url = "https://www.bilibili.com/video/av" + strconv.Itoa(videoInfo.Aid)
	data.Duration = videoInfo.Duration
	data.View = videoInfo.Stat.View
	data.Danmaku = videoInfo.Stat.Danmaku
	data.Reply = videoInfo.Stat.Reply
	data.Favorite = videoInfo.Stat.Favorite
	data.Coin = videoInfo.Stat.Coin
	data.Share = videoInfo.Stat.Share
	data.Like = videoInfo.Stat.Like
	data.NowRank = videoInfo.Stat.NowRank
	data.HisRank = videoInfo.Stat.HisRank
	data.Keywords = ""
	data.ActionTag = ""
	data.EmotionTag = ""
	data.SceneTag = ""
	data.StarTag = ""
	data.DialogTag = ""
	data.UpdateCount = 1
	data.UpdatedAt = time.Now().String()
	data.CreatedAt = time.Now().String()

	fmt.Println(strconv.Itoa(videoInfo.Aid) + ":数据解析完成！")
	return data
}

func insert(bilibili Models.Bilibili) sql.Result {
	println("准备导入数据：")
	sql := "INSERT INTO `scripts`.`bilibili`(`aid`, `title`, `url`, `duration`, `view`, `danmaku`, `reply`, `favorite`, `coin`, `share`, `like`, `now_rank`, `his_rank`, `keywords`, `action_tag`, `emotion_tag`, `scene_tag`, `star_tag`, `dialog_tag`, `update_count`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	db := GetDB()
	defer db.Close()
	result, err := db.Exec(sql, bilibili.Aid, bilibili.Title, bilibili.Url, bilibili.Duration, bilibili.View, bilibili.Danmaku, bilibili.Reply, bilibili.Favorite, bilibili.Coin, bilibili.Share, bilibili.Like, bilibili.NowRank, bilibili.HisRank, bilibili.Keywords, bilibili.ActionTag, bilibili.EmotionTag, bilibili.SceneTag, bilibili.StarTag, bilibili.DialogTag, bilibili.UpdateCount)
	if err != nil {
		println("数据导入失败")
		println(err)
	}
	println("数据导入完成！")
	return result
}

func findLastRecord() Models.Bilibili {
	sql := "select * from bilibili order by aid desc limit 1"
	db := GetDB()
	defer db.Close()
	record := Models.Bilibili{}
	db.QueryRow(sql).Scan(
		&record.Id,
		&record.Aid,
		&record.Title,
		&record.Url,
		&record.Duration,
		&record.View,
		&record.Danmaku,
		&record.Reply,
		&record.Favorite,
		&record.Coin,
		&record.Share,
		&record.Like,
		&record.NowRank,
		&record.HisRank,
		&record.Keywords,
		&record.ActionTag,
		&record.EmotionTag,
		&record.SceneTag,
		&record.StarTag,
		&record.DialogTag,
		&record.UpdateCount,
		&record.UpdatedAt,
		&record.CreatedAt)
	return record
}

/**
 *
 */
func findByAid(aid int) Models.Bilibili {
	sql := "select * from bilibili where aid=" + strconv.Itoa(aid) + " order by id desc limit 1"
	db := GetDB()
	defer db.Close()
	record := Models.Bilibili{}
	db.QueryRow(sql).Scan(
		&record.Id,
		&record.Aid,
		&record.Title,
		&record.Url,
		&record.Duration,
		&record.View,
		&record.Danmaku,
		&record.Reply,
		&record.Favorite,
		&record.Coin,
		&record.Share,
		&record.Like,
		&record.NowRank,
		&record.HisRank,
		&record.Keywords,
		&record.ActionTag,
		&record.EmotionTag,
		&record.SceneTag,
		&record.StarTag,
		&record.DialogTag,
		&record.UpdateCount,
		&record.UpdatedAt,
		&record.CreatedAt)
	return record
}

func getVideoInfo(aid int) HttpResponse.VideoInfoResponse {

	requestUrl := "https://api.bilibili.com/x/web-interface/view?aid=" + strconv.Itoa(aid)
	println("准备采集数据：" + requestUrl)
	time.Sleep(time.Duration(10))

	client := http.Client{}

	//提交请求
	request, err := http.NewRequest("GET", requestUrl, nil)

	//增加header选项
	request.Header.Add("Referer", "https://www.bilibili.com/")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	request.Header.Set("Connection", "keep-alive")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	result := HttpResponse.VideoInfoResponse{}
	json.Unmarshal(body, &result)

	println("数据采集完成！" + result.Data.Title)
	return result

}

func rsyncVideoTag(aid int, isOk chan bool, maxAid int, rc chan bool) {
	str := ""
	record := findByAid(aid)
	if record.Aid > 0 {
		fmt.Println("准备获取视频标签：" + strconv.Itoa(aid))
		videoTagInfo := getVideoTagInfo(aid)
		for i := 0; i < len(videoTagInfo.Data); i++ {
			str += videoTagInfo.Data[i].TagName + ","
		}
		str = string(str[0 : len(str)-1])
		fmt.Println(strconv.Itoa(aid) + ":视频标签为：" + str)
		record.Keywords = str
		fmt.Println(strconv.Itoa(aid) + "：准备同步标签：" + str)
		updateVideoTagByAid(aid, &record)
		fmt.Println(strconv.Itoa(aid) + ":同步完成：" + str)
	}
	rc <- true
	if aid >= maxAid {
		isOk <- false
	} else {
		isOk <- true
	}
}

func updateVideoTagByAid(aid int, bilibili *Models.Bilibili) sql.Result {
	sql := "update bilibili set keywords='" + bilibili.Keywords + "' where aid=?"
	fmt.Println(sql)
	db := GetDB()
	defer db.Close()
	result, _ := db.Exec(sql, strconv.Itoa(aid))
	return result
}

func getVideoTagInfo(aid int) HttpResponse.VideoTagInfoResponse {

	url := "https://api.bilibili.com/x/tag/archive/tags?aid=" + strconv.Itoa(aid)
	time.Sleep(time.Duration(10))

	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)

	//增加header选项
	request.Header.Add("Referer", "https://www.bilibili.com/")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	request.Header.Set("Connection", "keep-alive")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	result := HttpResponse.VideoTagInfoResponse{}
	json.Unmarshal(body, &result)
	return result

}
