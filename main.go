package main

import (
	_ "embed"
	"net/http"

	"github.com/syumai/workers"
)

//go:embed public/index.html
var htmlPage string

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(htmlPage))
	})
	workers.Serve(nil)
}
