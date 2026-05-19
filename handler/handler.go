package handler

import (
	"html/template"
	"net/http"
	"proj1/repository"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	repository.ShowAll(w)
}
func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/add.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		id := r.FormValue("id")
		name := r.FormValue("name")
		age := r.FormValue("age")
		sex := r.FormValue("sex")
		repository.AddUser(id, name, age, sex, w)
		return
	}
}
func FormHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)

}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func RegistrHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/registr.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		login := r.FormValue("login")
		password := r.FormValue("password")
		repository.Reg(login, password, w, r)

		return
	}
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		username := r.FormValue("username")
		password := r.FormValue("password")

		repository.Auth(username, password, w, r)
	}
}
