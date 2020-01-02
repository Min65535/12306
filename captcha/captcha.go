package captcha

/*
 * Captcha
 */
type Captcha struct {
	instances // 实现的接口集
}

/*
 * 默认Captcha
 */
var defaultCaptcha *Captcha

/*
 * GetCaptcha
 *
 * 获取Captcha实例
 *
 * @return *Captcha
 *
 */
func GetCaptcha() *Captcha {
	if defaultCaptcha == nil {
		defaultCaptcha = &Captcha{}
	}
	return defaultCaptcha
}

/*
 * Register
 *
 * 注册实现实例
 *
 * @return bool true - 成功 | false - 失败
 *
 */
func (this *Captcha) Register(name string, impl ICaptcha) bool {
	var ret bool = false
	this.instances, ret = this.append(instance{name: name, impl: impl})
	return ret
}

/*
 * GetRegisterNames
 *
 * 获取所有注册实现实例的名称
 *
 * @param name string 	实现的接口名称
 *
 * @return []string
 *
 */
func (this *Captcha) GetRegisterNames() []string {
	names := []string{}
	for _, item := range this.instances {
		names = append(names, item.name)
	}
	return names
}

/*
 * GetCaptchaImpl
 *
 * 注册实现实例
 *
 * @param name string 	实现的接口名称
 *
 * @return ICaptcha 	实现的接口
 *
 */
func (this *Captcha) GetCaptchaImpl(name string) (ICaptcha) {
	return this.get(name)
}

/*
 * Upload
 *
 * 上传验证码图片进行验证
 *
 *
 * @param name string 	实现的接口名称
 * @param image []byte	图片字节流
 *
 * @return string 	验证标识
 * @return string 	验证结果
 * @return error 	错误
 *
 */
func (this *Captcha) Upload(name string, image []byte) (string, string, error) {
	impl := this.get(name)
	return impl.Upload(image)
}

/*
 * Report
 *
 * 上报错误的验证结果
 *
 * @param name string 	实现的接口名称
 * @param id string 	验证标识
 *
 * @return error 		错误
 *
 */
func (this *Captcha) Report(name string, id string) (error) {
	impl := this.get(name)
	return impl.Report(id)
}

/*
 * 单个Captcha实例
 */
type instance struct {
	name string   // 名称
	impl ICaptcha // 实现
}

/*
 * Captcha实例集
 */
type instances []instance

/*
 * append
 *
 * 增加
 *
 * @return instances
 * @return bool
 *
 */
func (this instances) append(ins instance) (instances, bool) {
	for _, item := range this {
		if item.name == ins.name {
			return this, false
		}
	}
	this = append(this, ins)
	return this, true
}

/*
 * get
 *
 * 获取
 *
 * @param name string 	实现的接口名称
 *
 * @return ICaptcha 	实现的接口
 *
 */
func (this instances) get(name string) ICaptcha {
	for _, item := range this {
		if item.name == name {
			return item.impl
		}
	}
	return nil
}
