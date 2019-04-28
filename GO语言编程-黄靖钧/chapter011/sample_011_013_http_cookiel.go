package main

// go语言中通过net/http包中的SetCookie来设置Cookie
// http.SetCookie(w ResponseWriter, cookie *Cookie)

// type Cookie struct {
// 	Name  string
// 	Value string
//
// 	Path       string    // optional
// 	Domain     string    // optional
// 	Expires    time.Time // optional
// 	RawExpires string    // for reading cookies only
//
// 	// MaxAge=0 means no 'Max-Age' attribute specified.
// 	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// 	// MaxAge>0 means Max-Age attribute present and given in seconds
// 	MaxAge   int
// 	Secure   bool
// 	HttpOnly bool
// 	SameSite SameSite
// 	Raw      string
// 	Unparsed []string // Raw text of unparsed attribute-value pairs
// }

expiration := time.Now()
expiration = expiration.AddDate(1, 0, 0)
cookie := http.Cookie{Name: "username", Value: "wanghaohao", expiration}
http.SetCookie(w, &Cookie)

cookie, _ := r.Cookie("username")
fmt.Fprint(w.cookie)

// 还有另一种读取方式：
for _, cookie := range r.Cookies() {
	fmt.Fprint(w, cookie.Name)
}

// 这个练习需要做。登录网页获取用户权限。