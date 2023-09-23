package main_test

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	res1 := Solution(16)
	assert.Equal(t, 6, res1)

	res2 := Solution(2554)
	assert.Equal(t, 16, res2)

	//XXX
	res3 := Solution(6628)
	assert.Equal(t, 13, res3)
}

func Solution(storey int) int {
	logVal := math.Log10(float64(storey))
	if logVal == float64(int(logVal)) {
		return 1
	}

	numSlice := getNumSliceFromStorey(storey)

	//모든 자릿수가 5 이상인 경우 또는 첫번째 자리 숫자가 5 이상인 경우
	/*
		첫번째 자리가 5이상인경우(다음10지수, 다음으로 큰 수 중에 비용이 적은 것을 선택)
			두번째자리가 5 이상인경우
			5739
			두번쨰자리가 5이하인경우
			5468
			=> 두번째 자리가 뭐가 되었든간에, 5 이상이라 다은 10지수로 계산하는게 제일 최소 비용이다.
	*/
	if numSlice[0] >= 5 {
		rem := int(math.Pow10(len(numSlice)+1)) - storey
		var stones int
		for _, r := range getNumSliceFromStorey(rem) {
			stones += r
		}
		return 1 + stones
	}

	//모든 자릿수가 5 미만인 경우
	if isAllnumIsLessThanFive(numSlice) {
		var res int
		for _, n := range numSlice {
			res += n
		}
		return res
	}

	/*
		//혼재하는 경우ㅠㅠ
			첫번째 자리가 5미만인경우
			두번째 자리가 5이상인경우
			4523
			4566
			두번째 자리가 5이하면
			4268
			4238
	*/
	var candidates []int
	var can1 int
	for _, n := range numSlice {
		can1 += n
	}
	candidates = append(candidates, can1)

	nextNum := (numSlice[0] + 1) * int(math.Pow10(len(numSlice)))
	rem := nextNum - storey
	var remSum int
	for _, n := range getNumSliceFromStorey(rem) {
		remSum += n
	}
	can2 := numSlice[0] + 1 + remSum
	candidates = append(candidates, can2)

	var minNum = candidates[0]
	for _, c := range candidates {
		if c < minNum {
			minNum = c
		}
	}
	return minNum

}

/*
누적해서 보여줄 결과값
현재 자리
현재 자리 바로 다음 숫자

 1. 현재 자리가 첫번째 자리고 5보다 크다
    res에 1 더함
    rem에 10거듭제곱 - 숫자 뺀 값 더한다
 2. 현재 자리가 첫번쨰 자리고 5이하다.
 3. 다음 숫자가 5 이하라면 숫자를 그대로 더한다
 4. 다음 숫자가 6이상이라면 다음으로 큰 수를 더한다.
    rem에다 다음으로 큰수 - 현재수 빼준다

다음 iter에서는?
현재자리와 다음 자리, 남아있는 숫자만 가지고 판단하면 되니까 굳이 numSlice 전체를 볼 일은 없을듯
1) remainder가 있는 경우
2) 없는 경우
*/

// remainder가 생길지 안생길지 판단하는 함수
func hasRemainder(numSlice []int) bool {
	return numSlice[0] > 5 || (numSlice[0] <= 5 && numSlice[1] > 5)
}

func getRemainder(numSlice []int, storey int) int {
	if numSlice[0] > 5 {
		pow := int(math.Ceil(math.Log10(float64(storey))))
		maxBtn := math.Pow10(pow)
		return int(maxBtn) - storey
	}

	if numSlice[1] > 5 {
		return (numSlice[1]+1)*int(math.Pow10(len(numSlice)-1)) - getNumFromSlice(numSlice)
	}

	return 0
}

func isAllnumIsGreaterOrEqualThanFive(numSlice []int) bool {
	for _, n := range numSlice {
		if n < 5 {
			return false
		}
	}

	return true
}

func isAllnumIsLessThanFive(numSlice []int) bool {
	for _, n := range numSlice {
		if n >= 5 {
			return false
		}
	}

	return true
}

func getNumSliceFromStorey(storey int) []int {
	numStrSlice := (strings.Split(strconv.Itoa(storey), ""))

	var numSlice []int
	for _, ns := range numStrSlice {
		in, _ := strconv.Atoi(ns)
		numSlice = append(numSlice, in)
	}

	return numSlice
}

func getNumFromSlice(numSlice []int) int {
	resStr := ""
	for _, n := range numSlice {
		resStr += strconv.Itoa(n)
	}

	n, err := strconv.Atoi(resStr)
	if err != nil {
		return 0
	}
	return n
}

