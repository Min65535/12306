package web

import (
	"me.org/http"
	url2 "net/url"
	json "github.com/json-iterator/go"
	"errors"
	"regexp"
	"log"
	"time"
	"strconv"
	"strings"
)

type CheckDataJson struct {
	Flag bool `json:"flag"`
}
type CheckUserJson struct {
	ValidateMessagesShowId string        `json:"validateMessagesShowId"`
	Status                 bool          `json:"status"`
	HttpStatus             int           `json:"httpstatus"`
	Data                   CheckDataJson `json:"data"`
	Messages               []string      `json:"messages"`
	ValidateMessages       interface{}   `json:"validateMessages"`
}

type SubmitOrderRequestJson struct {
	ValidateMessagesShowId string      `json:"validateMessagesShowId"`
	Status                 bool        `json:"status"`
	HttpStatus             int         `json:"httpstatus"`
	Data                   string      `json:"data"`
	Messages               []string    `json:"messages"`
	ValidateMessages       interface{} `json:"validateMessages"`
}

type GetPassengerDTOsJson struct {
	ValidateMessagesShowId string           `json:"validateMessagesShowId"`
	Status                 bool             `json:"status"`
	HttpStatus             int              `json:"httpstatus"`
	Data                   GetPassengerData `json:"data"`
	Messages               []string         `json:"messages"`
	ValidateMessages       interface{}      `json:"validateMessages"`
}

type GetPassengerData struct {
	IsExist          bool                `json:"isExist"`
	ExMsg            string              `json:"exMsg"`
	TwoIsOpenClick   []string            `json:"two_isOpenClick"`
	OtherIsOpenClick []string            `json:"other_isOpenClick"`
	NormalPassengers NormalPassengersAll `json:"normal_passengers"`
	DjPassengers     interface{}         `json:"dj_passengers"`
}

type NormalPassengersAll []NormalPassenger

type NormalPassenger struct {
	Code                string `json:"code"`
	PassengerName       string `json:"passenger_name"`
	SexCode             string `json:"sex_code"`
	SexName             string `json:"sex_name"`
	BornDate            string `json:"born_date"`
	CountryCode         string `json:"country_code"`
	PassengerIdTypeCode string `json:"passenger_id_type_code"`
	PassengerIdTypeName string `json:"passenger_id_type_name"`
	PassengerIdNo       string `json:"passenger_id_no"`
	PassengerType       string `json:"passenger_type"`
	PassengerFlag       string `json:"passenger_flag"`
	PassengerTypeName   string `json:"passenger_type_name"`
	MobileNo            string `json:"mobile_no"`
	PhoneNo             string `json:"phone_no"`
	Email               string `json:"email"`
	Address             string `json:"address"`
	PostalCode          string `json:"postalcode"`
	FirstLetter         string `json:"first_letter"`
	RecordCount         string `json:"recordCount"`
	TotalTimes          string `json:"total_times"`
	IndexId             string `json:"index_id"`
}

type CheckOrderInfoJson struct {
	ValidateMessagesShowId string         `json:"validateMessagesShowId"`
	Status                 bool           `json:"status"`
	HttpStatus             int            `json:"httpstatus"`
	Data                   CheckOrderJson `json:"data"`
	Messages               []string       `json:"messages"`
	ValidateMessages       interface{}    `json:"validateMessages"`
}

type CheckOrderJson struct {
	IfShowPassCode     string `json:"ifShowPassCode"`
	CanChooseBeds      string `json:"canChooseBeds"`
	CanChooseSeats     string `json:"canChooseSeats"`
	ChooseSeats        string `json:"choose_Seats"`
	IsCanChooseMid     string `json:"isCanChooseMid"`
	IfShowPassCodeTime string `json:"ifShowPassCodeTime"`
	SubmitStatus       bool   `json:"submitStatus"`
	SmokeStr           string `json:"smokeStr"`
}
type GetQueueCountJson struct {
	ValidateMessagesShowId string       `json:"validateMessagesShowId"`
	Status                 bool         `json:"status"`
	HttpStatus             int          `json:"httpstatus"`
	Data                   GetQueueJson `json:"data"`
	Messages               []string     `json:"messages"`
	ValidateMessages       interface{}  `json:"validateMessages"`
}

