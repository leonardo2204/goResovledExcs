package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("nao existe raiz quadrada de numero negativo: %.2f",float64(err))
}

func Sqrt(x float64) (float64, error) {	
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	} else {
		return x * x, nil
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
