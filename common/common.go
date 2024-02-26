package common

import (
	"crypto/md5"
	"fmt"
)

func MD5(s string) string {
	md5Hash := md5.New()
	// 将文本写入哈希对象
	md5Hash.Write([]byte(s))
	// 获取哈希值的十六进制表示
	hashSum := md5Hash.Sum(nil)
	return fmt.Sprintf("%x", hashSum)
}