type GetQueueJson struct {
	Count  string `json:"count"`
	Ticket string `json:"ticket"`
	Op2    string `json:"op_2"`
	CountT string `json:"countT"`
	Op1    string `json:"op_1"`
}

type QueryOrderWaitTimeJson struct {
	ValidateMessagesShowId string         `json:"validateMessagesShowId"`
	Status                 bool           `json:"status"`
	HttpStatus             int            `json:"httpstatus"`
	Data                   QueryOrderJson `json:"data"`
	Messages               []string       `json:"messages"`
	ValidateMessages       interface{}    `json:"validateMessages"`
}

type QueryOrderJson struct {
	QueryOrderWaitTimeStatus bool        `json:"queryOrderWaitTimeStatus"`
	Count                    int         `json:"count"`
	WaitTime                 int         `json:"waitTime"`
	RequestId                int         `json:"requestId"`
	WaitCount                int         `json:"waitCount"`
	TourFlag                 string      `json:"tourFlag"`
	OrderId                  interface{} `json:"orderId"`
}

/*
 * CheckUser
 *
 * 验证登录
 *
 * @param client 	*http.Client	http客户端
 * @param url 		string			url
 * @param cookies 	[]*http.Cookie	cookies
 *
 * @return bool		成功标识
 * @return error
 *
 */
func CheckUser(client *http.Client, url string, cookies []*http.Cookie) (bool, error) {
	header := http.Header{
		"Accept":            {"*/*"},
		"Accept-Encoding":   {"deflate, br"},
		"Accept-Language":   {"zh-CN, zh; q = 0.9"},
		"Cache-Control":     {"no-cache"},
		"Connection":        {"keep-alive"},
		"Content-Type":      {"application/x-www-form-urlencoded; charset=UTF-8"},
		"Host":              {"kyfw.12306.cn"},
		"If-Modified-Since": {"0"},
		"Origin":            {"https://kyfw.12306.cn"},
		"Pragma":            {"no-cache"},
		"Referer":           {"https://kyfw.12306.cn/otn/leftTicket/init"},
		"X-Requested-With":  {"XMLHttpRequest"},
	}
	if client == nil {
		return false, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)
	form := make(url2.Values)
	form["_json_att"] = []string{}

	resp, err := client.AsyncPostForm(url, form)

	if err != nil {
		return false, err
	}

	str := new(CheckUserJson)
	err = json.Unmarshal(resp.Body, str)
	if err != nil {
		return false, err
	}

	if str.Status == false{
		err := errors.New(string(str.Messages[0]))
		return false, err
	}
	return true, nil
}

/*
 * SubmitOrderRequest
 *
 * 预提交订单
 *
 * @param client 			*http.Client	http客户端
 * @param url 				string			url
 * @param cookies 			[]*http.Cookie	cookies
 * @param secretStr 		string			从LeftTicketQuery取出的行程密钥
 * @param trainDate 		string			购票车次日期
 * @param backTrainDate 	string			回程日期
 * @param tourFlag 			string			行程类型，例：dc,表示单程;wc,表示往返程
 * @param personType 		string			乘客类型
 * @param fromStation 		string			始发站
 * @param toStation 		string			目的站
 *
 * @return bool		成功标识
 * @return error
 *
 */
