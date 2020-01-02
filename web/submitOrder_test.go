package web

import (
	"testing"
	"time"
	"fmt"
	"strconv"
)

func TestQueryOrderWaitTime(t *testing.T) {
	//dataTime := time.Now().Unix()*1000
	//dataTime1 := time.Now().UnixNano()
	//fmt.Println(dataTime)
	//fmt.Println(dataTime1)
	//

	fmt.Println(time.Now().Unix())               //获取当前秒
	fmt.Println(time.Now().UnixNano())           //获取当前纳秒
	fmt.Println(time.Now().UnixNano() / 1e6)     //将纳秒转换为毫秒
	fmt.Println(time.Now().UnixNano() / 1e9)     //将纳秒转换为秒
	c := time.Unix(time.Now().UnixNano()/1e9, 0) //将毫秒转换为 time 类型
	fmt.Println(c.String())                      //输出当前英文时间戳格式

	milliSecondStr := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	fmt.Printf("%t\n", milliSecondStr)


}
