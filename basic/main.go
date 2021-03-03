package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// computeE computes the approximation of e by running a fixed number of iterations.
func computeE(iterations int64) float64 {
	res := 2.0
	fact := 1.0

	for i := int64(2); i < iterations; i++ {
		fact *= float64(i)
		res += 1 / fact
	}
	return res
}

func main() {
	http.HandleFunc("/e", func(w http.ResponseWriter, r *http.Request) {

		v := r.FormValue("iters")
		s, _ := strconv.Atoi(v)

		w.Write([]byte(fmt.Sprintf("e = %0.4f\n", computeE(int64(s)))))
	})

	http.ListenAndServe(":8080", nil)

}
