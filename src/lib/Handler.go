package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"os"
	"strings"
)

type Handler struct {
	Conf Config
}

func (slf Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cdn-Source","gocdn")
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

	//分析版本
	reg = regexp.MustCompile(`(?U)^/v\.([^/]+)(/.*?)`)
	all = reg.FindAllStringSubmatch(ver_path, -1)
	ver := "default"
	min := false
	if len(all) > 0 {
		ver = all[0][1]
		reg = regexp.MustCompile(`\.min`)
		min = reg.MatchString(ver)
		filename = all[0][2]
	} else {

	}

	//获取文件ext
	reg = regexp.MustCompile(`(?U)\.([^.]+)$`)
	all = reg.FindAllStringSubmatch(filename, -1)
	ext := ""
	if len(all) > 0 {
		ext = all[0][1]
	} else {

	}

	//文件对象
	bf := BucketFile{
		Bucket{
			bucket.Name,
			bucket.IsLocal,
			bucket.Root,
			bucket.Deny,
		},
		filename,
		ver,
		min,
		ext,
		getMime(ext, slf.Conf),
	}
	if slf.checkDeny(bf) {
		slf.halt(w, r, 403)
		return
	}

	if bucket.IsLocal {
		slf.local(bf, w, r)
	} else {
		slf.remote(bf, w, r)
	}

	return
}

func (slf Handler) local(bf BucketFile, w http.ResponseWriter, r *http.Request) {
	filename := bf.Root + bf.Filename
	//fmt.Fprintln(w,filename)
	//判断文件是否存在
	stat,err := os.Stat(filename)
	if err != nil {
		slf.halt(w, r, 400)
		return
	}

	//判断目录
	if stat.IsDir() {
		slf.halt(w, r, 403)
		return
	}

	lenthg := stat.Size()
	w.WriteHeader(200)
	w.Header().Set("Content-type", bf.Mime)
	w.Header().Set("Content-length", strconv.FormatInt(lenthg, 10))

	//逐行读取文件
	file ,err := os.Open(filename)
	if err != nil {
		slf.halt(w, r, 403)
		return
	}
	defer file.Close()
	b := make([]byte, 1024)
	for {
		_,err := file.Read(b)
		if err != nil {
			break
		}
		w.Write(b)
	}


	//fmt.Fprintln(w,stat)
	//fmt.Fprintln(w,err)
}
func (slf Handler) remote(bf BucketFile, w http.ResponseWriter, r *http.Request) {
	//拼装url
	url := bf.Root + bf.Filename
	resp, err := http.Get(url)
	if err != nil {
		slf.halt(w, r, 400)
		return
	}

	if resp.StatusCode != 200 {
		slf.halt(w, r, resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	if err != nil {
		slf.halt(w, r, 403)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-type", bf.Mime)
	w.Header().Set("Content-length", strconv.FormatInt(resp.ContentLength, 10))
	w.Write(body)
}

func (slf Handler) checkDeny(bf BucketFile) bool  {
	dep := ","
	denys := dep+slf.Conf.Deny +dep+ bf.Deny+dep
	return strings.Contains(denys,dep+bf.Ext+dep)
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

	code_status, ok := codes[strconv.Itoa(code)]
	if !ok {
		code_status = codes["404"]
	}
	w.WriteHeader(code)
	fmt.Fprintln(w, "<h1>"+code_status+" </h1>")
}