/*경우의 수
규칙 정리
1) 숫자의 모든 자릿수가 5보다 크면 10배 큰 수로 한번 뺴고 보충하는 게 훨씬 이득임
2) 그렇지 않은 경우..
각 자릿수에 대해서 5보다 큰지 작은지 판단해보자
5이하다. ! 그러면 이 숫자 뒤에 다음 숫자가 있고,
그 숫자가 5보다 크다 그러면 (28인 경우 30으로) 잡고 더 빠진거는 보충
5보다 크다. 맨 앞자리가 아니라면 이전의 숫자가 알아서 처리해줘버린다.





1인 경우 1개
2인 경우 2개
3인 경우 3개
4인 경우 4개
5인 경우 5개
6인 경우 -10 + 4 => 5개 . 6개 쓰는 것보다 이득임
7인 경우 -10 + 3 => 4개
8인 경우 -10 + 2 => 3개
9인 경우 -10 + 1 => 2개
10인 경우 -10 => 1개

11인 경우 -10 -1 2개
12인 경우 -10 -2 3개
13인 경우 -10 -3 4개
14인 경우 -10 -4 5개
15인 경우 -10 -5 6개
16인 경우 -10 -10 + 4 6개
17인 경우 -10 -10 + 3 5개
18인 경우 -10 -10 + 2 4개
19인 경우 -10 -10 + 1 3개
20인 경우 -10 -10 2개

21인 경우 -10 -10 -1 3개
22인 경우 -10 -10 -2 4개
23인 경우 -10 -10 -3 5개
24인 경우 -10 -10 -4 6개
25인 경우 -10 -10 -5 7개
26인 경우 -10 -10 -10 + 4 7개
27인 경우 -10 -10 -10 + 3 6개
28인 경우 -10 -10 -10 + 2 5개
29인 경우 -10 -10 -10 + 1 4개
30인 경우 -10 -10 -10 3개

...
50인 경우 -10 * 5 5개
51인 경우 -10 * 5 + -1 6개
52인 경우 -10 * 5 + -1 * 2 7개
53인 경우 -10 * 5 -1 * 3 8개
54인 경우 -10 * 5 -1 * 4 9개 또는 -10 * 6 + 1 * 6
55인 경우 -10 * 5 -1 * 5 10개 또는 -10 * 6 + 1 * 5 11개
56인 경우 -10 * 5 -1 * 6 11개 또는 -10 * 6 +  1 * 4 10개
57인 경우 -10 * 5 -1 * 7 12개 또는 -10 * 6 + 1 * 3 9개 === (-10 * 5 , -10 * 1 + 1 * 3)


60인 경우 6개 쓰는 것보다, -100 + 10 + 10 + 10 + 10 5개 쓰는게 이득
66인 경우 :
-100 + 10 * 3 + 1 * 4  = 8개
-100 + 10 * 4  + -10 + 1 * 4 = 9개


101인 경우 -100 -1 2개
106인 경우 -100 -10 + 4 vs -100 -6
107인 경우 -100 -10 + 3 vs -100 -7

116인 경우 -100 -20 + 4

2554인 경우
각 자리수에 대해 2가지 경우가 있을 수 있다...?
그 두가지 경우 중 하나를 고르는 기준은 무엇일까? 그리고 그 기준대로 계산하는 게 맞는건가
1000: -1000 * 2 또는 -1000 * 3
100: 앞 자리수에 따라 달라지는데,
-1000 * 2를 선택한경우.
-100 * 6  또는 -100 * 5가 될수있다.

-1000 * 3을 선택한경우
446만큼 다시 올라가야 하는데
-100 * 5 또는 -100 * 4
(-10 * 4 + -10 * 1 + 4)

그러면 정리하면 주어진 숫자보다 최소로 더 큰 다음 시작점숫자를 찾는다. 끝자리수는 0으로 끝나야한다
2554인 경우 3000
12345인 경우 20000
857인 경우 900
593인 경우 600 이런식.

그 다음에 다음 기점의 숫자에 도달하려할때 이게 중간지점을 안넘은 경우랑 넘은 경우로 판단할 수 있는데
2554인 경우 2000~3000 중간인 2500을 넘었으니까,
1000을 두번빼는것보다 3번빼고 더 내려간거를 채우는게 이득이라고 판단하는 것

2222였다면, 그냥 1000을 두번빼는게 맞았겠지? (돌 8개)

2282였다면 , -1000 * 2 + -100 * 2 + -100 * 1 + 10 * 2 -2 (돌9개)
-1000 * 2 + -100 * 2 + -10 * 8 -1 * 2
결국 숫자가 5를 넘냐 안넘냐가 판단의 기준인것
맨 첫글자가 5를 넘는다 하면 아얘 다음 자릿수로 빼버리는 것이 이득일지도


6000인 경우 -10000 + 4000 5개 쓰는 게 이득
6600인 경우 -10000 + 3400

6628인 경우 -10000 + 3372 (16) 또는 -10000 + 3400 - 30 + 2 (13)

그러면 6666의 경우에는?
-10000 + 3334 => 1000 * 3 + 100 * 3 + 10 * 3 + 4 => 1개 + 3 + 3 + 3 + 4 => 14개
또는
다음 숫자가 7000이라고 보고

중간인 6500넘었으니까
-1000 * 7 + 100 * 3 + 10 * 3 + 1 * 4 ==> 17개..

=> 각 자리의 숫자가 5를 넘은 경우에는 2가지 가짓수가ㅏ 있음
다음으로 큰 최소수
또는 10의 거듭제곱 기준으로 한 다음 최소수 (이게 더 적을수도 있기 때문에ㅠㅠ)


...10의 n승 기준으로 기점이 나뉘어진다.
숫자가 주어지면 그 숫자보다 최소로 큰 10의 거듭제곱값을 찾는다?

그리고 10의 거듭제곱으로 이루어진 숫자라면 그거는 무조건 1개만 쓰면 된다
*/
