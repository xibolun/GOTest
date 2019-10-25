package basic

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)



func TestFileCnvert(t *testing.T) {
	Start("/Users/admin/projects/typora/content")
}

func TestPath(t *testing.T) {
	rootPath, _ := os.Executable()
	fmt.Println(rootPath)

}

func TestFile(t *testing.T) {
	dir := "/tmp/hello"
	path := "/tmp/hello.txt"
	// create dir
	os.Mkdir(dir, 0755)

	// create file
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	// filePath
	fmt.Printf("File.name(file full path): %v \n", file.Name())
	fmt.Println(filepath.Dir(file.Name()))
	fmt.Println(filepath.Abs(filepath.Dir(os.Args[0])))

	// current pwd
	pwd, err := os.Getwd()
	fmt.Printf("current pwd: %s\n", pwd)

	pwd, err = filepath.Abs("./")
	fmt.Printf("current pwd: %s\n", pwd)

	// exec path
	lookpath, _ := exec.LookPath(os.Args[0])
	fmt.Printf("exec path: %s\n", lookpath)

	// get file info
	fileInfo, err := file.Stat()
	fmt.Printf("%s\n", fileInfo.Name())

	// is exist
	if _, err := os.Stat("/tmp/bb.txt"); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("file is not exit\n")
		}
	}
	// is not exist		os.IsNotExist(err)
	// is permission  	os.IsPermission(err)

	// ll命令
	if err = filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("%t %s %d %s %s\n", info.IsDir(), info.Mode(), info.Size(), info.ModTime().Format("2006-01-02 15:04:05"), info.Name())
		return nil
	}); err != nil {
		t.Error(err)
	}

	// relative path

	// mv dir

	// mv file

	// rename dir
	os.Rename(dir, "boot")

	// rename file
	os.Rename(file.Name(), "hello.txt")

	// write file
	err = ioutil.WriteFile(file.Name(), []byte("hello world"), 0644)
	if err != nil {
		t.Error(err)
	}

	// read file
	bytes, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("file content: %s\n", string(bytes))

	// del dir
	os.RemoveAll(dir)

	// del file
	os.Remove(path)
}

func TestA(t *testing.T) {
	for i := 1; i < 20; i++ {
		fmt.Printf("%d,\n", i)
	}
}

func TestReadRoutes(t *testing.T) {

	type UriInfo struct {
		Method   string `json:"method"`
		URI      string `json:"uri"`
		Category string `json:"category"`
	}

	bytes, _ := ioutil.ReadFile("/Users/admin/projects/go/src/idcos.com/cloudboot/src/idcos.com/cloudboot/server/cloudbootserver/route2.go")

	s := string(bytes)

	var uris []UriInfo

	for _, row := range strings.Split(s, "\n") {
		var uriInfo UriInfo

		if !strings.Contains(row, "Put") && !strings.Contains(row, "Delete") && !strings.Contains(row, "Post") {
			continue
		}

		for i, str := range strings.Split(row, "\"") {
			if i == 2 {
				continue
			}

			if i == 0 {
				str = strings.Replace(str, "	mux.", "", -1)
				str = strings.Replace(str, "(", "", -1)
				uriInfo.Method = str
			}

			if i == 1 {
				uriInfo.URI = str
			}
		}

		operate := "rm_"

		switch uriInfo.Method {
		case "Put":
			operate = "mod_"
		case "Delete":
			operate = "rm_"
		case "Post":
			operate = "add_"
		default:
			operate = "null_"
		}

		category := strings.Replace(uriInfo.URI, "/api/cloudboot/v1/", "", -1)
		str := strings.Split(category, "/")[0]
		uriInfo.Category = operate + str

		uris = append(uris, uriInfo)
	}

	fmt.Println(ToJsonString(uris))

}
