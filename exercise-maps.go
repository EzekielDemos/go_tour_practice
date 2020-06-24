/*
Exercise: Maps
From: https://tour.golang.org/moretypes/23
*/
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	w := map[string]int{}
	sarray := strings.Fields(s)
	for _, v := range sarray {
		w[string(v)] += 1
	}
	return w
}

func main() {
	wc.Test(WordCount)
}
