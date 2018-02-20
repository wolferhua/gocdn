package main

import (
	"net/http"
	"lib"
)

func main() {
	var h lib.Handler
	http.ListenAndServe(":80", h)
}
