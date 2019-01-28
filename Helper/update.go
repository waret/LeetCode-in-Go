package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mozillazg/request"
)

/**
 * 获取最新的答题记录
 * 需要携带 cookie 访问 leetcode.com 的 API
 */

const (
	loginPageURL = "https://leetcode.com/accounts/login/"
)

func getRaw(URL string) []byte {
	log.Printf("开始下载 %s 的数据", URL)

	req := newReq()
	resp, err := req.Get(URL)
	if err != nil {
		log.Fatal("getRaw: Get Error: " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("getRaw: Read Error: " + err.Error())
	}
	return body
}

var req *request.Request

func newReq() *request.Request {
	if req == nil {
		req = signin()
	}
	return req
}

// 登录 leetcode
// 返回的 req 带有 cookie
func signin() *request.Request {
	log.Println("正在登录中...")
	cfg := getConfig()

	// 对 req 赋值
	req := request.NewRequest(new(http.Client))

	// 配置request
	req.Headers = map[string]string{
		"Accept-Encoding": "",
		"Referer":         "https://leetcode.com/",
	}

	// login
	csrfToken := getCSRFToken(req)

	log.Printf("csrfToken: %s", csrfToken)

	req.Data = map[string]string{
		"csrfmiddlewaretoken": csrfToken,
		"login":               cfg.Username,
		"password":            cfg.Password,
	}

	resp, err := req.Post(loginPageURL)
	resp.Body.Close() // **Don't forget close the response body**
	check(err)

	log.Println("成功登录")

	return req
}

func getCSRFToken(req *request.Request) string {
	resp, err := req.Get(loginPageURL)
	if err != nil {
		log.Panicf("无法 Get 到 %s: %s", loginPageURL, err)
	}

	cookies := resp.Cookies()

	for _, ck := range cookies {
		if ck.Name == "csrftoken" {
			return ck.Value
		}
	}

	panic("无法在 Cookies 中找到 csrftoken")
}
