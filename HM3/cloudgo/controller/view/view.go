package view

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bilibiliChangKai/GoLang/HM3/cloudgo/store"
)

// Controller 用于创建构造器
type Controller struct {
}

// Action 用于创建View.html并修改内容
func (c *Controller) Action(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	// 获得user
	ur, _ := store.GetUser(r.FormValue("name"))

	b, _ := ioutil.ReadFile("public/htmls/view.html")
	// 切割并改变tml内容
	s := string(b)
	s = strings.Replace(s, "#{target}", ur.Name, 1)
	s = strings.Replace(s, "#{target}", ur.ID, 1)
	s = strings.Replace(s, "#{target}", ur.PhoneNumber, 1)
	s = strings.Replace(s, "#{target}", ur.Email, 1)
	w.Header().Set("content-type", "text/html")
	fmt.Fprint(w, s)
}
