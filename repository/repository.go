package repository

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	models "proj1/User"
	"strconv"
	"text/template"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func Loading() {
	godotenv.Load()
	connStr = os.Getenv("DB_CONN_STR")
	jwtKey = []byte(os.Getenv("JWT_KEY"))
}

var jwtKey = []byte(os.Getenv("JWT_KEY"))
var connStr = os.Getenv("DB_CONN_STR")

func ShowAll(w http.ResponseWriter) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	res, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}
	defer res.Close()
	users := []models.User{}

	for res.Next() {

		u := models.User{}

		err := res.Scan(&u.Id, &u.Name, &u.Age, &u.Sex)
		if err != nil {
			panic(err)
		}

		users = append(users, u)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}
func AddUser(id, name, age, sex string, w http.ResponseWriter) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	intid, _ := strconv.Atoi(id)
	defer db.Close()
	var exists bool
	err = db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)",
		id,
	).Scan(&exists)

	if err != nil {
		panic(err)
	}
	tmpl, _ := template.ParseFiles("templates/add.html")
	if exists {

		tmpl.Execute(w, map[string]interface{}{
			"Error": "Пользователь уже существует",
		})
		return
	}
	intage, _ := strconv.Atoi(age)
	_, err = db.Exec("INSERT INTO users(id,name,age,sex)VALUES($1,$2,$3,$4)", intid, name, intage, sex)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, map[string]string{
		"Success": "Регистрация успешна",
	})
}
func Reg(log, pass string, w http.ResponseWriter, r *http.Request) {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(pass),
		bcrypt.DefaultCost,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec(
		"INSERT INTO admin(login,pass) VALUES($1, $2)",
		log,
		string(hash),
	)
	tmpl, _ := template.ParseFiles("templates/registr.html")
	if err != nil {
		tmpl.Execute(w, map[string]string{
			"Error": "Пользователь уже существует",
		})
		return
	}

	tmpl.Execute(w, map[string]string{
		"Success": "Регистрация успешна",
	})
}
func Auth(name, pass string, w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var password string

	err = db.QueryRow(
		"SELECT pass FROM admin WHERE login=$1",
		name,
	).Scan(&password)

	if err == sql.ErrNoRows {
		tmpl, _ := template.ParseFiles("templates/login.html")

		tmpl.Execute(w, map[string]interface{}{
			"Error": " неверный логин или пароль",
		})
		return
	}

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(password),
		[]byte(pass),
	)
	if err == nil {

		claims := jwt.MapClaims{
			"username": name,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			HttpOnly: true,
			Path:     "/",
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		tmpl, _ := template.ParseFiles("templates/login.html")

		tmpl.Execute(w, map[string]interface{}{
			"Error": " неверный логин или пароль",
		})

	}
}
