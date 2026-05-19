package sozd

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("JWT_KEY")

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {

			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		tokenStr := cookie.Value

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

			return jwtKey, nil
		})

		if err != nil || !token.Valid {

			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}
