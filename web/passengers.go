package web

import (
	"me.org/http"
	url2 "net/url"
	"fmt"
	"regexp"
	json "github.com/json-iterator/go"
	"strings"
	"errors"
)

type Passenger struct {
	TypeName   string `json:"passenger_type_name"`    //乘客类型
	DeleteTime string `json:"delete_time"`            //删除的最迟时间
	IsUserSelf string `json:"isUserSelf"`             //是否是账户拥有者本人
	IdTypeCode string `json:"passenger_id_type_code"` //乘客证件类型码
	Name       string `json:"passenger_name"`         //乘客姓名
	TotalTimes string `json:"total_times"`            //通过标志
	IdTypeName string `json:"passenger_id_type_name"` //证件类型名字
	Type       string `json:"passenger_type"`         //乘客类型码
	IdNum      string `json:"passenger_id_no"`        //证件号码
	MobileNum  string `json:"mobile_no"`              //手机号
}
type Passengers []*Passenger

func (this *Passenger) CheckStatus() (string, error) {
	if this.IdTypeCode == "2" {
		return "未通过", errors.New("未通过:本网站不再支持一代居民身份证，请更改为二代居民身份证")
	} else {
		if this.TotalTimes == "95" || this.TotalTimes == "97" {
			return "已通过", nil
		} else {
			if this.TotalTimes == "93" || this.TotalTimes == "99" {
				if this.IdTypeCode == "1" {
					return "已通过", nil
				} else {
					return "预通过", nil
				}
			} else {
				if this.TotalTimes == "94" || this.TotalTimes == "96" {
					if this.IdTypeCode == "1" {
						return "未通过", errors.New("未通过:根据本网站的服务条款，您需要提供真实、准确的本人资料，您注册的信息与其他用户重复，请您尽快到就近的办理客运售票业务的铁路车站完成身份核验，通过后即可在网上购票，本网站将核验与您身份信息重复的用户，谢谢您的支持。")
					} else {
						return "未通过", errors.New("未通过:身份信息核验未通过，详见《铁路互联网购票身份核验须知》。")
					}
				} else {
					if this.TotalTimes == "92" || this.TotalTimes == "98" {
						if this.IdTypeCode == "B" || this.IdTypeCode == "H" || this.IdTypeCode == "C" || this.IdTypeCode == "G" {
							return "请报验", errors.New("请报验:身份信息未经核验，需持在本网站填写的有效身份证件原件到车站售票窗口办理预核验，详见《铁路互联网购票身份核验须知》。")

						} else {
							return "待核验", errors.New("待核验:身份信息未经核验，需持二代居民身份证原件到车站售票窗口办理核验，详见《铁路互联网购票身份核验须知》。")
						}
					} else {
						if this.TotalTimes == "91" {
							if this.IdTypeCode == "B" || this.IdTypeCode == "H" || this.IdTypeCode == "C" || this.IdTypeCode == "G" {
								return "请报验", errors.New("请报验:身份信息未经核验，需持在本网站填写的有效身份证件原件到车站售票窗口办理预核验，详见《铁路互联网购票身份核验须知》。")
							}
						} else {
							return "请报验", errors.New("请报验:身份信息未经核验，需持在本网站填写的有效身份证件原件到车站售票窗口办理预核验，详见《铁路互联网购票身份核验须知》。")
						}
					}
				}
			}
		}
	}
	return "参数错误", errors.New("参数错误:IdTypeCode " + this.IdTypeCode + ",TotalTimes " + this.TotalTimes)
}

/*
 * passengerInit
 *
 * 常用联系人获取
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param header 	http.Header
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return []*Passenger	所有联系人信息
 * @return []*http.Cookie
 * @return error
 *
 */
func passengerInit(client *http.Client, url string, header http.Header, cookies []*http.Cookie) ([]*Passenger, []*http.Cookie, error) {
	if client == nil {
		return nil, nil, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)

	form := make(url2.Values)
	form["_json_att"] = []string{""}
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
	passengers, err := passengerMsgGet(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return passengers, _cookies, nil
}

/*
 * passengerMsgGet
 *
 * 常用联系人字段处理并返回
 *
 * @param data []byte	联系人首页的返回数据
 *
 * @return []*Passenger	所有联系人信息
 * @return error
 *
 */
func passengerMsgGet(data []byte) ([]*Passenger, error) {
	reg1 := regexp.MustCompile(`var passengers=`)
	find1 := reg1.FindIndex(data)
	if find1 == nil {
		err := errors.New("第一次截取未能匹配到参数")
		return nil, err
	}
	index1 := find1[1]
	str := data[index1:]
	reg2 := regexp.MustCompile(`];`)
	find2 := reg2.FindIndex([]byte(str))
	if find2 == nil {
		err := errors.New("第二次截取未能匹配到参数")
		return nil, err
	}
	index2 := find2[0]
	str2 := str[0:index2+1]

	str3 := strings.Replace(string(str2), "'", "\"", -1)
	passengers := make(Passengers, 0)

	err := json.Unmarshal([]byte(str3), &passengers)
	if err != nil {
		return nil, err
	}
	return passengers, nil
}

/*
 * GetPassengers
 *
 * 常用联系人首页开启获取联系人
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return []*Passenger	所有联系人信息
 * @return []*http.Cookie
 * @return error
 */
func GetPassengers(client *http.Client, url string, cookies []*http.Cookie) ([]*Passenger, []*http.Cookie, error) {
	header := http.Header{
		"Accept":                    {"application/json, text/javascript, */*; q=0.01"},
		"Accept-Encoding":           {"deflate, br"},
		"Accept-Language":           {"zh-CN, zh; q = 0.9"},
		"Connection":                {"keep-alive"},
		"Content-Type":              {"application/x-www-form-urlencoded"},
		"Host":                      {"kyfw.12306.cn"},
		"Referer":                   {"https://kyfw.12306.cn/otn/index/initMy12306"},
		"Upgrade-Insecure-Requests": {"1"},
	}
	return passengerInit(client, url, header, cookies)
}
