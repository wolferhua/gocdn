package main

import (
	"lib"
	"net/http"
)

var config lib.Config

func init() {
	config = lib.InitConfig()
}

func main() {
	h := lib.Handler{
		Conf: config,
	}
	http.ListenAndServe(":80", h)
}
