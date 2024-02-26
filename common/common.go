package common

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

func MD5(s string) string {
	md5Hash := md5.New()
	// 将文本写入哈希对象
	md5Hash.Write([]byte(s))
	// 获取哈希值的十六进制表示
	hashSum := md5Hash.Sum(nil)
	return fmt.Sprintf("%x", hashSum)
}

// 传入单位为秒
func GetTimeHourMinuteSecoud(t int) string {
	if t < 60 {
		return fmt.Sprintf("%v秒", t)
	}
	if t >= 60 && t < 3600 {
		return fmt.Sprintf("%v分%v秒", t/60, t%60)
	}
	return fmt.Sprintf("%v时%v分%v秒", t/3600, t%3600/60, t%3600%60)
}

// 获取一段文本信息中所有的链接
func GetLinks(text string) []string {
	// 匹配http和https开头的链接
	re := regexp.MustCompile(`(https?://\S+)"`)
	return re.FindAllString(text, -1)
}
