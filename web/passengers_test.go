package web

import (
	"testing"
	"log"
	"flag"
	"strings"
	"time"
	"me.org/http"
	"12306/captcha"
)

func TestGetPassengers(t *testing.T) {
	Username := ""
	Password := ""

	for _, arg := range flag.Args() {
		if strings.ContainsAny(arg, "=") {
			_arg := strings.Split(arg, "=")
			if _arg[0] == "username" {
				Username = _arg[1]
			} else if _arg[0] == "password" {
				Password = _arg[1]
			}
		}
	}
	if Username == "" || Password == "" {
		log.Fatalln("please input 12306 username and password in Run -> Edit Configurations -> [Go Test] -> TestLogin in login_test.go -> Program arguments set like username=用户名 password=用户密码")
	}

	log.Println("12306 username:", Username, "Password:", Password)

	client := http.NewClient(nil, "", "", 0)
	// 初始化
	cookies, err := Init(client, InitUrl)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Init new cookies:", cookies)
	// 验证
	result, cookies, err := Uamtk(client, UamtkUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Uamtk result:", result, "new cookies:", cookies)

	log.Println("client cookie:", client.Cookie())

	myCaptcha := captcha.GetCaptcha()

	for {

		// 获取验证码
		image, cookies, err := GetCaptcha(client, CaptchaUrl, nil)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("GetCaptcha image length:", len(image), "new cookies:", cookies)

		log.Println("client cookie:", client.Cookie())

		// 远程打码
		id, value, err := myCaptcha.Upload("Jsdama-12306", image)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("Upload CaptchaId:", id, "Captcha:", value)

		// 检查验证码
		result, cookies, succeed, err := CheckCaptcha(client, CaptchaCheckUrl, nil, value)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("CheckCaptcha result:", result, "succeed:", succeed, "new cookie:", cookies)

		if !succeed {
			myCaptcha.Report("Jsdama-12306", id)
			log.Println("sleep 2s")
			time.Sleep(2 * time.Second)
			continue
		}

		result, cookies, succeed, err = Login(client, LoginUrl, nil, Username, Password)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("Login result:", result, "succeed:", succeed, "new cookies:", cookies)

		log.Println("client cookie:", client.Cookie())

		if succeed {
			cookies, err = UserLogin(client, UserLoginUrl, nil)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("UserLogin new cookies:", cookies)
			log.Println("client cookie:", client.Cookie())
			result, cookies, err = UamtkLogin(client, UamtkUrl, nil)
			if err != nil {
				log.Fatalln(err)
			}

			log.Println("UamtkLogin result:", result, "new cookies:", cookies)

			log.Println("client cookie:", client.Cookie())

			result, cookies, err = UamauthClient(client, UamauthClientUrl, nil, result.GetTK())
			if err != nil {
				log.Fatalln(err)
			}

			log.Println("UamauthClient result:", result, "new cookies:", cookies)

			log.Println("client cookie:", client.Cookie())

			log.Println("welcome to", result.Username)

			passengers, cookies, err := GetPassengers(client, PassengerInitUrl, nil)

			if err != nil {
				log.Fatalln(err)
			}

			log.Println("PassengerInitFirst result:", result, "new cookies:", cookies)

			log.Println("client cookie:", client.Cookie())

			log.Println("GetPassengers", passengers)

			for _, passenger := range passengers {
				status, err := passenger.CheckStatus()
				log.Println(passenger, status, err)
			}
			break
		}
	}
}
