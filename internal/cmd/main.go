package main

import (
	"fmt"
	"net/http"
	"webchatproject/app"

	"github.com/julienschmidt/httprouter"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "True")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	_ = app.NewDB()
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Test hello")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: CORS(router),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
