package main

import (
	"lib"
	"net/http"
	"strconv"
)

var config lib.Config

func init() {
	config = lib.InitConfig()
}

func main() {
	h := lib.Handler{
		Conf: config,
	}
	http.ListenAndServe(config.Host+":"+strconv.Itoa(config.Port), h)
}
