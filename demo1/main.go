package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	. "webdemo/demo1/server"
)

func main() {

	go post("post-json")
	WebServerBase()
}

func get() {

	time.Sleep(1 * time.Second)

	response, err := http.Get("http://127.0.0.1:8081/login?username=zhangsan&password=123456")
	if err != nil {

		fmt.Println("err happens ", err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("in here , ", string(body))

	if response.StatusCode == 200 {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}
}
func post(t string) {
	time.Sleep(1 * time.Second)

	switch t {
	case "post":

		data := url.Values{}
		data.Set("firstname", "foo")
		data.Add("lastname", "bar")
		//两种情况的post ..
		//1.普通的post表单请求，Content-Type=application/x-www-form-urlencoded
		//2.有文件上传的表单，Content-Type=multipart/form-data
		resp, err := http.Post("http://localhost:8081/login", "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("POST OK: ", string(body))
		}

	case "postform":

		resp, err := http.PostForm("http://127.0.0.1:8081/login",
			url.Values{"firstname": {"ruifengyun"}, "lastname": {"johnnythree"}})
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("POST OK: ", string(body))
		}

	case "post-json":

		/*
		 [123 10 9 9 9 9 9 9 9 34 85 115 101 114 110 97 109 101 34 58 32 34 116 105 97
		 110 120 105 97 102 101 110 103 121 117 110 34 44 10 9 9 9 9 9 9 9 34 66 108 111
		  103 34 58 32 34 106 111 104 110 110 121 116 104 114 101 101 46 99 99 34 44 10
		  9 9 9 9 9 9 9 34 72 111 98 98 121 115 34 58 32 34 99 43 43 32 103 111 108 97
		  110 103 34 44 10 9 9 9 9 9 9 9 34 67 111 110 116 101 110 116 34 58 32 34 119 104
		   97 116 39 115 32 121 111 117 114 32 104 111 98 98 121 63 34 125]
		*/

		body := bytes.NewBuffer([]byte(`{
							"Username": "tianxiafengyun",
							"Blog": "johnnythree.cc",
							"Hobbys": "c++ golang",
							"Content": "what's your hobby?"}`))

		resp, err := http.Post("http://127.0.0.1:8081/login", "application/json", body)

		defer resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("POST OK: ", string(body))
		}

	}

}
