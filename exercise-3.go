// Paul Kenny G00326057
// Exercise 3, Serve a page using Bootstrap
// See https://golang.org/doc/articles/wiki/, code snippets used and adapted below.
// This exercise requires bootstrap.txt

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Load page from file, given title of file
func loadPage(title string) ([]byte, error) {
	filename := title + ".txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte("Error reading file"), err
	}
	return contents, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	content, _ := loadPage("bootstrap")
	fmt.Fprintf(w, "%s\n", content)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
