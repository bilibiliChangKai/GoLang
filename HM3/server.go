package main

import (
	"net/http"
	"reflect"
	"strings"

	"GoLang/HM3/controller/regist"
	"GoLang/HM3/controller/signin"
	"GoLang/HM3/controller/view"

	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()

	m.Get("/regist", registRouter)
	m.Get("/signin", signinRouter)
	m.Post("/sign_in_succeed", viewRouter)

	m.Group("/test", func(r martini.Router) {
		m.Post("/regist", registRouter)
		m.Post("/signin", signinRouter)
	})

	// 在8000端口运行服务器
	m.RunOnAddr(":8000")
}

func registRouter(w http.ResponseWriter, r *http.Request) {
	var methodName string
	if pathInfo := strings.Trim(r.URL.Path, "/"); pathInfo == "test/regist" {
		methodName = "TestRegist"
	} else {
		methodName = "Regist"
	}

	// 创建控制器,运行方法
	ob := &regist.RegistController{}
	callFunction(ob, methodName, &w, r)
}

func signinRouter(w http.ResponseWriter, r *http.Request) {
	var methodName string
	if pathInfo := strings.Trim(r.URL.Path, "/"); pathInfo == "test/signin" {
		methodName = "TestSignin"
	} else {
		methodName = "Signin"
	}

	// 创建控制器,运行方法
	ob := &signin.SigninController{}
	callFunction(ob, methodName, &w, r)
}

func viewRouter(w http.ResponseWriter, r *http.Request) {
	// 创建控制器
	ob := &view.ViewController{}
	callFunction(ob, "View", &w, r)
}

// 通过结构变量和函数名运行方法
func callFunction(i interface{}, methodName string, w *http.ResponseWriter, r *http.Request) {
	// 创建控制器,得到方法
	controller := reflect.ValueOf(i)
	method := controller.MethodByName(methodName + "Action")
	// 通过反射传递方法值
	req := reflect.ValueOf(r)
	res := reflect.ValueOf(*w)
	// 运行方法
	method.Call([]reflect.Value{res, req})
}