func SubmitOrderRequest(client *http.Client, url string, cookies []*http.Cookie, secretStr string, trainDate string, backTrainDate string, tourFlag string, personType string, fromStation string, toStation string) (bool, error) {
	header := http.Header{
		"Accept":           {"*/*"},
		"Accept-Encoding":  {"deflate, br"},
		"Accept-Language":  {"zh-CN, zh; q = 0.9"},
		"Cache-Control":    {"no-cache"},
		"Connection":       {"keep-alive"},
		"Content-Type":     {"application/x-www-form-urlencoded; charset=UTF-8"},
		"Host":             {"kyfw.12306.cn"},
		"Origin":           {"https://kyfw.12306.cn"},
		"Pragma":           {"no-cache"},
		"Referer":          {"https://kyfw.12306.cn/otn/leftTicket/init"},
		"X-Requested-With": {"XMLHttpRequest"},
	}
	if client == nil {
		return false, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)

	if backTrainDate == "" {
		w := time.Now()
		str := w.Format("2006-01-02 15:04:05")
		backTrainDate = (strings.Split(str, " "))[0]
	}
	form := http.NewForm()

	/*
	secretStr = "y3dkTHrvzO2n6KG66eOdOh1FIybLWtPVYrfSHbP1N3MNamHYAZrTdH2pDERsI9mjC7K+3hOes+LB0RSDwGkZljeW/bWHhoyrkiEmcwWEU3ZoFggtBMm2b7aGn8XuLSi+L9GY+p8WNBE0kBrY3e8MS0zAfMQzqgBH+MD0je861htDlrVnd0Dd233ykR0G9fy4bglMZTQogTGs0sPv0iXwRAvwzDSEhXbu3rUvYlKnHLR5Ul8cYYzeqg=="
	trainDate = "2017-12-10"
	backTrainDate = "2017-11-22"
	tourFlag = "dc"
	personType = "ADULT"
	fromStation = "深圳"
	toStation = "武汉"
	*/

	form.Add("secretStr", secretStr)
	form.Add("train_date", trainDate)
	form.Add("back_train_date", backTrainDate)
	form.Add("tour_flag", tourFlag)
	form.Add("purpose_codes", personType)
	form.Add("query_from_station_name", fromStation)
	form.Add("query_to_station_name", toStation)
	form.Add("undefined", "")

	resp, err := client.AsyncPostForm(url, form)

	log.Println("SubmitOrderRequest resp:", string(resp.Body))

	if err != nil {
		return false, err
	}
	str := new(SubmitOrderRequestJson)
	err = json.Unmarshal(resp.Body, str)
	if err != nil {
		return false, err
	}

	if str.Status == false {
		err := errors.New(string(str.Messages[0]))
		return false, err
	}

	return true, nil
}

/*
 * InitDc
 *
 * 模拟跳转页面InitDc
 *
 * @param client 	*http.Client		http客户端
 * @param url 		string				url
 * @param cookies 	[]*http.Cookie		cookies
 *
 * @return string	repeatSubmitToken	Token值
 * @return string	keyCheckIsChange	Key值
 * @return error
 *
 */
