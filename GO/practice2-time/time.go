/*
 * @Author: Galen Tong
 * @Date: 2022-04-17 15:38:51
 * @LastEditTime: 2022-04-17 15:51:13
 * @Description:
 */
package main

import (
	"fmt"
	"time"
)

var week time.Duration // 两个时刻相差的纳秒数, int64

func main() {
	t := time.Now()
	fmt.Println(t)                                              // 2022-04-17 15:46:13.463958217 +0800 CST m=+0.000043685
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) // 17.04.2022

	t = time.Now().UTC()
	fmt.Println(t)          // 2022-04-17 07:46:13.464108109 +0000 UTC
	fmt.Println(time.Now()) // 2022-04-17 15:46:13.464112616 +0800 CST m=+0.000198082

	week = 60 * 60 * 24 * 7 * 1e9 // need nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now) // 2022-04-24 07:46:13.464108109 +0000 UTC

	fmt.Println(t.Format(time.RFC822))         // 17 Apr 22 07:46 UTC
	fmt.Println(t.Format(time.ANSIC))          // Sun Apr 17 07:46:13 2022
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 17 Apr 2022 07:46

	s := t.Format("20060102")
	fmt.Println(t, "=>", s) // 2022-04-17 07:46:13.464108109 +0000 UTC => 20220417
}
