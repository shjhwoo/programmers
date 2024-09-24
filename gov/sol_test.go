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
	num    int
	expect []int
}

func TestGetPrimeNumbersBetween(t *testing.T) {
	tests := []Case2{
		{
			num:    10,
			expect: []int{2, 3, 5, 7},
		},
		{
			num:    59,
			expect: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59},
		},
	}

	for _, test := range tests {
		ans := GetPrimeNumbersLessThanOrEqualTo(test.num)
		t.Log(ans)
		assert.DeepEqual(t, test.expect, ans)
	}
}

func GetPrimeNumbersLessThanOrEqualTo(num int) []int { //가장 효율적이다.

	var isPrime = []bool{}
	for i := 0; i < num+1; i++ {
		if i <= 1 {
			isPrime = append(isPrime, false)
		} else {
			isPrime = append(isPrime, true)
		}
	}

	for i := 2; i*i < num; i++ {
		if isPrime[i] {
			for j := i * 2; j <= num; j += i {
				isPrime[j] = false
			}
		}
	}

	var result []int

	for idx, candidate := range isPrime {
		if candidate {
			result = append(result, idx)
		}
	}

	return result
}

/*
function get_primes(num){
    const prime = [false, false, ...Array(num-1).fill(true)] //0과 1은 소수가 아니라서 false로 둔 것임.

    for (let i = 2 ; i * i <=  num; i += 1){ //어짜피 a * b === n 이면, a로 나눠떨어지면 b로 나눠떨어지는 건 똑같으니까 b에 대해서까지 확인 필요없다.
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
