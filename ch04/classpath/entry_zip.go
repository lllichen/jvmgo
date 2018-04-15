package classpath

import (
	"archive/zip"
	"io/ioutil"
	"errors"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry{
	absPath , err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}


func (ze *ZipEntry) readClass(className string) ([]byte,Entry,error){
	//fmt.Println(ze.String())
	//if strings.Contains(ze.String(),"rt.jar") {
	//	fmt.Println(ze.String())
	//}
	r,err := zip.OpenReader(ze.absPath)
	if err != nil{
		return nil,nil,err
	}
	defer r.Close()
	for _,f:=range r.File {
		//if strings.Contains(f.Name,"Object") &&strings.Contains(f.Name,"lang") {
		//	fmt.Println(f.Name)
		//}
		if f.Name == className {
			rc,err := f.Open()
			if err != nil{
				return nil, nil, err
			}
			defer rc.Close()
			data,err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, ze ,nil
		}
	}
return nil,nil,errors.New("class not found: "+ className)
}

func (ze *ZipEntry) String()string {
	return ze.absPath
}