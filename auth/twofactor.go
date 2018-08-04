package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	name     string
	username string
	password string
	tfa      string
}

type Session struct {
	id       string
	username string
	ip       string
}

var sessions []*Session
var users []*User

var registerTemplate = template.Must(template.New("register").Parse(`
	<html>
	<head><title>Register</title></head>
	<body>
		{{with .Message}}<p>{{.Message}}</p>{{end}}
		<form><input type="text" prompt="name" name="name"/><br/>
		<input type="text" prompt="username" name="username"/><br/>
		<input type="password" prompt="password" name="password"/><br/>
		<input type="submit"/></form>
	</body>
	</html>`))
var loginTemplate = template.Must(template.New("login").Parse(`
	<html><head><title>Login</title></head>
	<body>
		<form method="post" action="/login">
		<input type="text" name="username"/>
		<input type="text" name="password"/>
		<input type="submit"/>
		</form>
	</body></html>`))
var tfaTemplate = ``
var welcomeTemplate = `<html><head><title>Welcome></title></head><body><p>Welcome</p></body></html>`

func main() {
	fmt.Println("Hello There... Starting up")

	du := new(User)
	du.name = "Default User"
	du.username = "user"
	du.password = "pass"
	users = append(users, du)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			u := new(User)
			u.name = r.PostFormValue("name")
			u.username = r.PostFormValue("username")
			u.password = r.PostFormValue("password")
			//TODO TFA Something

			if !u.IsValid() {
				registerTemplate.Execute(w, nil)
				//http.Error(w, u.Validate(), http.StatusBadRequest)
				return
			}

			//Check not exists
			for v := range users {
				if v.username == u.username {
					registerTemplate.Execute(w, nil)
					//http.Error(w, "username not unique", http.StatusBadRequest)
					return
				}
			}
			//Create a session
			s := new(Session)
			s.id = GenUUIDv4()
			s.username = fu.username
			s.ip = r.RemoteAddr
			//Set a cookie
			c := new(http.Cookie)
			c.Name = "xs"
			c.Value = s.id
			c.HttpOnly = true
			c.MaxAge = 60 * 60 * 4
			http.SetCookie(w, c)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		registerTemplate.Execute(w, nil)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		redir := r.FormValue("link")

		found := false
		var fu *User
		for _, u := range users {
			if username == u.username && password == u.password {
				//Create a session for this user
				found = true
				fu = u
				break
			}
		}
		if !found {
			loginTemplate.Execute(w, nil)
			return
		}

		//Create a session
		s := new(Session)
		s.id = GenUUIDv4()
		s.username = fu.username
		s.ip = r.RemoteAddr
		//Set a cookie
		c := new(http.Cookie)
		c.Name = "xs"
		c.Value = s.id
		c.HttpOnly = true
		c.MaxAge = 60 * 60 * 4
		http.SetCookie(w, c)
		//Ask for tfa if conditions are met
		//Redirect to where the user should go
		http.Redirect(w, r, redir, http.StatusFound)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := r.Cookie("xs")
		loginTemplate.Execute(w, nil)
		//w.Write("Welcome to the Site")
	})

	http.ListenAndServe(":8080", nil)
}

func GenTwoFactorAuth() {
}

func ValidateTwoFactorAuth() {
}

func GenUUIDv4() string {
	u := make([]byte, 16)
	rand.Read(u)
	//Set the version to 4
	u[6] = (u[6] | 0x40) & 0x4F
	u[8] = (u[8] | 0x80) & 0xBF
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}
