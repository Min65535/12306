package test

import (
	"me.org/http"
	"testing"
	"log"
)

func Test_Init(t *testing.T) {
	resp, err := http.Get("https://kyfw.12306.cn/otn/login/init#")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Body)
}



func Test_GetCaptcha(t *testing.T) {
	resp, err := http.Get("https://kyfw.12306.cn/passport/captcha/captcha-image?login_site=E&module=login&rand=sjrand")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Body)
}