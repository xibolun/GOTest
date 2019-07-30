package dell_irdac

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	main()
}

var client *http.Client

func main() {
	client = &http.Client{}

	postLogin()
}

func postLogin() {
	req, err := http.NewRequest(http.MethodPost, "https://10.0.10.100/data/login", strings.NewReader("user=root&password=calvin"))
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err!=nil{
		log.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	resp.Header.Set("Access-Control-Allow-Origin", "*")

	fmt.Println(body)
}

func fireUrl() {

}