func InitDc(client *http.Client, url string, cookies []*http.Cookie) (string, string, error) {
	header := http.Header{
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
		"Accept-Encoding":           {"deflate, br"},
		"Accept-Language":           {"zh-CN, zh; q = 0.9"},
		"Cache-Control":             {"no-cache"},
		"Connection":                {"keep-alive"},
		"Content-Type":              {"application/x-www-form-urlencoded"},
		"Host":                      {"kyfw.12306.cn"},
		"Origin":                    {"https://kyfw.12306.cn"},
		"Pragma":                    {"no-cache"},
		"Referer":                   {"https://kyfw.12306.cn/otn/leftTicket/init"},
		"Upgrade-Insecure-Requests": {"1"},
	}
	if client == nil {
		return "", "", ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)
	form := http.NewForm()
	form.Add("_json_att", "")
	resp, err := client.AsyncPostForm(url, form)
	if err != nil {
		return "", "", err
	}
	//log.Println("InitDc respone body:",string(resp.Body))

	reg1 := regexp.MustCompile(`var globalRepeatSubmitToken = '`)
	find1 := reg1.FindIndex(resp.Body)
	if find1 == nil {
		err := errors.New("repeatSubmitToken的第一次截取未能匹配到参数")
		return "", "", err
	}
	index1 := find1[1]
	str := resp.Body[index1:]
	reg2 := regexp.MustCompile(`';`)
	find2 := reg2.FindIndex([]byte(str))
	if find2 == nil {
		err := errors.New("repeatSubmitToken的第二次截取未能匹配到参数")
		log.Println(err)
	}
	index2 := find2[0]
	repeatSubmitToken := string(str[0:index2])

	regSec1 := regexp.MustCompile(`key_check_isChange':'`)
	findSec1 := regSec1.FindIndex(resp.Body)
	if findSec1 == nil {
		err := errors.New("keyCheckIsChange的第一次截取未能匹配到参数")
		return "", "", err
	}
	indexSec1 := findSec1[1]
	strSec1 := resp.Body[indexSec1:]

	regSec2 := regexp.MustCompile(`',`)
	findSec2 := regSec2.FindIndex([]byte(strSec1))
	if findSec2 == nil {
		err := errors.New("keyCheckIsChange的第二次截取未能匹配到参数")
		return "", "", err
	}
	indexSec2 := findSec2[0]
	keyCheckIsChange := string(strSec1[0:indexSec2])

	return repeatSubmitToken, keyCheckIsChange, nil

}

/*
 * GetPassengerDTOs
 *
 * 常用联系人确定
 *
 * @param client 				*http.Client	http客户端
 * @param url 					string			url
 * @param cookies 				[]*http.Cookie	cookies
 * @param repeatSubmitToken 	string			Token值
 *
 * @return *GetPassengerDTOsJson
 * @return error
 *
 */
func GetPassengerDTOs(client *http.Client, url string, cookies []*http.Cookie, repeatSubmitToken string) (*GetPassengerDTOsJson, error) {
	header := http.Header{
		"Accept":           {"*/*"},
		"Accept-Encoding":  {"deflate, br"},
		"Accept-Language":  {"zh-CN, zh; q = 0.9"},
		"Cache-Control":    {"no-cache"},
		"Connection":       {"keep-alive"},
		"Content-Type":     {"application/x-www-form-urlencoded; charset=UTF-8"},
		"Host":             {"kyfw.12306.cn"},
		"Pragma":           {"no-cache"},
		"Referer":          {"https://kyfw.12306.cn/otn/confirmPassenger/initDc"},
		"X-Requested-With": {"XMLHttpRequest"},
	}
	if client == nil {
		return nil, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)
	form := http.NewForm()
	form.Add("_json_att", "")
	form.Add("REPEAT_SUBMIT_TOKEN", repeatSubmitToken)
	resp, err := client.AsyncPostForm(url, form)
	if err != nil {
		return nil, err
	}

	str := new(GetPassengerDTOsJson)
	err = json.Unmarshal(resp.Body, str)
	if err != nil {
		return nil, err
	}

	if str.Status == false || str.Messages == nil {
		err := errors.New(str.Messages[0])
		return nil, err
	}
	return str, nil
}

/*
 * GetPassCodeNewSubmit
 *
 * 获取验证码
 *
 * @param client 				*http.Client	http客户端
 * @param url 					string			url
 * @param cookies 				[]*http.Cookie	cookies
 *
 * @return []*http.Cookie
 * @return error
 *
 */
func GetPassCodeNewSubmit(client *http.Client, url string, cookies []*http.Cookie) ([]*http.Cookie, error) {
	header := http.Header{
		"Accept":          {"image/webp,image/apng,image/*,*/*;q=0.8"},
		"Accept-Encoding": {"deflate, br"},
		"Accept-Language": {"zh-CN, zh; q = 0.9"},
		"Cache-Control":   {"no-cache"},
		"Connection":      {"keep-alive"},
		"Host":            {"kyfw.12306.cn"},
		"Pragma":          {"no-cache"},
		"Referer":         {"https://kyfw.12306.cn/otn/confirmPassenger/initDc"},
	}
	length := "17"
	return GetPassCodeNewCommon(client, url, header, cookies, length)
}

