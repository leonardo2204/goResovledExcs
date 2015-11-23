package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	iteracao := 0
	prox := 0
	primeiro := 0
	segundo := 1

	return func() int {
		if iteracao <= 1 {
			prox = iteracao
			iteracao += 1
		} else {
			prox = primeiro + segundo
			primeiro = segundo
			segundo = prox
		}
		return prox
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
