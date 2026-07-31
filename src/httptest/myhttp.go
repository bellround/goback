package httptest

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func RunHTTP() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
