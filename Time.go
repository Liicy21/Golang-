package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 获取当前时间
	now := time.Now()
	fmt.Println("当前完整时间：", now)
	fmt.Printf("类型：%T\n", now)
	// time.Now() 返回 time.Time 类型，不是字符串，里面包含年、月、日、时分秒、时区

	// 2. 单独取出 年、月、日、时、分、秒
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("年:%d 月:%d 日:%d 时:%d 分:%d 秒:%d\n", year, month, day, hour, minute, second)
	// month 是 time.Month 类型，打印出来默认是英文（October），转数字直接用int(month)

	// 3. 时间戳（Unix时间：从1970-01-01 00:00:00 UTC到现在的秒数）
	timestamp := now.Unix()        // 秒级时间戳
	timestampMs := now.UnixMilli() // 毫秒时间戳
	fmt.Println("秒级时间戳：", timestamp)
	fmt.Println("毫秒时间戳：", timestampMs)

	// 4. 时间戳转回 time.Time
	t := time.Unix(timestamp, 0)
	fmt.Println("时间戳转回时间：", t)

	// 5. 时间格式化！重点：Go特殊，必须用这个参考时间 "2006-01-02 15:04:05"
	// 记住口诀：06-01-02 15:04:05  （1月2号下午3点4分5秒06年）
	str := now.Format("2006-01-02 15:04:05")
	fmt.Println("格式化日期字符串：", str)
	str2 := now.Format("2006/01/02")
	fmt.Println("只保留年月日：", str2)

	// 6. 时间间隔 Duration 单位：time.Second time.Minute time.Hour
	d := 2 * time.Second
	fmt.Println("2秒：", d)

	// 7. 时间加减 Add
	later := now.Add(2 * time.Hour) // 当前时间 + 2小时
	fmt.Println("2小时之后：", later.Format("2006-01-02 15:04:05"))
	before := now.Add(-1 * time.Hour) // 减1小时
	fmt.Println("1小时之前：", before.Format("2006-01-02 15:04:05"))

	// 8. 两个时间比较：Before After Equal
	t1 := now.Add(1 * time.Hour)
	fmt.Println("t1 在now之后？", t1.After(now))
	fmt.Println("t1 在now之前？", t1.Before(now))

	// 9. 时间睡眠 Sleep（让程序卡住暂停）
	fmt.Println("准备睡1秒...")
	time.Sleep(1 * time.Second)
	fmt.Println("睡醒了！")

	// 10. 字符串解析成时间 Parse
	timeStr := "2026-10-24 10:30:00"
	parseTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("解析失败：", err)
	} else {
		fmt.Println("字符串转time：", parseTime)
	}
}
