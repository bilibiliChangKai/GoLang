package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"GoLang/HM3/store"

	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "hello"
	})
	m.Get("/123.txt", func() string {
		return "trytry"
	})

	m.Get("/regist", registControllor)
	m.Get("/signin", signinControllor)
	m.Post("/sign_in_succeed", viewControllor)

	//m.Post("/test/regist", testRegistControllor)
	m.Group("/test", func(r martini.Router) {
		m.Post("/regist", testRegistControllor)
		m.Post("/signin", testSigninControllor)
	})
	m.RunOnAddr(":8000")
}

func registControllor(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/htmls/regist.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func testRegistControllor(w http.ResponseWriter, r *http.Request) {
	// 解析表单
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	name := r.FormValue("name")

	// 设置头
	w.Header().Set("content-type", "text/plain")
	// 允许跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 判断是否注册过
	if ok := store.IsExistedUser(name); ok {
		fmt.Fprint(w, "True")
	} else {
		// 如果没有,则进行注册
		id, password, phonenumber, email := r.FormValue("id"), r.FormValue("password"), r.FormValue("phone"), r.FormValue("email")
		store.AddUser(*store.NewUser(id, name, password, email, phonenumber))
		fmt.Fprint(w, "False")
	}
	//fmt.Println("enter?")
}

func testSigninControllor(w http.ResponseWriter, r *http.Request) {
	// 解析表单
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	name := r.FormValue("name")

	// 设置头
	w.Header().Set("content-type", "text/plain")
	// 允许跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 判断是否注册过
	if ok := store.IsExistedUser(name); ok {
		fmt.Fprint(w, "True")
	} else {
		// 不进行注册
		fmt.Fprint(w, "False")
	}
}

func viewControllor(w http.ResponseWriter, r *http.Request) {
	var name string
	// post 请求,设置cookie
	if err := r.ParseForm(); err == nil {
		name = r.FormValue("name")
		cookie := http.Cookie{Name: "name", Value: name, Path: "/"}
		http.SetCookie(w, &cookie)
	}

	// get 请求,获得cookie
	coo := r.Header["cookie"]
	fmt.Println(coo)
	c, err := r.Cookie("name")
	if err != nil {
		panic(err)
	}
	name = c.Value

	// 获得user
	//ur, _ := store.GetUser(name)
	ur := store.UserItem{"123", "456", "789", "101", "112"}

	//t, err := template.ParseFiles("public/htmls/view.html")
	b, _ := ioutil.ReadFile("public/htmls/view.html")
	// 切割并改变tml内容
	s := string(b)
	s = strings.Replace(s, "#{target}", ur.Name, 1)
	s = strings.Replace(s, "#{target}", ur.ID, 1)
	s = strings.Replace(s, "#{target}", ur.PhoneNumber, 1)
	s = strings.Replace(s, "#{target}", ur.Email, 1)
	fmt.Println(s)
	w.Header().Set("content-type", "text/html")
	fmt.Fprint(w, s)
}

func signinControllor(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/htmls/signin.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}
