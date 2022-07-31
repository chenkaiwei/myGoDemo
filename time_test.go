package myDemo_test

import (
	"fmt"
	"testing"
	"time"
)

func Test_Time(t *testing.T) {

	// 2006-01-02 15:04:05 -0700  定义模板用的日期是固定死的，不要自行更改，就是1234567，6提前到年上

	date, err2 := time.ParseInLocation("20060102", "20220518", time.Local)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println("date", date)
	fmt.Println("String ", date.String())

	//转换回字符串
	fmt.Println("Format ", date.Format("2006-01-02"))
	fmt.Println("Format ", date.Format("2006年1月2日"))

}
