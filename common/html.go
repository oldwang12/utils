package common

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetHTMLTitle(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}
	title := doc.Find("title").Text()
	return title, nil
}

// 获取url返回的html内容
func GetHTMLContentByUrl(url string) (string, error) {
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

// parseJson, err := common.ExtractHTML(lls.HtmlResponse, `window.detailData\s*=\s*({.*?});`)
// 用于提取部分内容（json内容）
func ExtractHTML(htmlContent, htmlSelect string) (string, error) {
	// 使用正则表达式提取 htmlSelect 的值
	re := regexp.MustCompile(htmlSelect)
	matches := re.FindStringSubmatch(htmlContent)
	if len(matches) < 2 {
		return "", fmt.Errorf("html 正则匹配失败： %v", len(matches))
	}
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

func GetHtmlTitle(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	return doc.Find("title").Text(), err
}

// 获取html中script的变量

// 这是一个正则表达式，用来匹配包含"key = xx;"的字符串。
func GetHtmlScriptVar(html, key string) (string, error) {
	return ExtractHTML(html, fmt.Sprintf(`%v\s*=\s*(.*?);`, key))
}

// 这是一个正则表达式，用来匹配包含"key = {};"的字符串。
func GetHtmlScriptJsonVar(html, key string) (string, error) {
	return ExtractHTML(html, fmt.Sprintf(`%v\s*=\s*({.*?});`, key))
}
