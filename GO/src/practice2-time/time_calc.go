/*
 * @Author: Galen Tong
 * @Date: 2022-04-20 18:41:35
 * @LastEditTime: 2022-04-20 18:43:39
 * @Description: 计算函数执行时间
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// func sleep time : 3.002109652s
	start := time.Now()
	time.Sleep(3 * time.Second)
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("func sleep time : %s\n", duration)
}
