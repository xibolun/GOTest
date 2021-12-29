package ip

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"github.com/smartystreets/goconvey/convey"
	"net"
	"strings"
	"testing"
)

func Test_CalcMask(t *testing.T) {
	convey.Convey("cacl mask", t, func() {

		type Source struct {
			Mask   string
			Binary string
			CIDR   int
		}

		var source = []Source{
			{
				Mask:   "255.255.255.0",
				Binary: "11111111.11111111.11111111.00000000",
				CIDR:   24,
			},
			{
				Mask:   "FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FF00",
				Binary: "11111111.11111111.11111111.11111111.11111111.11111111.11111111.11111111.11111111.11111111.11111111.11111111.11111111.11111111.11111111.00000000",
				CIDR:   120,
			},
		}

		for _, item := range source {
			binStr := iplib.IPToBinaryString(net.ParseIP(item.Mask))

			convey.So(item.Binary, convey.ShouldEqual, binStr)

			cidr := 0
			for _, i := range strings.Replace(binStr, ".", "", -1) {
				if i != '1' {
					break
				}
				cidr += 1
			}
			convey.So(item.CIDR, convey.ShouldEqual, cidr)
		}

	})
}

func GetCIDR(ip, mask string) string {
	binStr := iplib.IPToBinaryString(net.ParseIP(mask))

	cidr := 0
	for _, i := range strings.Replace(binStr, ".", "", -1) {
		if i != '1' {
			break
		}
		cidr += 1
	}
	return fmt.Sprintf("%s/%d", ip, cidr)
}

func Test_IPRange(t *testing.T) {
	convey.Convey("test ip range ", t, func() {
		start := "61.155.168.80"
		mask := "255.255.255.248"

		cidr := GetCIDR(start, mask)

		ip, ne, _ := iplib.ParseCIDR(cidr)

		t.Log(ne.FirstAddress())
		t.Log(ne.LastAddress())


		iplib.NextIP(ip)

	})
}

func GetLastIP(start, mask string) string {
	_, ne, _ := iplib.ParseCIDR(GetCIDR(start, mask))
	return ne.LastAddress().String()
}

func Test_GetLastIP(t *testing.T) {
	//convey.Convey("test get last ip", t, func() {
	//	convey.So("10.22.10.2", convey.ShouldEqual, GetLastIP("10.22.10.0", "255.255.255.248"))
	//})

	t.Log(GetLastIP("61.155.168.80", "255.255.255.248"))
}

func Test_CIDR(t *testing.T) {
	convey.Convey("test cidr", t, func() {
		cidr := "10.22.10.0/29"

		ip, ne, err := iplib.ParseCIDR(cidr)
		convey.So(err, convey.ShouldBeNil)
		convey.So(ne, convey.ShouldNotBeNil)
		convey.So("10.22.10.0", convey.ShouldEqual, ip.String())

		mask, _ := ne.Mask().Size()
		//convey.So(int(4), convey.ShouldEqual, ne.Version)
		convey.So(29, convey.ShouldEqual, mask)

		// cidr ip range
		net4 := iplib.NewNet4(ip, mask)

		convey.So("10.22.10.1", convey.ShouldEqual, net4.FirstAddress().String())
		convey.So("10.22.10.6", convey.ShouldEqual, net4.LastAddress().String())
		convey.So("10.22.10.7", convey.ShouldEqual, net4.BroadcastAddress().String())
		convey.So("10.22.10.8", convey.ShouldEqual, iplib.NextIP(net4.BroadcastAddress()).String())

		// 31 bit
		net4 = iplib.Net4FromStr("10.22.10.0/31")
		convey.So("10.22.10.0", convey.ShouldEqual, net4.FirstAddress().String())
		convey.So("10.22.10.1", convey.ShouldEqual, net4.LastAddress().String())
		convey.So("10.22.10.1", convey.ShouldEqual, net4.BroadcastAddress().String())
		convey.So("10.22.10.2", convey.ShouldEqual, iplib.NextIP(net4.BroadcastAddress()).String())

		// 32 bit
		net4 = iplib.Net4FromStr("10.22.10.0/32")
		convey.So("10.22.10.0", convey.ShouldEqual, net4.FirstAddress().String())
		convey.So("10.22.10.0", convey.ShouldEqual, net4.LastAddress().String())
		convey.So("10.22.10.0", convey.ShouldEqual, net4.BroadcastAddress().String())
		convey.So("10.22.10.1", convey.ShouldEqual, iplib.NextIP(net4.BroadcastAddress()).String())
	})

}
