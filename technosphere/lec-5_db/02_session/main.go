package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

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

var sessions = map[string]string{}

func handle_main(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		w.Write([]byte(loginFormTmpl))
		return
	} else if err != nil {
		panic(err)
	}

	username, ok := sessions[sessionID.Value]

	// check if session exists
	if !ok {
		fmt.Fprint(w, "SESSION NOT FOUND")
	} else {
		fmt.Fprint(w, "Welcome, "+ username)
	}
}

func handle_cookie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inputLogin := r.Form["login"][0]
	expire := time.Now().Add(24 * time.Hour)
	//  setting random session id so hacker cannot access my session 
	// using my username
	sessionsID := RandStringRunes(32)
	sessions[sessionsID] = inputLogin

	cookie := http.Cookie{
		Name: "session_id",
		Value: sessionsID,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	// main page handler
	http.HandleFunc("/", handle_main)
	
	// cookie handler
	http.HandleFunc("/get_cookie", handle_cookie)

	http.ListenAndServe(":8081", nil)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}