package captcha

/*
 * ICaptcha接口
 */
type ICaptcha interface {
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
	Upload([]byte) (string, string, error)
	/*
	 * Report
	 *
	 * 上报错误的验证结果
	 *
	 * @param string 	验证标识
	 *
	 * @return error 	错误
	 *
	 */
	Report(string) error
}