package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type Case struct {
	input  int
	expect bool
}

func TestIsPrimeNumber(t *testing.T) {
	tests := []Case{
		{
			input:  41,
			expect: true,
		},
		{
			input:  78,
			expect: false,
		},
	}

	for _, test := range tests {
		ans := IsPrimeNumber(test.input)
		assert.DeepEqual(t, test.expect, ans)
	}
}

func IsPrimeNumber(n int) bool {
	var result = true
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			result = false
			break
		}
	}

	return result
}

type Case2 struct {
	start  int
	end    int
	expect []int
}

func TestGetPrimeNumbersBetween(t *testing.T) {
	tests := []Case2{
		{
			start:  1,
			end:    10,
			expect: []int{2, 3, 5, 7},
		},
		{
			start:  1,
			end:    60,
			expect: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59},
		},
	}

	for _, test := range tests {
		ans := GetPrimeNumbersBetween(test.start, test.end)
		t.Log(ans)
		assert.DeepEqual(t, test.expect, ans)
	}
}

func GetPrimeNumbersBetween(left, right int) []int {
	var result []int

	/*
		첫번째 수의 배수를 모두 체크한다
		두번째 수의 배수를 모두 체크한다 (앞전의 체크된 수 제외하고)
		세번째 수의 배수를 모두 체크한다 (앞전의 체크된 수 제외하고)
		...
		더 이상 체크할 수 있는 수가 없으면 종료한다

		=. 체크되지 않고 남아있는 수를 모아서 결과값으로 돌려준다.
	*/

	var checkedNumbers = make(map[int]bool)

	var start = left
	var end = right

	if left == 1 {
		start = 2
	}

	for i := start; i < end+1; i++ {
		for j := 2; j < (end+1/i)+1; j++ {
			notPrime := i * j
			if _, exist := checkedNumbers[notPrime]; !exist {
				checkedNumbers[notPrime] = true
			}
		}
	}

	for i := start; i < end+1; i++ {
		if !checkedNumbers[i] {
			result = append(result, i)
		}
	}

	return result
}

/*
function get_primes(num){
    const prime = [false, false, ...Array(num-1).fill(true)]

    for (let i = 2 ; i * i <=  num; i += 1){
        if (prime[i]){
            for (let j = i * 2 ; j <= num ; j += i){
                prime[j] = false
            }
        }
    }

    return prime.map((num,idx) => {
        if (num){
            return idx
        }else{
            return 0
        }}).filter(num => num > 0)
}
*/
