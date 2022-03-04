package elastic

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_common(t *testing.T) {
	convey.Convey("test es common", t, func() {

		err := InitEs()
		convey.So(err, convey.ShouldBeNil)

		convey.Println(elasticsearch.Version)

		convey.So(elasticsearch.Version, convey.ShouldNotBeZeroValue)

	})
}
