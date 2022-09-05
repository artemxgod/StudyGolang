package main

import (
	"fmt"
	"net/http"
	"time"
)

// simple login page
var loginFormTmpl = `
<html>
	<body>
	<form action="/get_cookie" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	</body>
</html>
`

func hande_main(w http.ResponseWriter, r *http.Request) {
	// taking session_id from cookie
	sessionID, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		// if no cookie found user should login
		w.Write([]byte(loginFormTmpl))
		return
	} else if err != nil {
		panic(err)
	}
	fmt.Fprint(w, "Welcome, "+sessionID.Value)
}

func handle_cookie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// reading login from html form
	inputLogin := r.Form["login"][0]
	// in 1 year we will have to login again (new session)
	expire := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie {
		Name: "session_id",
		Value: inputLogin,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	// handler
	http.HandleFunc("/", hande_main)
	
	//create cookie
	http.HandleFunc("/get_cookie", handle_cookie)


	http.ListenAndServe(":8081", nil)
}

// cookies should be in headers, before main content

