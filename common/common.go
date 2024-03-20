package common

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

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

var sendGroupMessageTime = make(map[int64]time.Time)

func SendQQMessage(cqUrl, message string, userid, groupid int64) ([]byte, error) {
	params := url.Values{}
	if userid != 0 {
		params.Add("user_id", fmt.Sprintf("%v", userid))
	} else if groupid != 0 {
		if !sendGroupMessageTime[groupid].IsZero() {
			if time.Since(sendGroupMessageTime[groupid]) < 5*time.Minute {
				return nil, fmt.Errorf("发送频率太快，忽略。")
			}
		}
		sendGroupMessageTime[groupid] = time.Now()
		params.Add("group_id", fmt.Sprintf("%v", groupid))
	}
	params.Set("auto_escape", "false") // 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 reply 字段是字符串时有效
	params.Add("message", message)
	return HttpRequest(
		fmt.Sprintf("%v/send_msg?%v", cqUrl, params.Encode()),
		http.MethodGet, nil, nil)
}

func Md5(input string) string {
	// 创建一个 md5 hash 对象
	hash := md5.New()

	// 将字符串转换为字节切片，并写入 hash 对象
	hash.Write([]byte(input))

	// 获取 hash 值的摘要（字节切片）
	hashSum := hash.Sum(nil)

	// 将摘要转换为十六进制字符串
	return fmt.Sprintf("%x", hashSum)
}

// 标准文件名格式
func StandardDownloadFileName(s string) string {
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "(", "-")
	s = strings.ReplaceAll(s, ")", "")

	for {
		s = strings.TrimPrefix(s, "-")
		if s == strings.TrimPrefix(s, "-") {
			break
		}
	}
	for {
		s = strings.TrimSuffix(s, "-")
		if s == strings.TrimSuffix(s, "-") {
			break
		}
	}
	return s
}
