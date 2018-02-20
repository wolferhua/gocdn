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

	reg := regexp.MustCompile(`(?U)^/([^/]+)(/.*?)`)
	all := reg.FindAllStringSubmatch(path,-1)
	if len(all)<1 {
		slf.h404(w,r)
		return
	}

	//仓库
	bucketname := all[0][1]
	//包含版本路径
	ver_path := all[0][2]

	filename := ver_path

	reg = regexp.MustCompile(`(?U)^/v\.([^/]+)(/.*?)`)
	all = reg.FindAllStringSubmatch(ver_path,-1)
	fmt.Println(all)
	ver := "default"
	min := false
	if len(all) > 0 {
		ver = all[0][1]
		reg = regexp.MustCompile(`\.min`)
		min = reg.MatchString(ver)
		filename = all[0][2]
	}else{

	}




	fmt.Println(bucketname)
	fmt.Println(ver)
	fmt.Println(min)
	fmt.Println(filename)


	//fmt.Println(all)
	//fmt.Println(all[0][1])
	//fmt.Println(all[1])

	fmt.Fprintln(w, r.URL.Path)
}




func (slf Handler) h404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintln(w,"<h1>404 Not Found!</h1>")
}