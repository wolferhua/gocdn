package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

type Config struct {
	Buckets map[string]Bucket
	Mimes map[string]string
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
	fmt.Println(conf.Buckets)
	fmt.Println("config load")
	//mime
	mimes, err  := os.Open("conf/mime.types")
	defer mimes.Close()



	return conf
}
