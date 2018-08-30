package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strings"
	"time"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func main() {
	addr := ":3001"
	fmt.Printf("Servidor iniciado en el puerto %v\n", addr)
	http.ListenAndServe(addr, router())
}

func router() http.Handler {
	r := chi.NewRouter()

	// Rutas Protegidas
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())

			var exp int64
			//uffff.. Golang y sus tipos de datos
			if expv, ok := claims["exp"]; ok {
				switch v := expv.(type) {
				case float64:
					exp = int64(v)
				case int64:
					exp = v
				case json.Number:
					exp, _ = v.Int64()
				default:
				}
			}
			w.Write([]byte(fmt.Sprintf("Area protegida por JWT. Bienvenido: %v\nEste Token expira el: %v\n", claims["user_id"], time.Unix(exp, 0))))
		})
	})

	// Rutas Publicas
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Bienvenido Anonimo, Nada por aqui!\n"))
		})
		r.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
			auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

			if len(auth) != 2 || auth[0] != "Basic" {
				http.Error(w, "Falló la Autorización\n", http.StatusUnauthorized)
				return
			}

			payload, _ := base64.StdEncoding.DecodeString(auth[1])
			pair := strings.SplitN(string(payload), ":", 2)

			if len(pair) != 2 || !validate(pair[0], pair[1]) {
				http.Error(w, "Falló la Autorización: credenciales inválidas\n", http.StatusUnauthorized)
				return
			} else {
				//el JWT tendrá 180 segundos de vida
				expiration := int64(time.Now().Unix()) + 180
				w.Write([]byte(fmt.Sprintf("Este es tu JWT para esta sesión: %v\nExpira el: %s\n\n", generate_jwt(pair[0], expiration), time.Unix(expiration, 0))))
			}

		})
	})

	return r
}

//Validacion de usuarios
func validate(username, password string) bool {
	out := false
	database, _ := sql.Open("sqlite3", "./usuarios.db")
	err := database.QueryRow("select username, password from usuarios where username LIKE ? and password LIKE ?", username, password).Scan(&username)
	if err != nil && err == sql.ErrNoRows {
		out = false
	} else {
		//el usuario existe en la DB. lo valido
		out = true
	}
	return out
}

//Generacion de Jwt
func generate_jwt(username string, expiration int64) string {
	//el claim "exp" es un unix timestamp que especifica el momento de expiracion del JWT. La rfc 7519 sugiere usar unos  pocos minutos. Yo le puse 2
	_, tokenString, _ := tokenAuth.Encode(jwtauth.Claims{"user_id": username, "exp": expiration})
	fmt.Printf("DEBUG: JWT: %s\n Claim: \"user_id\": %s\n\"exp\": %s", tokenString, username, time.Unix(expiration, 0))
	return tokenString
}
