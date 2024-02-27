package common

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"

	"mvdan.cc/xurls/v2"

	"github.com/avast/retry-go"
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

// s: https://www.baidu.com/query
// seq: /
// 返回: query
func GetLastIndexValue(s, sep string) string {
	strs := strings.Split(s, sep)
	return strs[len(strs)-1]
}

// s: https://www.baidu.com/query
// seq: /
// 返回: https://www.baidu.com/
func GetBeforeLastIndexValue(s, sep string) string {
	lastIndexValue := GetLastIndexValue(s, sep)
	return strings.TrimRight(s, lastIndexValue)
}

// 多次执行函数
func RetryFunc(fn func() error, retryAttempts int, waitTime time.Duration) error {
	return retry.Do(
		fn,
		retry.Delay(waitTime),
		retry.Attempts(uint(retryAttempts)),
		retry.DelayType(retry.FixedDelay),
	)
}
