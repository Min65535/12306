package web

import (
	"testing"
	"log"
)

func TestResult_UnmarshalJSON(t *testing.T) {
	datas := []string{
		`{"result_message":"验证码校验成功","result_code":"4"}`,
		`{"result_message":"验证码校验成功","result_code":4}`,
		`{"result_message":"验证通过","result_code":0,"apptk":null,"newapptk":"Ewpl0qViwc1d9z8yfemma8Zp2kBUVO8nqZesGlxGdt0lml2l0"}`,
	}

	result := &Result{}
	for _, data := range datas {
		err := result.UnmarshalJSON([]byte(data))

		if err != nil {
			log.Fatalln(err)
		}
		log.Println(result)
	}
}
