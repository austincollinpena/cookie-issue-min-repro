package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"net/http"
	"time"
)

const defaultPort = ":8089"

//// Middleware decodes the share session cookie and packs the session into context
func AddUserIDToCTX() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			expire := time.Now().Add(time.Minute * 30)

			cookie := http.Cookie{
				Name:    "hI",
				Value:   "there",
				Expires: expire,
				Domain: "app.localhost",
			}
			http.SetCookie(w, &cookie)
			next.ServeHTTP(w, r)
		})
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	expire := time.Now().Add(time.Minute * 30)

	cookie := http.Cookie{
		Name:    "hI",
		Value:   "there",
		Expires: expire,
		Domain: "localhost",
	}

	http.SetCookie(w, &cookie)
	w.Write([]byte("Hello"))
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://app.localhost", "http://api.localhost"},
		AllowCredentials: true,
		Debug:            true,
		AllowedHeaders:   []string{"Content-Type", "Sentry-Trace"},
		MaxAge:           300,
	}).Handler)
	r.Use(AddUserIDToCTX())


	r.HandleFunc("/", HandleHome)
	panic(http.ListenAndServe("0.0.0.0:8089", r))
}

