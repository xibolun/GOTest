package basic

import (
	"fmt"
	"strings"
	"testing"
)

type Varyse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Varyse) validate() error {
	var errstr []string
	if p.Age <= 0 {
		errstr = append(errstr, fmt.Sprintf("age must big than 0"))
	}

	if p.Name == "" {
		errstr = append(errstr, fmt.Sprintf("person name can not be null"))
	}

	return fmt.Errorf(strings.Join(errstr, "\n"))

}

func TestValidte(t *testing.T) {
	p := &Varyse{
		Name: "",
		Age:  0,
	}

	t.Log(p.validate())
}
