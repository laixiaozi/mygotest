package ryanutility

import "time"

//GetTodayString  返回当前日期时间的字符串格式
func GetTodayString() string {
	return time.Now().In(time.FixedZone("Asia/shanghai", 3600*8)).Format("2006-01-02 15:04:05")
}
