// Paul Kenny G00326057
// Exercise 1, Guessing game

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Write "Guessing game" to response writer
	fmt.Fprintf(w, "Guessing game\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
