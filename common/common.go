package common

import (
	"crypto/md5"
	"fmt"
	"strings"

	"mvdan.cc/xurls/v2"
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
	return xurls.Strict().FindAllString(text, -1)
}

func GetLastIndexValue(s, sep string) string {
	strs := strings.Split(s, sep)
	return strs[len(strs)-1]
}
