package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	Buckets map[string]Bucket
	Deny  string
	Mimes   map[string]string
}

func InitConfig() (conf Config) {
	//基础配置
	file, err := os.Open("conf/conf.json")
	defer file.Close()
	if err != nil {
		fmt.Println("load config error:" + err.Error())
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	//conf = Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	//mime
	mimes, err := ioutil.ReadFile("conf/mime.types")

	if err != nil {
		fmt.Println("load config error:" + err.Error())
		os.Exit(1)
	}
	mimecontent := string(mimes)
	mimecontent = strings.TrimSpace(mimecontent)
	mimelist := strings.Split(mimecontent, ";")

	conf.Mimes = make(map[string]string)
	for _, item := range mimelist {
		items := strings.Fields(item)
		if len(items) > 1 {
			exts := items[1:]
			for _, ext := range exts {
				conf.Mimes[ext] = items[0]
			}
		}
	}

	fmt.Println("config load")
	//fmt.Println(conf)
	return conf
}
