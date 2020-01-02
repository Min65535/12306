package captcha

import (
	"log"
	"12306/captcha/jsdama"
)

/*
 * 初始化
 */
func init() {
	myCaptcha := GetCaptcha()
	jsdama12306 := jsdama.NewJsdama12306(8184, "z7fB0y0GWsD8wmirBX1mLljAzMsgNUGWNobIEQDZ", "xxxxtrip", "xxxx123")
	ret := myCaptcha.Register(jsdama12306.Name, jsdama12306)
	log.Println("Captcha Register:", jsdama12306.Name, ret)
}
