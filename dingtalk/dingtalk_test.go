package dingtalk

import (
	"testing"
)

// 生成钉钉Webhook签名
const accessToken = "89c9387eb2c329b534f7fceb0978aba05f532d886e09c47c8e355bc79ae62ea0" 
const secret = "SEC19016b1729cc55c41ea2e5a8de53be4bec446de59f5294f5429865664f321b5e"

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
