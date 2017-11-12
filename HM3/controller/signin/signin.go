package signin

import (
  "net/http"
  "html/template"
  "fmt"

  "GoLang/HM3/store"
)

// SigninController 用于创建构造器
type SigninController struct {
}

// SigninAction 调用signin.html网页
func (this *SigninController) SigninAction(w http.ResponseWriter, r *http.Request)  {
  t, err := template.ParseFiles("public/htmls/signin.html")
  if err != nil {
    panic(err)
  }
  t.Execute(w, nil)
}

// TestSigninAction 测试账号是否存在
func (this *SigninController) TestSigninAction(w http.ResponseWriter, r *http.Request)  {
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
