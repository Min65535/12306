package captcha

import (
	"testing"
	"log"
)

func TestCaptcha_GetRegisterName(t *testing.T) {
	myCaptcha := GetCaptcha()
	log.Println(myCaptcha.GetRegisterNames())
}
