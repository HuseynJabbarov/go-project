package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	fmt.Println("Starting server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.RequestURI)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {

		queryParams := r.URL.Query()

		n1 := queryParams.Get("n1")
		n2 := queryParams.Get("n2")

		if !isDigit(n1) || !isDigit(n2) {
			http.Error(w, "Provide two digits to be added", http.StatusBadRequest)
		} else {
			num1, _ := strconv.Atoi(n1)
			num2, _ := strconv.Atoi(n2)
			num3 := num1 + num2
			//n3 := strconv.Itoa(num3)

			fmt.Fprintf(w, "Received digits: %s and %s , their sum is %v", n1, n2, num3)
		}

	})

	http.ListenAndServe(":8080", nil)

}

func isDigit(s string) bool {
	return '0' <= s[0] && s[0] <= '9'
}
