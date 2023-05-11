package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./dist"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			r.URL.Path = "/" // 重定向到 /
		}
		fileServer.ServeHTTP(w, r)
	})
	fmt.Println("[HTTP] start WEB server at 0.0.0.0:80")
	panic(http.ListenAndServe(":80", nil))
}
