package common

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func HttpRequest(url, method string, requestBody []byte, headers map[string]string) ([]byte, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp.Body)
}

// 分享内容中提取链接
func ExtractURL(text string) (string, error) {
	// Define the regular expression pattern for matching URLs
	urlPattern := `https?://[^\s]+`

	// Compile the regular expression
	re := regexp.MustCompile(urlPattern)

	// Find the first match in the text
	match := re.FindString(text)

	if match == "" {
		return "", fmt.Errorf("no url found in provided text")
	}
	return match, nil
}

// 获取重定向地址
func HttpGetRedirectURL(url string) (string, error) {
	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 获取重定向URL
	redirectURL := resp.Request.URL
	if redirectURL.String() == "" {
		return "", fmt.Errorf("empty redirect URL")
	}
	return redirectURL.String(), nil
}

// 获取url返回的html内容
func HttpGetUrlHTMLContent(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}
	return document.Html()
}

func ExtractHTML(htmlContent, htmlSelect string) (string, error) {
	// 使用正则表达式提取 ytInitialPlayerResponse 的值
	re := regexp.MustCompile(htmlSelect)
	matches := re.FindStringSubmatch(htmlContent)
	if len(matches) < 2 {
		return "", fmt.Errorf("html 正则匹配失败，请联系管理员: %v", len(matches))
	}

	// 清理并返回 JSON 字符串
	jsonString := strings.TrimSpace(matches[1])
	return jsonString, nil
}

func GetM3U8Key(s string) (string, error) {
	parts := strings.Split(s, ",")
	for _, part := range parts {
		if strings.HasPrefix(part, "URI=\"") {
			return part[5 : len(part)-1], nil
		}
	}
	return "", fmt.Errorf("提取m3u8 key失败")
}
