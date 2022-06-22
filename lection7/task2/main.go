package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	
	input:= "1 9 3 4 -5"
	 
	var results []string
	var slice1 []string
	 
	slice1=strings.Split(input, " ")
	fmt.Printf("%q, \n", slice1)
	
	sort.Strings(slice1)
	fmt.Println("sort =", slice1)

	min:=slice1[0]
	max:=slice1[len(slice1)-1]
	fmt.Println("min, max=",min,max)

	results=append(results,string(max),string(min))
	fmt.Println("results=",results)

}