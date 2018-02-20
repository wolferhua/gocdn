package lib

import (
	"fmt"
	"os"
	"encoding/json"
)


type Config struct {
	Buckets map[string] Bucket
}


func InitConfig() (conf Config) {
	file, err := os.Open("conf/conf.json")
	defer file.Close()
	if err !=nil {
		fmt.Println("load config error:"+err.Error())
		os.Exit(1);
	}
	decoder := json.NewDecoder(file)
	//conf = Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Buckets)
	fmt.Println("config load")
	return conf
}