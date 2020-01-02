package web

import (
	"me.org/http"
	"errors"
	"math/rand"
	"fmt"
	url2 "net/url"
	json "github.com/json-iterator/go"
)

/*
 * GetCaptcha
 *
 * 获取验证码

 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return []byte	图片字节流
 * @return []*http.Cookie
 * @return error
 *
 */
func GetCaptcha(client *http.Client, url string, cookies []*http.Cookie) ([]byte, []*http.Cookie, error) {
	header := http.Header{
		"Accept":          {"image/webp,image/apng,image/*,*/*;q=0.8"},
		"Accept-Encoding": {"gzip, deflate, br"},
		"Accept-Language": {"zh-CN, zh; q = 0.9"},
		"Cache-Control":   {"no-cache"},
		"Connection":      {"keep-alive"},
		"Host":            {"kyfw.12306.cn"},
		"Pragma":          {"no-cache"},
		"Referer":         {"https://kyfw.12306.cn/otn/login/init"},
	}
	if client == nil {
		return nil, nil, ClientIsNilErr
	}
	client.SetHeader(header)
	client.AddCookies(cookies)
	sjrand := rand.Float64()
	url = fmt.Sprintf("%s?login_site=E&module=login&rand=sjrand&%0.16f", url, sjrand)
	resp, err := client.AsyncGet(url, nil)
	if err != nil {
		return nil, nil, err
	}
	_cookies := resp.Cookies()
	client.AddCookies(_cookies)
	return resp.Body, _cookies, nil
}

/*
 * CheckCaptcha
 *
 * 提交验证码

 * @param client 		*http.Client	http客户端
 * @param url 			string			url
 * @param cookies 		[]*http.Cookie	cookies
 * @param answer 		string			验证码坐标
 *
 * @return *Result
 * @return []*http.Cookie
 * @return bool			true - 验证成功 | false - 验证失败
 * @return error
 *
 */
func CheckCaptcha(client *http.Client, url string, cookies []*http.Cookie, answer string) (*Result, []*http.Cookie, bool, error) {
	headers := http.Header{
		"Accept":           {"application/json, text/javascript, */*; q=0.01"},
		"Accept-Encoding":  {"deflate, br"},
		"Accept-Language":  {"zh-CN, zh; q = 0.9"},
		"Cache-Control":    {"no-cache"},
		"Connection":       {"keep-alive"},
		"Content-Type":     {"application/x-www-form-urlencoded; charset=UTF-8"},
		"Host":             {"kyfw.12306.cn"},
		"Origin":           {"https://kyfw.12306.cn"},
		"Pragma":           {"no-cache"},
		"Referer":          {"https://kyfw.12306.cn/otn/login/init"},
		"X-Requested-With": {"XMLHttpRequest"},
	}

	if client == nil {
		return nil, nil, false, ClientIsNilErr
	}

	client.SetHeader(headers)
	client.AddCookies(cookies)

	data := make(url2.Values)

	data["answer"] = []string{answer}
	data["login_site"] = []string{"E"}
	data["rand"] = []string{"sjrand"}

	resp, err := client.AsyncPostForm(url, data)
	if err != nil {
		return nil, nil, false, err
	}

	_cookies := resp.Cookies()
	client.AddCookies(_cookies)

	if resp.StatusCode != 200 {
		s := fmt.Sprintf("http status code is %d", resp.StatusCode)
		return nil, _cookies, false, errors.New(s)
	}

	result := &Result{
		ResultMessage: "新对象",
	}

	err = json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil,_cookies, false, err
	}

	succeed := false
	if result.ResultCode == 4 {
		succeed = true
	}

	return result,_cookies, succeed, nil
}
