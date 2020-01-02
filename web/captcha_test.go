package web

import (
	"testing"
	"log"
	"me.org/http"
	"12306/captcha"
)

func TestGetCaptcha(t *testing.T) {
	client := http.NewClient(nil, "", "", 0)
	image, _, err := GetCaptcha(client, CaptchaUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(len(image))
}

func TestCheckCaptcha(t *testing.T) {
	client := http.NewClient(nil, "", "", 0)
	cookies, err := Init(client, InitUrl)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("init cookies:", cookies)
	result, cookies, err := Uamtk(client, UamtkUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("uamtk result:", result, "cookies:", cookies)

	log.Println("client cookie:", client.Cookie())

	image, cookies, err := GetCaptcha(client, CaptchaUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("client cookie:", client.Cookie())

	myCaptcha := captcha.GetCaptcha()
	id, value, err := myCaptcha.Upload("Jsdama-12306", image)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("CaptchaId:", id, "Captcha:", value)
	result, cookies, succeed, err := CheckCaptcha(client, CaptchaCheckUrl, nil, value)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("CheckCaptcha result:", result, "succeed:", succeed,"new cookie:",cookies)

	if !succeed {
		myCaptcha.Report("Jsdama-12306", id)
	}
}
