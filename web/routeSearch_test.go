package web

import (
	"testing"
	//"log"
	//"flag"
	//"strings"
	//"time"
	"me.org/http"
	//"captcha"
	//
	//"os"
	//"fmt"

	"log"
)

func TestLeftTicketLog(t *testing.T) {
	/*Username := ""
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
	client.SetInsecureSkipVerify(true)

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

			cookies1, err := UserLoginSec(client, UserLoginSecUrl, nil)
			if err != nil {
				log.Fatalln(err)
			}

			log.Println("UserLoginSec result:", result, "new cookies:", cookies1)

			log.Println("client cookie:", client.Cookie())

			cookies2, err := InitMy12306(client, InitMy12306Url, nil)
			if err != nil {
				log.Fatalln(err)
			}

			log.Println("InitMy12306 result:", result, "new cookies:", cookies2)

			log.Println("client cookie:", client.Cookie())

			cookies3, err := LeftTicketInit(client, LeftTiceketInitUrl, nil)
			if err != nil {
				log.Fatalln(err)
			}

			log.Println("LeftTiceketInit result:", result, "new cookies:", cookies3)

			log.Println("client cookie:", client.Cookie())

			cookies4, err := GetPassCodeNew(client, GetPassCodeNewUrl, nil)
			if err != nil {
				log.Fatalln(err)
			}

			log.Println("GetPassCodeNew result:", result, "new cookies:", cookies4)

			log.Println("client cookie:", client.Cookie())

			//passengers, cookies, err := GetPassengers(client, PassengerInitUrl, nil)
			//
			//if err != nil {
			//	log.Fatalln(err)
			//}
			//
			//log.Println("PassengerInitFirst result:", result, "new cookies:", cookies)
			//
			//log.Println("client cookie:", client.Cookie())
			//
			//log.Println("GetPassengers", passengers)
			//
			//for _, passenger := range passengers {
			//	status, err := passenger.CheckStatus()
			//	log.Println(passenger, status, err)
			//}

			//sign, err := LeftTiceketLog(client, LeftTiceketLogUrl, nil, "2017-12-10", "SZQ", "WHN", "ADULT")
			fromStation := []string{0: "深圳", 1: "SZQ"}
			toStation := []string{0: "武汉", 1: "WHN"}
			sign, err := LeftTiceketLog(client, LeftTiceketLogUrl, nil, fromStation, toStation, "2017-12-10", "2017-12-10", "dc", "ADULT")
			if err != nil {
				log.Fatalln(err)
			}

			log.Println("LeftTiceketLog result:", sign)

			log.Println("client cookie:", client.Cookie())

			msgJson, err := LeftTiceketQuery(client, LeftTiceketQueryUrl, nil, fromStation, toStation, "2017-12-10", "2017-12-10", "dc", "ADULT")
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("LeftTiceketQuery result:", msgJson)
			log.Println("client cookie:", client.Cookie())

			f, err := os.Create(`I:\gospace\project\12306\data\msg.txt`)
			if err != nil {
				fmt.Println("Create err:", err)
				return
			}
			defer f.Close()
			f.WriteString(string(msgJson.Data.Result[0]))

			break
		}
	}*/

	//fromStation := []string{0: "深圳", 1: "SZQ"}
	//toStation := []string{0: "武汉", 1: "WHN"}

	/*tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	url := "https://kyfw.12306.cn/otn/leftTicket/query"
	realUrl := url + "?leftTicketDTO.train_date=" + "2017-12-10" + "&leftTicketDTO.from_station=" + fromStation[1] + "&leftTicketDTO.to_station=" + toStation[1] + "&purpose_codes=" + "ADULT"
	resp, err := client.Get(realUrl)
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("ReadAll err:", err)
		return
	}
	f, err := os.Create(`I:\gospace\project\12306\certificate\msg.txt`)
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()
	f.WriteString(string(body))
	fmt.Println(string(body))*/

}

func TestLeftTicketQuery(t *testing.T) {

	fromStation := []string{0: "深圳", 1: "SZQ"}
	toStation := []string{0: "武汉", 1: "WHN"}
	client := http.NewClient(nil, "", "", 0)
	client.SetInsecureSkipVerify(true)
	res, err := LeftTicketQuery(client, LeftTicketQueryUrl, nil, fromStation, toStation, "2017-12-10", "ADULT")

	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	log.Println(res.Data)
	log.Println(res.Data.Result)
}
