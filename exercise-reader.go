/*
Exercise: Readers
From: https://tour.golang.org/methods/22
*/

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (reader MyReader) Read(b []byte) (int, error) {
	var length = len(b)

	for i := 0; i < length; i++ {
		b[i] = byte('A')
	}

	return length, nil
}

func main() {
	reader.Validate(MyReader{})
}
