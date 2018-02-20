package lib

type Bucket struct {
	Name    string
	IsLocal bool
	Root    string
}

type BucketFile struct {
	Bucket
	Filename string
	Ver      string
	IsMin    bool
	Ext string
	Mime string
}

func getMime(ext string,config Config) string {
	mime,ok := config.Mimes[ext]
	if !ok {
		mime = "application/octet-stream"
	}
	return mime
}