/*
 * CheckOrderInfo
 *
 * 购票人确定
 *
 * @param client 				*http.Client	http客户端
 * @param url 					string			url
 * @param cookies 				[]*http.Cookie	cookies
 * @param passengerTicketStr 	string			乘客预定座位信息和乘客信息
 * @param oldPassengerStr 		string			乘客信息
 * @param tourFlag 				string			行程类型:dc,表示单程;wc,表示往返
 * @param repeatSubmitToken 	string			Token值
 *
 * @return *CheckOrderInfoJson
 * @return error
 *
 */
func CheckOrderInfo(client *http.Client, url string, cookies []*http.Cookie, passengerTicketStr string, oldPassengerStr string, tourFlag string, repeatSubmitToken string) (*CheckOrderInfoJson, error) {
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
		"Referer":          {"https://kyfw.12306.cn/otn/confirmPassenger/initDc"},
		"X-Requested-With": {"X-Requested-With"},
	}

	if client == nil {
		return nil, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)
	form := http.NewForm()
	form.Add("cancel_flag", "2")
	form.Add("bed_level_order_num", "000000000000000000000000000000")

	/*
		座位编号（seatType）参考：

		‘硬卧’ => ‘3’,
		‘软卧’ => ‘4’,
		‘二等座’ => ‘O’,
		‘一等座’ => ‘M’,
		‘硬座’ => ‘1’,
		‘商务座’ => ‘9’,

		passengerTicketStr组成的格式：seatType,0,票类型（成人票填1）,乘客名,passenger_id_type_code,passenger_id_no,mobile_no,’N’//多个乘车人用’_’隔开

		oldPassengerStr组成的格式：乘客名,passenger_id_type_code,passenger_id_no,passenger_type，’_’ //多个乘车人用’_’隔开，注意最后的需要多加一个’_’
	 */

	//3,0,1,杜方敏,1,421022199411011854,17512032601,N_3,0,1,崔圣雨,1,420621199304263816,18086507020,N_3,0,1,杜先强,1,421022197205221835,13509694716,N

	//3,0,1,杜方敏,1,421022199411011854,17512032601,N

	//3,0,1,刘芬,1,421022199408110341,15307149873,N_3,0,1,杜方敏,1,421022199411011854,17512032601,N

	form.Add("passengerTicketStr", passengerTicketStr)

	//杜方敏,1,421022199411011854,3_崔圣雨,1,420621199304263816,3_杜先强,1,421022197205221835,1_

	//杜方敏,1,421022199411011854,3_

	//刘芬,1,421022199408110341,3_杜方敏,1,421022199411011854,3_

	form.Add("oldPassengerStr", oldPassengerStr)
	form.Add("tour_flag", tourFlag)
	form.Add("randCode", "")
	form.Add("_json_att", "")
	form.Add("REPEAT_SUBMIT_TOKEN", repeatSubmitToken)
	resp, err := client.AsyncPostForm(url, form)
	if err != nil {
		return nil, err
	}
	str := new(CheckOrderInfoJson)

	err = json.Unmarshal(resp.Body, str)
	if err != nil {
		return nil, err
	}

	if str.Status == false{
		err := errors.New(string(str.Messages[0]))
		return nil, err
	}
	return str, nil
}

