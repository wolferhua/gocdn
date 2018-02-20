package main

import (
	"net/http"
	"lib"
)

var conf lib.Config

func init() {
	conf = lib.InitConfig()
}

func main() {
	var h lib.Handler
	http.ListenAndServe(":80", h)
}
