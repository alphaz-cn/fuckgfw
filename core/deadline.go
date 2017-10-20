/*
* @Author: happyyi
* @Date:   2017-10-20 20:44:32
* @Last Modified by:   happyyi
* @Last Modified time: 2017-10-20 21:31:05
*/
package core

import (
	"errors"
	// "log"
	"fmt"
	"time"
	"github.com/araddon/dateparse"
)

func GetXorCode(date string) string{
	bDate := []byte(date)
	x := byte(0) 
	for _ , v := range bDate {
  		x ^= v
	}
	xor1 := fmt.Sprintf("%02x",x)
	return xor1
}

func CheckDeadLine(VerfyCode string) (time.Time, error){
	date := VerfyCode[0:8]
	xorcode := VerfyCode[len(VerfyCode)-2:]
	t, err := dateparse.ParseLocal(date)
	if err != nil {
		return t, errors.New("ERROR: 日期解析错误")
	}
	if xorcode == GetXorCode(date) {
		//log.Printf("OK, xorcode %s %s", xorcode, t)
		return t, nil
	} else {
		// log.Printf("ERROR, xorcode %s xor=%s t=%s", GetXorCode(date) ,xorcode, t)
		return t, errors.New("ERROR: 配置日期校验错误,Xor=" + GetXorCode(date))
	}
	return t,errors.New("日期校验未知错误")
}