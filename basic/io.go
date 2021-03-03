package basic

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var path = "/tmp/Swagger UI.html"

func ReadFile() {

	content, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(content)
}

func OpenFile() {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(file.Name())

	bufio.NewReader(file)
}

func ReaderFile() {

}
