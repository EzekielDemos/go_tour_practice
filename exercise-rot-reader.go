/*
Exercise: rot13Reader
From: https://tour.golang.org/methods/23
*/

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	var length = len(b)
	var tb = make([]byte, length)
	n, err := reader.r.Read(tb)

	if err != nil {
	 return n, err
	} else {
		for i := 0; i < length; i++ {
			if (tb[i] + 13 > 122) {
			  b[i] = tb[i] - 13
			} else {
			  b[i] = tb[i] + 13
			}
		}
	}

	return length, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
