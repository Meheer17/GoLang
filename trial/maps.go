package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	list := make(map[string]int) 
	words := strings.Split(s, " ")
	for i := range words {
		v, ok := list[words[i]]
		if ok == false {
			list[words[i]] = 1
		} else {
			list[words[i]] = v + 1
		}
	}
	return list
}

func main() {
	wc.Test(WordCount)
}