/*
 * GetQueueCount
 *
 * 准备进入排队
 *
 * @param client 					*http.Client	http客户端
 * @param url 						string			url
 * @param cookies 					[]*http.Cookie	cookies
 * @param trainDate 				string			乘车日期//Sun+Dec+10+2017+00:00:00+GMT+0800
 * @param trainNo 					string			列车编号//690000K2380A
 * @param stationTrainCode 			string			列车车次//K238
 * @param seatType 					string			座位类型//3,硬卧;4,软卧;0,二等座;M,一等座;1,硬座;
 * @param fromStationTelecode 		string			始发站的三字码
 * @param toStationTelecode 		string			终点站的三字码
 * @param leftTicket 				string			LeftTicketQuery里的余票短密钥       //Wey1KErtCNm8%2B1Bt55Vy9F9hqwuUR0KY%2F2P5DarryhPxyPnkDbWGDEP775g%3D
 * @param trainLocation 			string			LeftTicketQuery里的短密钥后面第三段里的Q7        //|20171210|3|Q7|01|11|0|0||||10|||有||有|有|||||10401030|1413
 * @param repeatSubmitToken 		string			Token值
 *
 * @return *GetQueueCountJson
 * @return error
 *
 */
func GetQueueCount(client *http.Client, url string, cookies []*http.Cookie, trainDate string, trainNo string, stationTrainCode string, seatType string, fromStationTelecode string, toStationTelecode string, leftTicket string, trainLocation string, repeatSubmitToken string) (*GetQueueCountJson, error) {
	header := http.Header{
		"Accept":           {"application/json, text/javascript, */*; q=0.01"},
		"Accept-Encoding":  {"deflate, br"},
		"Accept-Language":  {"zh-CN, zh; q = 0.9"},
		"Cache-Control":    {"no-cache"},
		"Connection":       {"keep-alive"},
		"Content-Type":     {"application/x-www-form-urlencoded; charset=UTF-8"},
		"Host":             {"kyfw.12306.cn"},
		"Pragma":           {"no-cache"},
		"Referer":          {"https://kyfw.12306.cn/otn/confirmPassenger/initDc"},
		"X-Requested-With": {"X-Requested-With"},
	}

	if client == nil {
		return nil, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)
	form := http.NewForm()
	form.Add("train_date", trainDate)                    //Sun+Dec+10+2017+00:00:00+GMT+0800
	form.Add("train_no", trainNo)                        //690000K2380A
	form.Add("stationTrainCode", stationTrainCode)       //K238
	form.Add("seatType", seatType)                       //3
	form.Add("fromStationTelecode", fromStationTelecode) //OSQ
	form.Add("toStationTelecode", toStationTelecode)     //WCN
	form.Add("leftTicket", leftTicket)                   //Wey1KErtCNm8%2B1Bt55Vy9F9hqwuUR0KY%2F2P5DarryhPxyPnkDbWGDEP775g%3D
	form.Add("purpose_codes", "00")                      //00
	form.Add("train_location", trainLocation)            //Q7
	form.Add("_json_att", "")                            //
	form.Add("REPEAT_SUBMIT_TOKEN", repeatSubmitToken)   //Token值

	resp, err := client.AsyncPostForm(url, form)
	if err != nil {
		return nil, err
	}
	str := new(GetQueueCountJson)

	err = json.Unmarshal(resp.Body, str)
	if err != nil {
		return nil, err
	}

	if str.Status == false {
		err := errors.New(string(str.Messages[0]))
		return nil, err
	}
	return str, nil

}

/*
 * ConfirmSingleForQueue
 *
 * 确认购买
 *
 * @param client 				*http.Client	http客户端
 * @param url 					string			url
 * @param cookies 				[]*http.Cookie	cookies
 * @param passengerTicketStr 	string			乘客预定座位信息和乘客信息
 * @param oldPassengerStr 		string			乘客信息
 * @param keyCheckIsChange 		string			Key值
 * @param leftTicket 			string			LeftTicketQuery里的余票短密钥       //Wey1KErtCNm8%2B1Bt55Vy9F9hqwuUR0KY%2F2P5DarryhPxyPnkDbWGDEP775g%3D
 * @param trainLocation 		string			LeftTicketQuery里的短密钥后面第三段里的Q7        //|20171210|3|Q7|01|11|0|0||||10|||有||有|有|||||10401030|1413
 * @param repeatSubmitToken 	string			Token值
 *
 * @return *CommonJson
 * @return error
 *
 */
