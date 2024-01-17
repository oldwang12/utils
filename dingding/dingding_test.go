package dingding

import (
	"testing"
)

// 生成钉钉Webhook签名
const accessToken = "" 
const secret = ""

func TestGenerateWebhookSignature(t *testing.T) {
	d := NewDingTalk(accessToken, secret)
	if err := d.SendText("123456"); err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestMarkdown(t *testing.T) {
	d := NewDingTalk(accessToken, secret)
	if err := d.SendMarkDown("新主题","# 你好 \n ![RUNOOB 图标](http://static.runoob.com/images/runoob-logo.png)"); err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
