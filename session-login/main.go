package main

import (
	"net/http"
	"text/template"
	"time"

	uuid "github.com/satori/go.uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()

}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	showSessions()
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		ps := r.FormValue("password")

		//is there a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(ps))
		if err != nil {
			http.Error(w, "Username or password do not match", http.StatusForbidden)
			return
		}

		//create session
		sId := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
		http.SetCookie(w, c)

		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u user
	if r.Method == http.MethodPost {

		//get from form values
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		ro := r.FormValue("role")

		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "User name already taken", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		})

		dbSessions[sID.String()] = session{un, time.Now()}

		//store password in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		dbUsers[un] = user{un, bs, f, l, ro}
		//redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")

	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	//clean up dbSessions
	if time.Since(dbSessionsCleaned) > time.Second*30 {
		go cleanSessions()
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
