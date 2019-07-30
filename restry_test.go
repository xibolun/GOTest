package basic

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://10.0.30.21:8092/rbac/api/authInfo?appId=cloudboot&token=eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiLotoXnuqfnrqHnkIblkZgiLCJ1c2VySWQiOiI1OWRmNTk2MGNkNmFjMzVmNTMxMzViMzEiLCJuYW1lIjoi6LaF57qn566h55CG5ZGYIiwibG9naW5JZCI6ImFkbWluIiwibG9naW5OYW1lIjoiYWRtaW4iLCJ0ZW5hbnRJZCI6ImRlZmF1bHQiLCJ0aW1lb3V0IjoyMTYwMCwiZXhwIjoxNTQ3NzUyMjYzLCJjcmVhdFRpbWUiOjE1NDc3MzA2NjMyNTAsInRlbmFudE5hbWUiOiLnrqHnkIbnp5_miLcifQ.1L95aztIyAiz2Jvf-iAF00BhuBouj3xHduxweMFnzo0", nil)
	req.Header.Set("access-token", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiLotoXnuqfnrqHnkIblkZgiLCJ1c2VySWQiOiI1OWRmNTk2MGNkNmFjMzVmNTMxMzViMzEiLCJuYW1lIjoi6LaF57qn566h55CG5ZGYIiwibG9naW5JZCI6ImFkbWluIiwibG9naW5OYW1lIjoiYWRtaW4iLCJ0ZW5hbnRJZCI6ImRlZmF1bHQiLCJ0aW1lb3V0IjoyMTYwMCwiZXhwIjoxNTQ3NzUyMjYzLCJjcmVhdFRpbWUiOjE1NDc3MzA2NjMyNTAsInRlbmFudE5hbWUiOiLnrqHnkIbnp5_miLcifQ.1L95aztIyAiz2Jvf-iAF00BhuBouj3xHduxweMFnzo0")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)

}

//func TestRestryTest(t *testing.T) {
//	resp, err := resty.R().
//		SetHeader("access-token", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiLotoXnuqfnrqHnkIblkZgiLCJ1c2VySWQiOiI1OWRmNTk2MGNkNmFjMzVmNTMxMzViMzEiLCJuYW1lIjoi6LaF57qn566h55CG5ZGYIiwibG9naW5JZCI6ImFkbWluIiwibG9naW5OYW1lIjoiYWRtaW4iLCJ0ZW5hbnRJZCI6ImRlZmF1bHQiLCJ0aW1lb3V0IjoyMTYwMCwiZXhwIjoxNTQ3NzUyMjYzLCJjcmVhdFRpbWUiOjE1NDc3MzA2NjMyNTAsInRlbmFudE5hbWUiOiLnrqHnkIbnp5_miLcifQ.1L95aztIyAiz2Jvf-iAF00BhuBouj3xHduxweMFnzo0").
//		SetQueryParam("appId", "cloudboot").
//		SetQueryParam("token", "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiLotoXnuqfnrqHnkIblkZgiLCJ1c2VySWQiOiI1OWRmNTk2MGNkNmFjMzVmNTMxMzViMzEiLCJuYW1lIjoi6LaF57qn566h55CG5ZGYIiwibG9naW5JZCI6ImFkbWluIiwibG9naW5OYW1lIjoiYWRtaW4iLCJ0ZW5hbnRJZCI6ImRlZmF1bHQiLCJ0aW1lb3V0IjoyMTYwMCwiZXhwIjoxNTQ3NzUyMjYzLCJjcmVhdFRpbWUiOjE1NDc3MzA2NjMyNTAsInRlbmFudE5hbWUiOiLnrqHnkIbnp5_miLcifQ.1L95aztIyAiz2Jvf-iAF00BhuBouj3xHduxweMFnzo0").
//		Get(url)
//}
