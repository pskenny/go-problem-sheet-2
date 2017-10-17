// Paul Kenny G00326057
// Exercise 2, Make the text a H1

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Write "Guessing game" to response writer
	fmt.Fprintf(w, "<h1>Guessing game</h1>\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
