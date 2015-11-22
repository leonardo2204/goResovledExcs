package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	
	var myMap = make(map[string]int)
	
	for _,value := range(strings.Fields(s)) {
		val,exists := myMap[value]
		if !exists {
			myMap[value] = 1
		}else{
			myMap[value] = val +1;
		}
	}
	
	return myMap
}

func main() {
	wc.Test(WordCount)
}
