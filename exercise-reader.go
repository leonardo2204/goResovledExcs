package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func(reader MyReader)Read(a []byte) (int,error) {
	a[0] = 'A'
	return 1,nil
}

func main() {
	reader.Validate(MyReader{})
}
