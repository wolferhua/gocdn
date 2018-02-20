package lib

import (
	"fmt"
	"net/http"
	"regexp"
)

type Handler struct {
	Conf Config
}

func (slf Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	//分析path
	///bucket/version/filename
	// /hashmap/v.1/test/1.png  //正常文件
	// /hashmap/v.1.min/test/1.png  //压缩

	reg := regexp.MustCompile(`(?U)^/([^/]+)(/.*?)`)
	all := reg.FindAllStringSubmatch(path, -1)
	if len(all) < 1 {
		slf.halt(w, r, 404)
		return
	}

	//仓库
	bucketname := all[0][1]

	//判断仓库是否存在
	_, ok := slf.Conf.Buckets[bucketname]
	if !ok {
		slf.halt(w, r, 404)
		return
	}
	var bucket Bucket = slf.Conf.Buckets[bucketname]

	//包含版本路径
	ver_path := all[0][2]

	filename := ver_path

	reg = regexp.MustCompile(`(?U)^/v\.([^/]+)(/.*?)`)
	all = reg.FindAllStringSubmatch(ver_path, -1)
	fmt.Println(all)
	ver := "default"
	min := false
	if len(all) > 0 {
		ver = all[0][1]
		reg = regexp.MustCompile(`\.min`)
		min = reg.MatchString(ver)
		filename = all[0][2]
	} else {

	}

	//文件对象
	bf := BucketFile{
		Bucket{
			bucket.Name,
			bucket.IsLocal,
			bucket.Root,
		},
		filename,
		ver,
		min,
	}

	if bucket.IsLocal {
		slf.local(bf, w, r)
	} else {
		slf.remote(bf, w, r)
	}

	return
	//fmt.Println(bucketname)
	//fmt.Println(bucket)
	//fmt.Println(ver)
	//fmt.Println(min)
	//fmt.Println(filename)

	//fmt.Println(all)
	//fmt.Println(all[0][1])
	//fmt.Println(all[1])

	//fmt.Fprintln(w, r.URL.Path)
}

func (slf Handler) halt(w http.ResponseWriter, r *http.Request, code int) {

	//codes := make(map[string]string)
	codes := map[string]string{
		//1 消息
		"100": "100 Continue",
		"101": "101 Switching Protocols",
		"102": "102 Processing",
		//2 成功
		"200": "200 OK",
		"201": "201 Created",
		"202": "202 Accepted",
		"203": "203 Non-Authoritative Information",
		"204": "204 No Content",
		"205": "205 Reset Content",
		"206": "206 Partial Content",
		"207": "207 Multi-Status",
		//3 重定向
		"300": "300 Multiple Choices",
		"301": "301 Moved Permanently",
		"302": "302 Move temporarily",
		"303": "303 See Other",
		"304": "304 Not Modified",
		"305": "305 Use Proxy",
		"306": "306 Switch Proxy",
		"307": "307 Temporary Redirect",
		//4 请求错误
		"400": "400 Bad Request",
		"401": "401 Unauthorized",
		"402": "402 Payment Required",
		"403": "403 Forbidden",
		"404": "404 Not Found",
		"405": "405 Method Not Allowed",
		"406": "406 Not Acceptable",
		"407": "407 Proxy Authentication Required",
		"408": "408 Request Timeout",
		"409": "409 Conflict",
		"410": "410 Gone",
		"411": "411 Length Required",
		"412": "412 Precondition Failed",
		"413": "413 Request Entity Too Large",
		"414": "414 Request-URI Too Long",
		"415": "415 Unsupported Media Type",
		"416": "416 Requested Range Not Satisfiable",
		"417": "417 Expectation Failed",
		"421": "421 too many connections",
		"422": "422 Unprocessable Entity",
		"423": "423 Locked",
		"424": "424 Failed Dependency",
		"425": "425 Unordered Collection",
		"426": "426 Upgrade Required",
		"449": "449 Retry With",
		"451": "451 Unavailable For Legal Reasons",
		//5 服务器错误（5、6字头）
		"500": "500 Internal Server Error",
		"501": "501 Not Implemented",
		"502": "502 Bad Gateway",
		"503": "503 Service Unavailable",
		"504": "504 Gateway Timeout",
		"505": "505 HTTP Version Not Supported",
		"506": "506 Variant Also Negotiates",
		"507": "507 Insufficient Storage",
		"509": "509 Bandwidth Limit Exceeded",
		"510": "510 Not Extended",
		"600": "600 Unparseable Response Headers",
	}

	code_status, ok := codes[string(code)]
	if !ok {
		code_status = codes["404"]
	}
	w.WriteHeader(code)
	fmt.Fprintln(w, "<h1>"+code_status+" </h1>")
}

func (slf Handler) local(bf BucketFile, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, bf.Name)
	fmt.Fprintln(w, bf)
}
func (slf Handler) remote(bf BucketFile, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, bf.Name)
	fmt.Fprintln(w, bf)
}
