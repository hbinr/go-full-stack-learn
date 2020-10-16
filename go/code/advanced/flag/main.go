package main

import (
	"flag"
	"fmt"
	"time"
)

// FlagDemo flag库使用demo
func FlagDemo() {
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 38, "年龄")
	married := flag.Bool("married", true, "婚否")
	time := flag.Duration("time", 0, "时间间隔")
	// 解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, time)
}

// FlagDemo2 flag库使用demo2
func FlagDemo2() {
	var (
		name    string
		age     int
		married bool
		delay   time.Duration
	)
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 28, "年龄")
	flag.BoolVar(&married, "married", true, "婚否")
	flag.DurationVar(&delay, "delay", 0, "时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)

	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

}
func main() {
	// FlagDemo()
	FlagDemo2()
}
