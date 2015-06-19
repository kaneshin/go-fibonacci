package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/kaneshin/go-fibonacci/fibonacci"
)

func main() {
	http.HandleFunc("/fibonacci", func(w http.ResponseWriter, r *http.Request) {
		num := r.URL.Query().Get("num")
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		val := fibonacci.Fibonacci(n)
		fmt.Fprint(w, fmt.Sprintf("%d\n", val))
	})
	http.HandleFunc("/", http.NotFound)
	http.ListenAndServe(":8888", nil)
}
