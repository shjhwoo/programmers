package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution(2000))
	fmt.Println(solution(3))
}

var cache = make(map[int]int64)

func solution(n int) int64 {
	cache[0] = 0

	if n == 1 || n == 2 {
		cache[n] = int64(n)
		return int64(n % 1234567)
	}

	if cache[n] != 0 {
		return int64(cache[n] % 1234567)
	}

	ans := (solution(n-2) + solution(n-1))
	cache[n] = ans

	return cache[n] % 1234567
}
