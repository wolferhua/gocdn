package lib

import (
	"net/http"
	"fmt"
	"regexp"
)

type Handler struct {
	
}

func (slf Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	//分析path
	///bucket/version/filename
	// /hashmap/v.1/test/1.png  //正常文件
	// /hashmap/v.1.min/test/1.png  //压缩

	reg := regexp.MustCompile(`(?U)^/([^/]+)(/.*)`)
	all := reg.FindAllStringSubmatch(path,-1)
	fmt.Println(all)
	if len(all)<1 {
		slf.h404(w,r)
		return
	}

	//仓库
	bucket := all[0][1]
	//包含版本路径
	verpath := all[1][0]

	reg = regexp.MustCompile(`(?U)^/(v.[^/]+)(/.*)`)
	all = reg.FindAllStringSubmatch(verpath,-1)
	fmt.Println(all)
	ver := "default"
	min := false
	if len(all) > 1 {
		ver = all[0][1]
		reg = regexp.MustCompile("\\.min")
		min = reg.MatchString(ver)
	}else{
	}



	fmt.Println(bucket)
	fmt.Println(verpath)


	fmt.Println(ver)
	fmt.Println(min)
	fmt.Println(all)
	//fmt.Println(all[0][1])
	//fmt.Println(all[1])

	fmt.Fprintln(w, r.URL.Path)
}




func (slf Handler) h404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintln(w,"<h1>404 Not Found!</h1>")
}