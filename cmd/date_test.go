/*
* @Author: happyyi
* @Date:   2017-10-20 16:50:41
* @Last Modified by:   happyyi
* @Last Modified time: 2017-10-20 19:39:10
*/
package cmd

import (
	// "net"
	// "io"
	"log"
	"fmt"
	"testing"
	"time"
	// "datetime"
)
func Reverse(s string) string {
    n := len(s)
    runes := make([]rune, n)
    for _, rune := range s {
        n--
        runes[n] = rune
    }
    return string(runes[n:])
}

func GetXorCode(date string) string{
	bDate := []byte(date)
	x := byte(0) 
	for _ , v := range bDate {
  		x ^= v
	}
	xor1 := fmt.Sprintf("%02x",x)
	return xor1
}

func TestTime(t *testing.T) {
	p := log.Println
	for i := 0; i < 100; i++ {
		now := time.Now()
		now = now.AddDate(0,0,i)
		s := fmt.Sprintf("%d%02d%02d",now.Year(),now.Month(),now.Day())
		p(s)
		xor1 := GetXorCode(s)
		xor2 := GetXorCode(s + xor1)
		p(s + xor1 + xor2)	
	}
	// GetXorCode("20170101")
	// GetXorCode("20170102")

	// p(now.Format("%Y%m%d"))
	// p(t.Format(time.RFC3339))
	// log.Printf("%s",start)
	// t = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
	// log.Printf(t)
}