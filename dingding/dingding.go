package dingding

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
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

const Text = "text"

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

func (d *DingTalk) Send(messageType, content string) error {
	if messageType == "" {
		messageType = Text
	}

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	sign := d.sign(timestamp)
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%v&timestamp=%v&sign=%v",
		d.AccessToken, timestamp, sign)

	payload := map[string]interface{}{
		"msgtype": messageType,
		"text": map[string]string{
			"content": content,
		},
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
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
