package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	a, err := rot.r.Read(b)
	s := string(b[:a])

	for i, v := range s {
		c := fmt.Sprintf("%c", v)
		if c >= "a" && c <= "m" {
			b[i] += 13
		} else if c >= "n" && c <= "z" {
			b[i] -= 13
		} else if c >= "A" && c <= "M" {
			b[i] += 13
		} else if c >= "N" && c <= "Z" {
			b[i] -= 13
		}
	}

	return a, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
