package config

import (
	"strings"
)

/*
 * 联众答题
 *
 * 12306实例
 *
 */
type Jsdama12306 struct {
	*Jsdama
}

/*
 * NewJsdama12306
 *
 * 新建12306接口
 *
 * @param id 		uint	软件id
 * @param secret 	string	软件秘钥
 * @param username 	string	联众用户名
 * @param password 	string	联众密码
 *
 * @return *Jsdama12306
 *
 */
func NewJsdama12306(id uint, secret, username, password string) *Jsdama12306 {
	jsdama12306 := &Jsdama12306{
		Jsdama: New("Jsdama-12306", id, secret, username, password),
	}
	return jsdama12306
}

/*
 * Upload
 *
 * 上传验证码图片进行验证
 *
 * @param []byte 	图片字节流
 *
 * @return string 	验证标识
 * @return string 	验证结果
 * @return error 	错误
 *
 */
func (this *Jsdama12306) Upload(image []byte) (string, string, error) {
	recognition, err := this.Jsdama.Upload(1321, image, 0, 0)
	if err != nil {
		return "", "", err
	}
	value := recognition.Recognition
	if strings.Contains(recognition.Recognition, "&") {
		value = strings.Replace(recognition.Recognition, "&", ",", -1)
	}
	return recognition.CaptchaId, value, nil
}
