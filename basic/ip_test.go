package basic

import (
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"testing"
)

func Test_isPrivate(t *testing.T) {
	Convey("test ip functions", t, func() {
		So(net.ParseIP("172.60.239.61").IsPrivate(), ShouldEqual, false)
		So(net.ParseIP("172.60.116.24").IsPrivate(), ShouldEqual, false)
		So(net.ParseIP("172.60.239.61").IsPrivate(), ShouldEqual, false)
		So(net.ParseIP("172.60.239.61").IsPrivate(), ShouldEqual, false)
		So(net.ParseIP("fe80::6e92:bfff:fe04:b309").IsPrivate(), ShouldEqual, false)
		So(net.ParseIP("fe80::569f:35ff:fe0f:6340").IsPrivate(), ShouldEqual, false)
	})

}
