package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution(4))
	fmt.Println(solution(3))
}

//멀리뛰기에 사용될 칸의 수를 나눠 생각하자
//1이상 2000이하 값이 들어올 수 있다.
//n이 1인 경우 1가지
//n이 2인 경우 (1,1) (2) 2가지  -- 2가 최대 1번
//n이 3인 경우 (1,1,1) (2,1) (1,2) 3가지 --2가 최대 1번
//n=4 인 경우 (1,1,1,1) (2,1,1) (1,2,1) (1,1,2) (2,2) 5가지 --2가 최대 2번
//n=5인 경우 (1,1,1,1,1) (2,1,1,1) (1,2,1,1) (1,1,2,1) (1,1,1,2) (2,2,1) (2,1,2) (1,2,2) 7가지 -- 2가 최대 2번
//n=6 --2가 최대 3번

//n이 짝수인 경우

//홀수인 경우: 1은 무조건 1개 들어갈 수 밖에 없고
// 1,2로 홀수-1를 만드는 방법

//아니면 2가 들어갈 수 있는 최대횟수를 구한다
/*
	2가 0번 --- 2 0개와 1 숫자 개를 순서 구분해서 세울 수 있는 방법 (x)
	2가 1번 --- 2 1개와 1 (숫자-2)개 를 순서 구분해서 세울 수 있는 방법 : 숫자 - 1
	2가 2번 --- 2 2개와 1 (숫자-4)개를 순서 구분해서 세울 수 있는 방법: 숫자 - 2
	...
	(숫자를 2로 나눈 몫 =q) 번 --- 2 q개와 1 (숫자-2q)개를 순서 구분해서 세울 수 있는 방법

	이 나열한 방법의 수를 모두 더해준다.
	그리고 그 값을 1234567로 나눠준다
*/
func solution(n int) int64 {
	if n == 1 {
		return 1
	}

	q := n / 2 //몫
	var count int64
	for i := 0; i <= q; i++ {
		r := permutationCount(int64(i), int64(n-2*i))
		if i%2 == 0 && i == q {
			count = count + 1
			continue
		}
		count = count + int64(r)
	}
	return count % 1234567
}

func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func permutationCount(X, Y int64) int64 {
	n := X + Y
	r := X
	permutations := factorial(n) / factorial(n-r)
	return permutations
}
