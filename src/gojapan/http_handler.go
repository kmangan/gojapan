package gojapan

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	japaneseResult := translate(r.URL.Path[1:])
	fmt.Fprintf(w, "Translated %s: %s!", r.URL.Path[1:], japaneseResult)
}

func SetupHttpListener() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}