package web

import (
	"me.org/http"
	url2 "net/url"
	"fmt"
	"errors"
	json "github.com/json-iterator/go"
)

/*
 * Init
 *
 * 获取cookies
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 *
 * @return []*http.Cookie
 * @return error
 *
 */
func Init(client *http.Client, url string) ([]*http.Cookie, error) {
	header := http.Header{
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"},
		"Accept-Encoding":           {"gzip, deflate, br"},
		"Accept-Language":           {"zh-CN, zh; q = 0.9"},
		"Cache-Control":             {"max-age = 0"},
		"Connection":                {"keep-alive"},
		"Host":                      {"kyfw.12306.cn"},
		"Upgrade-Insecure-Requests": {"1"},
	}
	if client == nil {
		return nil, ClientIsNilErr
	}
	client.SetHeader(header)
	resp, err := client.AsyncGet(url, nil)
	if err != nil {
		return nil, err
	}
	cookies := resp.Cookies()
	client.AddCookies(cookies)
	return cookies, nil
}

/*
 * uamtk
 *
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param header 	http.Header
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return *Result
 * @return []*http.Cookie
 * @return error
 *
 */
func uamtk(client *http.Client, url string, header http.Header, cookies []*http.Cookie) (*Result, []*http.Cookie, error) {

	if client == nil {
		return nil, nil, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)

	form := make(url2.Values)
	form["appid"] = []string{"otn"}

	resp, err := client.AsyncPostForm(url, form)
	if err != nil {
		return nil, nil, err
	}

	_cookies := resp.Cookies()
	client.AddCookies(_cookies)

	if resp.StatusCode != 200 {
		s := fmt.Sprintf("http status code is %d", resp.StatusCode)
		return nil, nil, errors.New(s)
	}

	result := &Result{
		ResultMessage: "新对象",
	}

	err = json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, nil, err
	}

	return result, _cookies, nil
}

/*
 * Uamtk
 *
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return *Result
 * @return []*http.Cookie
 * @return error
 *
 */
func Uamtk(client *http.Client, url string, cookies []*http.Cookie) (*Result, []*http.Cookie, error) {
	header := http.Header{
		"Accept":           {"application/json, text/javascript, */*; q=0.01"},
		"Accept-Encoding":  {"deflate, br"},
		"Accept-Language":  {"zh-CN, zh; q = 0.9"},
		"Cache-Control":    {"no-cache"},
		"Connection":       {"keep-alive"},
		"Content-Type":     {"application/x-www-form-urlencoded; charset=UTF-8"},
		"Host":             {"kyfw.12306.cn"},
		"Pragma":           {"no-cache"},
		"Referer":          {"https://kyfw.12306.cn/otn/login/init"},
		"X-Requested-With": {"XMLHttpRequest"},
	}

	return uamtk(client, url, header, cookies)
}

/*
 * Login
 *
 * 账户登录

 * @param client 	*http.Client	http客户端
 * @param url 		string	 		url
 * @param cookies 	[]*http.Cookie	cookies
 * @param username 	string			用户名
 * @param password 	string			用户密码
 *
 * @return *Result
 * @return []*http.Cookie
 * @return bool		true - 登录成功 | false - 登录失败
 * @return error
 *
 */
func Login(client *http.Client, url string, cookies []*http.Cookie, username string, password string) (*Result, []*http.Cookie, bool, error) {
	header := http.Header{
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

	client.SetHeader(header)
	client.AddCookies(cookies)

	form := make(url2.Values)
	form["username"] = []string{username}
	form["password"] = []string{password}
	form["appid"] = []string{"otn"}

	resp, err := client.AsyncPostForm(url, form)
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
		return nil, _cookies, false, err
	}

	succeed := false

	if result.ResultCode == 0 {
		succeed = true
	}

	return result, _cookies, succeed, nil
}

/*
 * UserLogin
 *
 * 账户登录后重定向
 *
 * 自动重定向到: https://kyfw.12306.cn/otn/passport?redirect=/otn/login/userLogin
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string	 		url
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return []*http.Cookie
 * @return error
 *
 */
