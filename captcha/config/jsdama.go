package config

import (
	"me.org/http"
	json "github.com/json-iterator/go"
	"encoding/base64"
	"errors"
)

/*
 * 联众答题
 *
 * 数据
 *
 * https://www.jsdama.com/
 */
type Jsdama struct {
	Name             string       `json:"-"`                          // 名称
	SoftwareId       uint         `json:"softwareId"`                 // 软件id
	SoftwareSecret   string       `json:"softwareSecret"`             // 软件秘钥
	Username         string       `json:"username"`                   // 用户名称
	Password         string       `json:"password"`                   // 用户密码
	CaptchaData      string       `json:"captchaData,omitempty"`      // 图片base64值
	CaptchaType      uint         `json:"captchaType,omitempty"`      // 验证码类型
	CaptchaMinLength uint         `json:"captchaMinLength,omitempty"` // 最小长度
	CaptchaMaxLength uint         `json:"captchaMaxLength,omitempty"` // 最大长度
	CaptchaId        string       `json:"captchaId,omitempty"`        // 验证码id
	client           *http.Client `json:"-"`                          // http客户端
}

/*
 * 联众答题
 *
 * 响应
 *
 * 文档 http://l-files.ufile.ucloud.com.cn/api/%E8%81%94%E4%BC%97%E7%AD%94%E9%A2%98API%E6%8E%A5%E5%8F%A3%E5%8D%8F%E8%AE%AE.pdf
 */
type Reply struct {
	Code    int       `json:"code"`    // 错误码
	Message string    `json:"message"` // 错误信息
	Data    ReplyData `json:"data"`    // 数据
}

/*
 * 联众答题
 *
 * 响应数据
 *
 */
type ReplyData struct {
	Result bool `json:"result"` // 结果
	Recognition                 // 验证码结果
	Points                      // 我的点数
}

/*
 * 联众答题
 *
 * 验证码结果
 *
 */
type Recognition struct {
	CaptchaId   string `json:"captchaId"`   // 验证码id
	Recognition string `json:"recognition"` // 识别结果
}

/*
 * 联众答题
 *
 * 我的点数
 *
 */
type Points struct {
	UserPoints      int `json:"userPoints"`      // 剩余总点数
	AvailablePoints int `json:"availablePoints"` // 可用点数
	LockPoints      int `json:"lockPoints"`      // 锁定点数
}

/*
 * New
 *
 * 新建
 *
 * @param name 		string	名称
 * @param id 		uint	软件id
 * @param secret 	string	软件秘钥
 * @param username 	string	联众用户名
 * @param password 	string	联众密码
 *
 * @return *Jsdama
 *
 */
func New(name string, id uint, secret, username, password string) *Jsdama {

	client := http.NewClient(nil, "", "", 30)

	return &Jsdama{
		Name:           name,
		SoftwareId:     id,
		SoftwareSecret: secret,
		Username:       username,
		Password:       password,
		client:         client,
	}
}

/*
 * Upload
 *
 * 上传验证码图片进行验证
 *
 * @param image []byte		图片字节流
 *
 * @return *Recognition 	验证码结果
 * @return error 			错误
 *
 */
func (this *Jsdama) Upload(captchaType uint, image []byte, min uint, max uint) (*Recognition, error) {
	url := "https://v2-api.config.com/upload"
	// 复制对象
	jsdama := *this
	// 设置值
	jsdama.CaptchaType = captchaType
	jsdama.CaptchaMinLength = min
	jsdama.CaptchaMaxLength = max
	jsdama.CaptchaData = base64.StdEncoding.EncodeToString(image)
	// 转换为json数据
	data, err := json.Marshal(jsdama)
	if err != nil {
		return nil, err
	}
	// 异步Post
	resp, err := this.client.AsyncPostJson(url, data)
	if err != nil {
		return nil, err
	}
	// 解析返回值
	reply := &Reply{}
	err = json.Unmarshal(resp.Body, reply)
	if err != nil {
		return nil, err
	}
	if reply.Code != 0 {
		return nil, errors.New(reply.Message)
	}
	// 获取值
	recognition := &Recognition{
		CaptchaId:   reply.Data.Recognition.CaptchaId,
		Recognition: reply.Data.Recognition.Recognition,
	}
	return recognition, nil
}

/*
 * Report
 *
 * 上传验证码图片进行验证
 *
 * @param captchaId string	验证码id
 *
 * @return error 	错误
 *
 */
func (this *Jsdama) Report(captchaId string) error {
	url := "https://v2-api.config.com/report-error"
	// 复制结构体
	jsdama := *this
	// 设置值
	jsdama.CaptchaId = captchaId
	// 转换为json数据
	data, err := json.Marshal(jsdama)
	if err != nil {
		return err
	}
	// 异步Post
	resp, err := this.client.AsyncPostJson(url, data)
	if err != nil {
		return nil
	}
	// 解析返回值
	reply := &Reply{}
	err = json.Unmarshal(resp.Body, reply)
	if err != nil {
		return err
	}
	if reply.Code != 0 {
		return errors.New(reply.Message)
	}
	return nil
}

/*
 * CheckPoints
 *
 * 获取我的点数
 *
 * @return *Points 	我的点数
 * @return error 	错误
 *
 */
func (this *Jsdama) CheckPoints() (*Points, error) {
	url := "https://v2-api.config.com/check-points"
	// 转换为json数据
	data, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	// 异步Post
	resp, err := this.client.AsyncPostJson(url, data)
	if err != nil {
		return nil, nil
	}
	// 解析返回值
	reply := &Reply{}
	err = json.Unmarshal(resp.Body, reply)
	if err != nil {
		return nil, nil
	}
	if reply.Code != 0 {
		return nil, errors.New(reply.Message)
	}
	// 获取值
	points := &Points{
		UserPoints:      reply.Data.Points.UserPoints,
		AvailablePoints: reply.Data.Points.AvailablePoints,
		LockPoints:      reply.Data.Points.LockPoints,
	}
	return points, nil
}
