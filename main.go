package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("收到请求: %s", r.RequestURI)
		_, _ = io.WriteString(w, "Hello go!\n")
	})

	port := "8080"
	log.Printf("启动端口: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
