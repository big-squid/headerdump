package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Printf("value of environment variable %v: %v\n", "ECHO_VAR", os.Getenv("ECHO_VAR"))
		r.Write(w)
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
