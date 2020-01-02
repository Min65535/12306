package web

import (
	"testing"
)

func TestSingerOrder(t *testing.T) {
	//12306账号密码
	Username := ""
	Password := ""
	fromStation := []string{0: "深圳北", 1: "IOQ"}
	toStation := []string{0: "岳阳东", 1: "YIQ"}
	fromDate := "2018-02-10"
	toDate := ""
	tourFlag := "dc"
	personType := "ADULT"
	stationTrainCode := "G1008"
	/*
	*座位类型
	*硬卧 		value:3
	*软卧 		value:4
	*二等座 		value:O
	*一等座 		value:M
	*硬座 		value:1
	*商务座 		value:9
	*/
	/*
	 *票类型
	 *成人票 	value:1
	 *儿童票 	value:2  //儿童票的名字应填带儿童的那个成人的名字
	 *学生票 	value:3
	 *残军票 	value:4
	 */
	/*
	 *高铁买了学生票就只能买二等座
	 *普通买了学生票不能买软卧
	 *学生不能添加儿童票
	 *一次单最多只能买5张票
	 *
	 */

	passengers := []string{0: "3,1,张三", 1: "3,1,李四", 2: "3,2,王五"} //座位类型,票类型（成人票填1）,乘客名
	SingerOrder(Username, Password, fromStation, toStation, fromDate, toDate, tourFlag, personType, stationTrainCode, passengers)

	/*fromDate := "2017-12-10"
	fromDateTemp := fromDate + ` 00:00:00`
	date, _ := time.Parse("2006-01-02 15:04:05", fromDateTemp)
	weekStr := date.Weekday().String()
	monthStr := date.Month().String()
	dayStr := strconv.Itoa(date.Day())
	yearStr := strconv.Itoa(date.Year())
	lastStr := weekStr[:3] + `+` + monthStr[:3] + `+` + dayStr + `+` + yearStr + `+00:00:00+GMT+0800`
	log.Println(lastStr)*/

	/*s := "NH5NdrdCqfQfzXmGPdxzDlZLmxVRaXMv9LNfx7T0gnY87fzl6jcflkDNPxZsM9NMAKRUu7vOWHau%0A0uU5hOZlttmxTAcB%2FnkmLNRF9zydIzE7mb8Xr7DtfFwM%2B%2F8YqrUMrURCVuhJ%2B2CDXyC%2FHfRtSMaZ%0AIrnBx8qOxPIZmqvHCszeX5fVpt51KF5BgPNod%2Bhg026TEmQg2Vzvm6E7lnS6LN2nXWDb20qUbWVo%0AxzgJK266h%2F0sPsay0A%3D%3D|预订|690000K2380A|K238|OSQ|TYV|OSQ|WCN|15:40|08:52|17:12|Y|Wey1KErtCNm8%2B1Bt55Vy9F9hqwuUR0KY%2F2P5DarryhPxyPnkDbWGDEP775g%3D|20171210|3|Q7|01|11|0|0||||10|||有||有|有|||||10401030|1413"
	str := "K238"
	fmt.Println(strings.Contains(s, str))

	arr := strings.Split(s, "|")

	secretStr := arr[0]
	log.Println(secretStr)

	w := `y3dkTHrvzO2n6KG66eOdOh1FIybLWtPVYrfSHbP1N3MNamHYAZrTdH2pDERsI9mjC7K+3hOes+LB0RSDwGkZljeW/bWHhoyrkiEmcwWEU3ZoFggtBMm2b7aGn8XuLSi+L9GY+p8WNBE0kBrY3e8MS0zAfMQzqgBH+MD0je861htDlrVnd0Dd233ykR0G9fy4bglMZTQogTGs0sPv0iXwRAvwzDSEhXbu3rUvYlKnHLR5Ul8cYYzeqg==`;
	log.Println("w的长度：", len(w))

	secretStr_t := strings.Replace(secretStr, `%0A`, ``, -1)

	log.Println(secretStr_t)

	ww, err := url.QueryUnescape(secretStr_t)
	if err != nil {
		log.Println(err)
	}
	log.Println("ww：", ww)
	log.Println("ww的长度：", len(ww))
*/


}
