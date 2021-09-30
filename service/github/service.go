package github

import (
	"fmt"
	"time"

	"github.com/SparkElf/blog-server/config"
	"github.com/SparkElf/blog-server/model"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

func Init() {
	go startTask()
}
func GetUserTotalStars(username string) (total int, err error) {
	client := resty.
		New().
		SetTimeout(5 * time.Second)

	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	resp, err := client.R().
		SetAuthToken(config.APIToken).
		Get(url)
	value := gjson.Get(resp.String(), "#.stargazers_count").Array()
	for i := range value {
		total += int(value[i].Int())
	}
	println("GetValues:", value)
	return total, err
}
func startTask() {
	//每隔1秒向github请求一次数据，超时时间为5秒
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()
	for range t.C {
		//println(time.Now().Format("2006-01-02 15:04:05"))
		v, e := GetUserTotalStars(config.Username)
		if e != nil {
			println(e.Error())
		}
		if model.U.GetTotalStars() != v {
			model.U.SetTotalStars(v)
		}
		//println(time.Now().Format("2006-01-02 15:04:05"))
	}

}
