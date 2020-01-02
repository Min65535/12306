package web

import (
	"me.org/http"
	json "github.com/json-iterator/go"
	"log"
	"fmt"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	url2 "net/url"
	_ "net/http"
	"time"
)

/*
 * LeftTiceketInit
 *
 * 余票查询首页
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return []*http.Cookie
 * @return error
 *
 */
func LeftTicketInit(client *http.Client, url string, cookies []*http.Cookie) ([]*http.Cookie, error) {
	header := http.Header{
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"},
		"Accept-Encoding":           {"deflate, br"},
		"Accept-Language":           {"zh-CN, zh; q = 0.9"},
		"Cache-Control":             {"no-cache"},
		"Connection":                {"keep-alive"},
		"Host":                      {"kyfw.12306.cn"},
		"Pragma":                    {"no-cache"},
		"Referer":                   {"https://kyfw.12306.cn/otn/index/initMy12306"},
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
 * GetPassCodeNewCommon
 *
 * 获取验证码

 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param cookies 	[]*http.Cookie	cookies
 * @param length 	string			随机数的长度
 *
 * @return []*http.Cookie
 * @return error
 *
 */
func GetPassCodeNewCommon(client *http.Client, url string, header http.Header, cookies []*http.Cookie, length string) ([]*http.Cookie, error) {
	if client == nil {
		return nil, ClientIsNilErr
	}
	client.SetHeader(header)
	client.AddCookies(cookies)
	sjrand := rand.Float64()
	if length == "" {
		url = fmt.Sprintf("%s?module=passenger&rand=randp&%0.16f", url, sjrand)
	} else {
		url = fmt.Sprintf("%s?module=passenger&rand=randp&%0."+length+"f", url, sjrand)
	}
	resp, err := client.AsyncGet(url, nil)
	if err != nil {
		return nil, err
	}
	_cookies := resp.Cookies()
	client.AddCookies(_cookies)
	return _cookies, nil
}

/*
 * GetPassCodeNew
 *
 * 获取验证码

 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param cookies 	[]*http.Cookie	cookies
 * @param length 	string			随机数的长度
 *
 * @return []*http.Cookie
 * @return error
 *
 */
func GetPassCodeNew(client *http.Client, url string, cookies []*http.Cookie) ([]*http.Cookie, error) {
	header := http.Header{
		"Accept":          {"image/webp,image/apng,image/*,*/*;q=0.8"},
		"Accept-Encoding": {"deflate, br"},
		"Accept-Language": {"zh-CN, zh; q = 0.9"},
		"Cache-Control":   {"no-cache"},
		"Connection":      {"keep-alive"},
		"Host":            {"kyfw.12306.cn"},
		"Pragma":          {"no-cache"},
		"Referer":         {"https://kyfw.12306.cn/otn/leftTicket/init"},
	}
	length := ""
	return GetPassCodeNewCommon(client, url, header, cookies, length)
}

/*
 * LeftTicketLog
 *
 * 余票查询确认
 *
 * @param client 		*http.Client	http客户端
 * @param url 			string			url
 * @param cookies 		[]*http.Cookie	cookies
 * @param fromStation 	[]string  		出发地参数,示例:"深圳","SZQ"
 * @param toStation 	[]string  		目的地参数,示例:"武汉","WHN"
 * @param toStation 	string  		目的地参数,示例:"武汉","WHN"
 * @param toStation 	string  		目的地参数,示例:"武汉","WHN"
 *
 * @return bool		成功标志
 * @return error
 *
 */
func LeftTicketLog(client *http.Client, url string, cookies []*http.Cookie, fromStation []string, toStation []string, fromDate string, toDate string, flag string, personType string) (bool, error) {
	header := http.Header{
		"Host":                      {"kyfw.12306.cn"},
		"Accept":                    {"*/*"},
		"Accept-Language":           {"zh-CN, zh; q = 0.9"},
		"Accept-Encoding":           {"gzip, deflate, sdch, br"},
		"Referer":                   {"https://kyfw.12306.cn/otn/leftTicket/init"},
		"If-Modified-Since":         {"0"},
		"Cache-Control":             {"no-cache"},
		"Upgrade-Insecure-Requests": {"1"},
		"Connection":                {"keep-alive"},
		"Pragma":                    {"no-cache"},
	}
	if client == nil {
		return false, ClientIsNilErr
	}
	client.SetHeader(header)
	client.AddCookies(cookies)

	client.DelCookie("uamtk")

	f := strconv.QuoteToASCII(fromStation[0])
	f = strings.Replace(f, `\`, `%`, 2)
	f = f[1:13]
	fromStr := f + url2.QueryEscape(",") + fromStation[1]

	t := strconv.QuoteToASCII(toStation[0])
	t = strings.Replace(t, `\`, `%`, 2)
	t = t[1:13]
	toStr := t + url2.QueryEscape(",") + toStation[1]

	if toDate == "" {
		w := time.Now()
		str := w.Format("2006-01-02 15:04:05")
		toDate = (strings.Split(str, " "))[0]
	}

	client.AddCookie(&http.Cookie{
		Name:  "current_captcha_type",
		Value: "Z",
	})
	client.AddCookie(&http.Cookie{
		Name:  "_jc_save_fromStation",
		Value: fromStr,
	})
	client.AddCookie(&http.Cookie{
		Name:  "_jc_save_toStation",
		Value: toStr,
	})
	client.AddCookie(&http.Cookie{
		Name:  "_jc_save_fromDate",
		Value: fromDate,
	})
	client.AddCookie(&http.Cookie{
		Name:  "_jc_save_toDate",
		Value: toDate,
	})
	client.AddCookie(&http.Cookie{
		Name:  "_jc_save_wfdc_flag",
		Value: flag,
	})

	log.Println("LeftTicketLog Header:", client.SafeHeader)

	query := http.NewQuery()

	query.Add("leftTicketDTO.train_date", fromDate)
	query.Add("leftTicketDTO.from_station", fromStation[1])
	query.Add("leftTicketDTO.to_station", toStation[1])
	query.Add("purpose_codes", personType)

	resp, err := client.AsyncGet(url, query)
	if err != nil {
		return false, err
	}

	str := new(ResponseJson)
	err1 := json.Unmarshal(resp.Body, str)
	if err1 != nil {
		return false, err1
	}

	if str.Status == false || str.Messages == nil {
		err := errors.New(str.Messages[0])
		return false, err
	}
	return true, nil
}

/*
 * LeftTicketQuery
 *
 * 余票查询
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return MsgJson	查询出来的具体参数
 * @return error
 *
 */
func LeftTicketQuery(client *http.Client, url string, cookies []*http.Cookie, fromStation []string, toStation []string, fromDate string, personType string) (*MsgJson, error) {
	header := http.Header{
		"Host":                      {"kyfw.12306.cn"},
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"},
		"Accept-Language":           {"zh-CN, zh; q = 0.9"},
		"Accept-Encoding":           {"gzip, deflate, sdch, br"},
		"Referer":                   {"https://kyfw.12306.cn/otn/leftTicket/init"},
		"If-Modified-Since":         {"0"},
		"Cache-Control":             {"no-cache"},
		"Upgrade-Insecure-Requests": {"1"},
		"Connection":                {"keep-alive"},
		"Pragma":                    {"no-cache"},
	}
	if client == nil {
		return nil, ClientIsNilErr
	}
	client.SetHeader(header)
	client.AddCookies(cookies)

	realUrl := url + "?leftTicketDTO.train_date=" + fromDate + "&leftTicketDTO.from_station=" + fromStation[1] + "&leftTicketDTO.to_station=" + toStation[1] + "&purpose_codes=" + personType

	resp, err := client.Get(realUrl)
	if err != nil {
		return nil, err
	}
	str := new(MsgJson)
	err = json.Unmarshal(resp.Body, &str)
	if err != nil {
		return nil, err
	}
	return str, nil
}
