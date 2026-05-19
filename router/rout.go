package rout

import (
	"fmt"
	"net/http"
	sozd "proj1/func"
	"text/template"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	sozd.ShowAll(w)
}
func addHandler(w http.ResponseWriter, r *http.Request) {
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
		sozd.AddUser(id, name, age, sex, w)
		return
	}
}
func formHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)

}
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func registrHandler(w http.ResponseWriter, r *http.Request) {

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
		sozd.Reg(login, password, w, r)

		return
	}
}
func loginHandler(w http.ResponseWriter, r *http.Request) {

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

		sozd.Auth(username, password, w, r)
	}
}

func Rout() {
	http.HandleFunc("/", sozd.AuthMiddleware(formHandler))
	http.HandleFunc("/add", sozd.AuthMiddleware(addHandler))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/users", sozd.AuthMiddleware(usersHandler))
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/registr", registrHandler)
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
