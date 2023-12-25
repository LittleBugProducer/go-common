package main

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon/v2"
)

func main() {
	// lang := carbon.NewLanguage()
	// lang.SetLocale("zh-CN")
	// c := carbon.SetLanguage(lang)
	// fmt.Println(carbon.Now().ToDateTimeString())
	// fmt.Println(c.Tomorrow().ToMonthString())
	// fmt.Println(c.Now().Season())
	// fmt.Println("Hello, World!")

	fmt.Println(time.Now().Format(carbon.RFC3339MilliLayout))
	fmt.Println(time.Now().Format(carbon.RFC3339Layout))

	tm, err := time.Parse("2006-01-02T15:04:05.999Z07:00", "2023-12-20T14:10:47+08:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time.Now())
	fmt.Println(tm)
	fmt.Println(tm.Local().Format("2006-01-02T15:04:05.000Z07:00"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05.000Z07:00"))
	fmt.Println(tm.Local().Format(time.StampMilli))
}
