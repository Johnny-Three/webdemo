package server

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Postjson struct {
	Username string
	Blog     string
	Hobbys   string
	Content  string
}

func WebServerBase() {
	fmt.Println("This is webserver base!")

	//第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	http.HandleFunc("/login", loginTask)

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe("localhost:8081", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func loginTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("loginTask is running...")
	//获取客户端通过GET/POST方式传递的参数
	r.ParseForm()

	if r.Method == "GET" {

		param_userName, found1 := r.Form["username"]
		param_password, found2 := r.Form["password"]

		if !(found1 && found2) {
			fmt.Fprint(w, "Get 方法访问出错，请注意参数和URL拼写。。")
			return
		}

		result := NewBaseJsonBean()
		userName := param_userName[0]
		password := param_password[0]

		s := "userName:" + userName + ",password:" + password
		fmt.Println(s)

		if userName == "zhangsan" && password == "123456" {
			result.Code = 100
			result.Message = "login success .."
		} else {
			result.Code = 101
			result.Message = "wrong username or password .."
		}

		//return client the json data ..
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))

		//write to w , another two ways ..
		//fmt.Fprint(w, fmt.Sprintf("[%s] , %s ,%s,%s,%s!", req.Method, firstname, lastname, aihao, content))
		//w.Write([]byte(fmt.Sprintf("[%s] , %s ,%s", req.Method, firstname, lastname)))
	} else if r.Method == "POST" {

		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")

		fmt.Printf("parse post msg ... [%s %s]\n", firstname, lastname)
		w.Write([]byte(fmt.Sprintf("[%s] Hello, %s %s!", r.Method, firstname, lastname)))

		//test json style ..
		result, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Println("in r post ", err)
			return
		}
		fmt.Printf("parse post result is %q\n", result)

		//write to struct
		var pj Postjson
		json.Unmarshal([]byte(result), &pj)

		fmt.Println(pj.Username)
		fmt.Println(pj.Blog)
		fmt.Println(pj.Hobbys)
		fmt.Println(pj.Content)

		//handle unknown type ..
		var f interface{}
		err = json.Unmarshal(result, &f)
		if err != nil {
			fmt.Println(err)
			return
		}
		m, ok := f.(map[string]interface{})
		if ok == false {

			fmt.Println("m type convert fails ~! ")
		}

		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}

	}
}
