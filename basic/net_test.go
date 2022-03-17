package basic

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_NetAPI(t *testing.T) {
	Convey("test net api", t, func() {

	})
}

type Stu struct {
	Age int64
}

func Test_newPoint(t *testing.T) {
	s := new(Stu)

	t.Error(s == nil)
}
