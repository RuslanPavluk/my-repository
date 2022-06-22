package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var results[]int
	
	unique:=make(map[int]bool)

	for _, val:=range arr{

		if _, in := unique[val]; !in {

			unique[val]=true
			results=append(results, val)
			fmt.Println("results=",results)
		}
	}
}