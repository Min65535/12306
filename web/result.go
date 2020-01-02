package web

import (
	json "github.com/json-iterator/go"
	"strconv"
	"errors"
)

/*
 * 12306返回结果
 */
type Result struct {
	ResultCode    int    `json:"result_code"`
	ResultMessage string `json:"result_message"`
	Uamtk         string `json:"uamtk,omitempty"`
	AppTK         string `json:"apptk,omitempty"`
	NewAppTK      string `json:"newapptk,omitempty"`
	Username      string `json:"username,omitempty"`
}

type ResponseJson struct {
	ValidateMessagesShowId string      `json:"validateMessagesShowId"`
	Status                 bool        `json:"status"`
	HttpStatus             int         `json:"httpstatus"`
	Messages               []string    `json:"messages"`
	ValidateMessages       interface{} `json:"validateMessages"`
}

type DataJson struct {
	Result []string    `json:"result"`
	Flag   string      `json:"flag"`
	Map    interface{} `json:"map"`
}

type MsgJson struct {
	ValidateMessagesShowId string      `json:"validateMessagesShowId"`
	Status                 bool        `json:"status"`
	HttpStatus             int         `json:"httpstatus"`
	Data                   DataJson    `json:"data"`
	Messages               []string    `json:"messages"`
	ValidateMessages       interface{} `json:"validateMessages"`
}

type CommonJson struct {
	ValidateMessagesShowId string      `json:"validateMessagesShowId"`
	Status                 bool        `json:"status"`
	HttpStatus             int         `json:"httpstatus"`
	Data                   interface{} `json:"data"`
	Messages               []string    `json:"messages"`
	ValidateMessages       interface{} `json:"validateMessages"`
}

/*
 * GetTK
 *
 * 获取tk值
 *
 * 对应官网的checkUAM函数
 *
 * @return string
 *
 */
func (this *Result) GetTK() string {
	if this.NewAppTK != "" {
		return this.NewAppTK
	}
	return this.AppTK
}

/*
 * UnmarshalJSON
 *
 * JSON反序列化
 *
 * @param data []byte
 *
 * @return error
 *
 */
func (this *Result) UnmarshalJSON(data []byte) error {
	temp := make(map[string]interface{})
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	for key, value := range temp {
		if key == "result_message" {
			this.ResultMessage = value.(string)
		} else if key == "result_code" {
			switch value.(type) {
			case string:
				num, err := strconv.Atoi(value.(string))
				if err != nil {
					return err
				}
				this.ResultCode = num
			case float64:
				this.ResultCode = int(value.(float64))
			default:
				return errors.New("unsupport type")
			}
		} else if key == "uamtk" {
			this.Uamtk = value.(string)
		} else if key == "apptk" {
			if value != nil {
				this.AppTK = value.(string)
			}
		} else if key == "newapptk" {
			if value != nil {
				this.NewAppTK = value.(string)
			}
		} else if key == "username" {
			this.Username = value.(string)
		}
	}
	return nil
}

// 验证码
// 4 验证码校验成功
// 5 验证码校验失败
// 7 验证码已经过期

// Uamtk
// 0 验证通过
// 1 用户未登录

// 登录
// 0 登录成功
// 1 登录名不存在。
// 1 密码输入错误。如果输错次数超过4次，用户将被锁定。
