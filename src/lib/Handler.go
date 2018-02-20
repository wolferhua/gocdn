package lib

import (
	"net/http"
	"fmt"
)

type Handler struct {
	
}

func (slf Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Implement the Handle interface.")
}