func UserLogin(client *http.Client, url string, cookies []*http.Cookie) ([]*http.Cookie, error) {
	header := http.Header{
		"Accept":           {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"},
		"Accept-Encoding":  {"deflate, br"},
		"Accept-Language":  {"zh-CN, zh; q = 0.9"},
		"Cache-Control":    {"no-cache"},
		"Connection":       {"keep-alive"},
		"Content-Type":     {"application/x-www-form-urlencoded"},
		"Host":             {"kyfw.12306.cn"},
		"Origin":           {"https://kyfw.12306.cn"},
		"Pragma":           {"no-cache"},
		"Referer":          {"https://kyfw.12306.cn/otn/login/init"},
		"X-Requested-With": {"XMLHttpRequest"},
	}
	if client == nil {
		return nil, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)
	client.DelCookie("_passport_ct")

	form := make(url2.Values)
	form["_json_att"] = []string{}

	resp, err := client.AsyncPostForm(url, form)

	if err != nil {
		return nil, err
	}

	_cookies := resp.Cookies()
	client.AddCookies(_cookies)

	if resp.StatusCode != 200 {
		s := fmt.Sprintf("http status code is %d", resp.StatusCode)
		return _cookies, errors.New(s)
	}

	return _cookies, nil
}

/*
 * UamtkLogin
 *
 * 在Login后调用用于获取tk值
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return *Result
 * @return []*http.Cookie
 * @return error
 *
 */
func UamtkLogin(client *http.Client, url string, cookies []*http.Cookie) (*Result, []*http.Cookie, error) {
	header := http.Header{
		"Accept":           {"application/json, text/javascript, */*; q=0.01"},
		"Accept-Encoding":  {"deflate, br"},
		"Accept-Language":  {"zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3"},
		"Connection":       {"keep-alive"},
		"Content-Type":     {"application/x-www-form-urlencoded; charset=UTF-8"},
		"Host":             {"kyfw.12306.cn"},
		"Pragma":           {"no-cache"},
		"Referer":          {"https://kyfw.12306.cn/otn/passport?redirect=/otn/login/userLogin"},
		"X-Requested-With": {"XMLHttpRequest"},
	}
	return uamtk(client, url, header, cookies)
}

/*
 * UamauthClient
 *
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param cookies 	[]*http.Cookie	cookies
 * @param tk		string
 *
 * @return *Result
 * @return []*http.Cookie
 * @return error
 *
 */
func UamauthClient(client *http.Client, url string, cookies []*http.Cookie, tk string) (*Result, []*http.Cookie, error) {
	header := http.Header{
		"Accept":           {"*/*"},
		"Accept-Encoding":  {"deflate, br"},
		"Accept-Language":  {"zh-CN, zh; q = 0.9"},
		"Cache-Control":    {"no-cache"},
		"Connection":       {"keep-alive"},
		"Content-Type":     {"application/x-www-form-urlencoded; charset=UTF-8"},
		"Origin":           {"https://kyfw.12306.cn"},
		"Host":             {"kyfw.12306.cn"},
		"Referer":          {"https://kyfw.12306.cn/otn/passport?redirect=/otn/login/userLogin"},
		"X-Requested-With": {"XMLHttpRequest"},
	}
	if client == nil {
		return nil, nil, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)
	client.DelCookie("_passport_session")

	form := make(url2.Values)
	form["tk"] = []string{tk}

	resp, err := client.AsyncPostForm(url, form)
	if err != nil {
		return nil, nil, err
	}

	_cookies := resp.Cookies()
	client.AddCookies(_cookies)

	if resp.StatusCode != 200 {
		s := fmt.Sprintf("http status code is %d", resp.StatusCode)
		return nil, nil, errors.New(s)
	}


	result := &Result{
		ResultMessage: "新对象",
	}

	err = json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, nil, err
	}

	return result, _cookies, nil
}

/*
 * UserLoginSec
 *
 * 账户登录后重定向
 *
 * 自动重定向到: https://kyfw.12306.cn/otn/index/initMy12306
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string	 		url
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return []*http.Cookie
 * @return error
 *
 */
func UserLoginSec(client *http.Client, url string, cookies []*http.Cookie) ([]*http.Cookie, error) {
	header := http.Header{
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"},
		"Accept-Encoding":           {"deflate, br"},
		"Accept-Language":           {"zh-CN, zh; q = 0.9"},
		"Cache-Control":             {"no-cache"},
		"Connection":                {"keep-alive"},
		"Content-Type":              {"application/x-www-form-urlencoded"},
		"Host":                      {"kyfw.12306.cn"},
		"Pragma":                    {"no-cache"},
		"Referer":                   {"https://kyfw.12306.cn/otn/passport?redirect=/otn/login/userLogin"},
		"Upgrade-Insecure-Requests": {"1"},
	}
	if client == nil {
		return nil, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)

	resp, err := client.AsyncGet(url)

	if err != nil {
		return nil, err
	}

	_cookies := resp.Cookies()
	client.AddCookies(_cookies)

	if resp.StatusCode != 200 {
		s := fmt.Sprintf("http status code is %d", resp.StatusCode)
		return _cookies, errors.New(s)
	}

	return _cookies, nil
}

/*
 * InitMy12306
 *
 * 账户登录后的首页
 *
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string	 		url
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return []*http.Cookie
 * @return error
 *
 */
func InitMy12306(client *http.Client, url string, cookies []*http.Cookie) ([]*http.Cookie, error) {
	header := http.Header{
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"},
		"Accept-Encoding":           {"deflate, br"},
		"Accept-Language":           {"zh-CN, zh; q = 0.9"},
		"Cache-Control":             {"no-cache"},
		"Connection":                {"keep-alive"},
		"Host":                      {"kyfw.12306.cn"},
		"Pragma":                    {"no-cache"},
		"Referer":                   {"https://kyfw.12306.cn/otn/passport?redirect=/otn/login/userLogin"},
		"Upgrade-Insecure-Requests": {"1"},
	}
	if client == nil {
		return nil, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)

	resp, err := client.AsyncGet(url)

	if err != nil {
		return nil, err
	}

	_cookies := resp.Cookies()
	client.AddCookies(_cookies)

	if resp.StatusCode != 200 {
		s := fmt.Sprintf("http status code is %d", resp.StatusCode)
		return _cookies, errors.New(s)
	}

	return _cookies, nil
}
