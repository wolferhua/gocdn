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
}

func mime(ext string)  {

}
