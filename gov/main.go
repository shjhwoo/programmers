package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	line := ""
	for i := 0; i < a; i++ {
		line += "*"
	}

	for j := 0; j < b; j++ {
		fmt.Println(line)
	}

	//fmt.Println(a + b)
}
