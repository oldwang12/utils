package dingding

import (
	"testing"
)

// 生成钉钉Webhook签名
const accessToken = ""
const secret = ""

func TestGenerateWebhookSignature(t *testing.T) {
	d := NewDingTalk(accessToken, secret)
	if err := d.Send(Text, "123456"); err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