func ConfirmSingleForQueue(client *http.Client, url string, cookies []*http.Cookie, passengerTicketStr string, oldPassengerStr string, keyCheckIsChange string, leftTicket string, trainLocation string, repeatSubmitToken string) (*CommonJson, error) {
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
		"Referer":          {"https://kyfw.12306.cn/otn/confirmPassenger/initDc"},
		"X-Requested-With": {"XMLHttpRequest"},
	}

	if client == nil {
		return nil, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)
	form := http.NewForm()
	form.Add("passengerTicketStr", passengerTicketStr)
	form.Add("oldPassengerStr", oldPassengerStr)
	form.Add("randCode", "")
	form.Add("purpose_codes", "00")
	form.Add("key_check_isChange", keyCheckIsChange)
	form.Add("leftTicketStr", leftTicket)
	form.Add("train_location", trainLocation)
	form.Add("choose_seats", "")
	form.Add("seatDetailType", "000")
	form.Add("roomType", "00")
	form.Add("dwAll", "N")
	form.Add("_json_att", "")
	form.Add("REPEAT_SUBMIT_TOKEN", repeatSubmitToken)

	resp, err := client.AsyncPostForm(url, form)
	if err != nil {
		return nil, err
	}
	str := new(CommonJson)
	err = json.Unmarshal(resp.Body, str)
	if err != nil {
		return nil, err
	}

	if str.Status == false{
		err := errors.New(string(str.Messages[0]))
		return nil, err
	}
	return str, nil
}

/*
 * QueryOrderWaitTime
 *
 * 循环提交
 *
 * @param client 				*http.Client	http客户端
 * @param url 					string			url
 * @param cookies 				[]*http.Cookie	cookies
 * @param tourFlag 				string			行程类型:dc,表示单程;wc,表示往返
 * @param repeatSubmitToken 	string			Token值
 *
 * @return *CommonJson
 * @return error
 *
 */
func QueryOrderWaitTime(client *http.Client, url string, cookies []*http.Cookie, tourFlag string, repeatSubmitToken string) (*QueryOrderWaitTimeJson, error) {
	header := http.Header{
		"Accept":           {"application/json, text/javascript, */*; q=0.01"},
		"Accept-Encoding":  {"deflate, br"},
		"Accept-Language":  {"zh-CN, zh; q = 0.9"},
		"Cache-Control":    {"no-cache"},
		"Connection":       {"keep-alive"},
		"Host":             {"kyfw.12306.cn"},
		"Pragma":           {"no-cache"},
		"Referer":          {"https://kyfw.12306.cn/otn/confirmPassenger/initDc"},
		"X-Requested-With": {"XMLHttpRequest"},
	}

	if client == nil {
		return nil, ClientIsNilErr
	}

	client.SetHeader(header)
	client.AddCookies(cookies)

	milliSecondStr := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	query := http.NewQuery()
	query.Add("random", milliSecondStr)
	query.Add("tourFlag", tourFlag)
	query.Add("_json_att", "")
	query.Add("REPEAT_SUBMIT_TOKEN", repeatSubmitToken)
	resp, err := client.AsyncGet(url, query)
	if err != nil {
		return nil, err
	}

	log.Println("QueryOrderWaitTime resp:", string(resp.Body))

	str := new(QueryOrderWaitTimeJson)
	err = json.Unmarshal(resp.Body, str)
	if err != nil {
		return nil, err
	}

	if str.Status == false{
		err := errors.New(string(str.Messages[0]))
		return nil, err
	}
	return str, nil
}
