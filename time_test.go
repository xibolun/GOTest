package basic

import (
	"fmt"
	"testing"
	"time"
)

// 日期处理
// 最近一周
// 最近半个月
// 最近一个月
// 最近半年
// 最近一年
// 是否闰年
// 一年共多少天
// 最后一天
// 当前月份最后一天
func Test_time(t *testing.T) {
	// current time
	fmt.Printf("current time : %s\n", time.Now())

	// format
	fmt.Printf("current time (ANSIC): %s\n", time.Now().Format(time.ANSIC))
	fmt.Printf("current time (Stamp): %s\n", time.Now().Format(time.Stamp))
	fmt.Printf("current time (RFC1123): %s\n", time.Now().Format(time.RFC1123))
	fmt.Printf("current time (UnixDate): %s\n", time.Now().Format(time.UnixDate))
	fmt.Printf("current time (Stamp): %s\n", time.Now().Format(time.Stamp))
	fmt.Printf("current time (UTC): %s\n", time.Now().UTC().String())
	fmt.Printf("current time (YYYY-MM-DD): %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("current time (YYYY-MM-DD HH:mm:ss): %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("current time (YYYY-MM-DD HH:mm): %s\n", time.Now().Format("2006-01-02 15:04"))
	fmt.Printf("current time (YYYYMMDDHHmmss): %s\n", time.Now().Format("20060102150405"))

	// year month day weekday
	fmt.Printf("year day: %d\n", time.Now().YearDay())
	fmt.Printf("day of year: %d\n", DayOfYear())
	fmt.Printf("day of half year: %d\n", DayOfHalfYear())
	fmt.Printf("current Year: %d\n", time.Now().Year())
	fmt.Printf("current Month: %s\n", time.Now().Month())
	fmt.Printf("current Month(int): %d\n", time.Now().Month())
	fmt.Printf("current Day: %d\n", time.Now().Day())
	fmt.Printf("current Weekday: %s\n", time.Now().Weekday())
	fmt.Printf("current Weekday(int): %d\n", time.Now().Weekday())

	// 日子过了多少天
	fmt.Printf("日子过了多少天: %d\n", DayFromNow(time.Date(2015, 5, 8, 0, 0, 0, 0, time.Now().Location())))
	fmt.Printf("日子过了多少天: %d\n", DaySinceTime(time.Date(2015, 5, 8, 0, 0, 0, 0, time.Now().Location())))

	// string date to time
	fmt.Printf("string date to time %s\n", StrToTime("2018-05-08", "2006-01-02"))
	fmt.Printf("isEqual: %t\n", time.Now().Equal(time.Now())) //time.now两次是不相等的
	fmt.Printf("isEqual: %t\n", StrToTime("2018-05-08", "2006-01-02").Equal(StrToTime("2018-05-08", "2006-01-02")))

	startTime := time.Now()
	fmt.Printf("startTime is: %s\n", startTime)
	//time.Sleep(10 * time.Second)
	// sub time
	fmt.Printf("sub time: %f\n", time.Now().Sub(startTime).Seconds())
	fmt.Printf("isAfter: %t\n", time.Now().After(startTime))
	fmt.Printf("time add: %s\n", startTime.AddDate(1, 1, 1).Format("2006-01-02"))
	fmt.Printf("isBefore: %t\n", time.Now().Before(startTime))
	fmt.Printf("since from startTime: %f\n", time.Since(startTime).Seconds())
	fmt.Printf("time truncate: %s\n", startTime.Truncate(2*time.Hour))
	fmt.Printf("day truncate: %s\n", startTime.Truncate(24*time.Hour))
	fmt.Printf("day minus: %s\n", startTime.Add(-24*time.Hour))

}

// 一年有多少天
// 先获取最后一天，再获取第一天，
func DayOfYear() int {
	year := time.Now().Year()
	lastDay := time.Date(year, 12, 31, 24, 59, 59, 0, time.Now().Location())
	firstDay := time.Date(year, 1, 1, 0, 0, 0, 0, time.Now().Location())
	return int(lastDay.Sub(firstDay).Hours() / 24)
}

// 一年的最后一天
func LastDayOfYear() time.Time {
	return time.Date(time.Now().Year(), 12, 31, 24, 59, 59, 0, time.Now().Location())
}

// 半年有多少天
func DayOfHalfYear() int {
	year := time.Now().Year()
	lastDay := time.Date(year, 6, 30, 24, 59, 59, 0, time.Now().Location())
	firstDay := time.Date(year, 1, 1, 0, 0, 0, 0, time.Now().Location())
	return int(lastDay.Sub(firstDay).Hours() / 24)
}

// 距离现在多少天
func DayFromNow(t time.Time) int {
	return int(time.Now().Sub(t).Hours() / 24)
}

// 距离现在多少天
func DaySinceTime(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// string转time
func StrToTime(str, layout string) time.Time {
	time2, _ := time.ParseInLocation(layout, str, time.Local)
	return time2
}

//TimestampToTime
func TimestampToTime(stamp int64) time.Time {
	return time.Unix(stamp, 0)
}

func TestStringToTime(t *testing.T) {
	time1, _ := time.ParseInLocation("2006-01-02", "2099-09-01", time.Local)
	time2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-12-15 17:54:31", time.Local)

	fmt.Println(time1)
	fmt.Println(time2.Format("20060102150405"))
}

func Test_LongToStr(t *testing.T) {

	t1 := TimestampToTime(1547696100).String()
	t2 := StrToTime("12/20/2018 15:35:50", "01/02/2006 15:04:05").String()

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t1 > t2)

	//time, err := time.Parse("2006-01-02 15:04:05", t1)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//fmt.Println(time.String())

}

func Test_Ticket(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)

	i := 0
	go func() {
		for { //循环
			<-ticker.C
			i++
			fmt.Println("i =", i)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()
}

func Test_Duration(t *testing.T) {
	t1 := time.Now()
	time.Sleep(1 * time.Second)

	t2 := time.Now()

	dur := t2.Sub(t1)

	/**
	1.000838681
	1000838681
	0.016680644683333332
	0.00027801074472222223
	*/

	fmt.Println(dur.Seconds())
	fmt.Println(dur.Nanoseconds())
	fmt.Println(dur.Minutes())
	fmt.Println(dur.Hours())
}

func Test_After(t *testing.T) {
	//t1 := time.After(5 * time.Microsecond)
	t2 := time.NewTicker(1 * time.Second)
	timeout := time.NewTicker(2 * time.Second)
	i := 0
	//go func() {
DONE:
	for {
		select {
		//case <-t1:
		//	fmt.Println("1s定时输出")
		case <-t2.C:
			fmt.Println("t2 1s输出")
			i++
			if i < 5 {
				continue
			}
			if i > 5 {
				t2.Stop()
				break DONE
			}
		case <-timeout.C:
			return
		}

	}

	fmt.Println("for done")
}
