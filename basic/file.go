package basic

import (
	"fmt"
	"io/ioutil"
	"strings"

	"xibolun/gotest/util"
)

func Start(dir string) {
	files := util.GetFileList(dir, ".DS_Store")

	for _, item := range files {
		content, err := ioutil.ReadFile(item.Name())
		if err != nil {
			fmt.Printf("read file %s err,%s", item.Name(), err.Error())
			continue
		}

		str := string(content)

		str = strings.Replace(str, "+++", "---", -1)
		str = strings.Replace(str, "date =", "\ndate : ", -1)
		str = strings.Replace(str, "title = ", "\ntitle : ", -1)
		str = strings.Replace(str, "categories = ", "\ncategories : ", -1)
		str = strings.Replace(str, "tags = ", "\ntags : ", -1)
		str = strings.Replace(str, "toc = ", "\ntoc : ", -1)

		fname := strings.Replace(item.Name(), "/Users/admin/projects/typora/content", "/tmp/content", -1)

		if err := util.CreateFileRecursive(fname); err != nil {
			fmt.Printf(err.Error())
			continue
		}
		if err := ioutil.WriteFile(fname, []byte(str), 0644); err != nil {
			fmt.Printf("fail write file %s,err: %s\n", fname, err.Error())
			continue
		}

	}
}

//
//func GetFileList(path string, files []*os.File) {
//	fis, err := ioutil.ReadDir(path)
//	if err != nil {
//		return
//	}
//	for _, f := range fis {
//		fmt.Println(f.Name())
//	}
//}
