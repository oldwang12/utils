package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DingTalk struct {
	Secret      string
	AccessToken string
}

type DingTalkReqponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

const DingTalkPrefix = "https://oapi.dingtalk.com/robot/send"

func NewDingTalk(accessToken, secret string) *DingTalk {
	return &DingTalk{
		Secret:      secret,
		AccessToken: accessToken,
	}
}

// 生成钉钉Webhook签名
func (d *DingTalk) sign(timestamp int64) string {
	strToHash := fmt.Sprintf("%d\n%s", timestamp, d.Secret)
	hmac256 := hmac.New(sha256.New, []byte(d.Secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}

func (d *DingTalk) SendText(content string) error {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	sign := d.sign(timestamp)
	url := fmt.Sprintf("%v?access_token=%v&timestamp=%v&sign=%v", DingTalkPrefix, d.AccessToken, timestamp, sign)

	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": content,
		},
	}
	requestBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	return d.Request(url, http.MethodPost, requestBody, headers)
}

func (d *DingTalk) SendMarkDown(title, text string) error {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	sign := d.sign(timestamp)
	url := fmt.Sprintf("%v?access_token=%v&timestamp=%v&sign=%v", DingTalkPrefix, d.AccessToken, timestamp, sign)

	payload := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  text,
		},
	}
	requestBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	return d.Request(url, http.MethodPost, requestBody, headers)
}

func (d *DingTalk) Request(url, method string, requestBody []byte, headers map[string]string) error {
	b, err := HttpRequest(url, http.MethodPost, requestBody, headers, nil)
	if err != nil {
		return err
	}

	var res DingTalkReqponse
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}

	if res.ErrCode != 0 || res.ErrMsg != "ok" {
		return fmt.Errorf("failed to send message: %v", res)
	}
	return nil
}
