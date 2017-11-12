package regist

import (
	"fmt"
	"html/template"
	"net/http"

	"GoLang/HM3/store"
)

// RegistController 用于创建构造器
type RegistController struct {
}

// RegistAction 调用regist.html网页
func (this *RegistController) RegistAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/htmls/regist.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

// TestRegistAction 测试账号是否存在,若不存在,直接注册
func (this *RegistController) TestRegistAction(w http.ResponseWriter, r *http.Request) {
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
}
