package main_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	numbers []int
	expect  []int
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			numbers: []int{2, 3, 3, 5},
			expect:  []int{3, 5, 5, -1},
		},
		{
			numbers: []int{9, 1, 5, 3, 6, 2},
			expect:  []int{-1, 5, 6, 6, -1, -1},
		},
	}

	for _, test := range tests {
		if !assert.True(t, slices.Compare(test.expect, solution(test.numbers)) == 0) {
			t.Log(solution(test.numbers))
		}
	}
}

func solution(numbers []int) []int {
	/*
			   빨리 찾으려면..

			   혹시 캐시같은걸 써야하나..

			   예를 들어서,
			   9,1,5,3,6,2 배열이라면

			   9의 경우는 어쩔 수 없이 다 돌아봐야 해. 처음 한번은. 안 돌고서는 진짜 모르니까!

			   그러면 이때 뒷 큰수를 찾기.. 없으니 -1

			   1의 경우는 뒷 숫자가 1이 아니라면 빨리 뒷 큰수 지정 가능함.

			   9: 15362
			   1: 5362
			   5: 362
			   3: 62
			   6: 2
			   2: -


			   방법> 뒷 큰수를 찾을 떄마다 캐시에다가 저장해두는 거야.
			   그러니까 특정 인덱스에 있는 숫자의 뒷 큰수 인덱스를 맵으로 저장해두는 거지!

			   위의 경우를 예로 들자면
			   key value
			   [0] : -1
		       [1] : 2
			    [2]: 4
		       [3]: 4
			   [4]: -1
			   [5]: -1

			   그러면 이런 경우에 유용할 거야.
			   2,1,1,1,1,1,7,10 배열이라고 해봐.
			   [0]: 6
			   [1]: 6
			   ... (더이상 묻따 없음)
			   [6]: 7
			   [7]: -1

			   왜냐하면 중간의 1 다섯개들은, 다음 질문을 통해서 뒷 큰 수를 찾을 건데,
			   처음에 등장하는 1은, 일단 뒤에 뭐가 더 있는지 몰라, 그치만 앞에서 7이라는 큰 수 정보를 찾았어.
			   그래서 일단은 7을 맵핑해놓는거지.
			   그런데 얘보다 더 가까운 큰 수가 있을 수도 있겠지? 예를 들면 5 같은..
			   그래서 일단은, 2번 인덱스부터, 6번 인덱스 까지만 다시 순회를 하는거야.
			   여기서 만약에, 7보다 더 큰 수가 나온다면 인덱스 6을 버리고 그 수에 대한 인덱스를 찾으면 되겠지?
			   오호.!! 해결점을 찾았어~~ 내부 반복문에서 굳이 끝까지 다 안돌아도 된다는 거자나.

			   그러면 일단 이런 맵부터 먼저 만들어 보는거야.

			   그래도 시간 초과했다고 뜨네 마지막 4케이스에서.

			   그러면 이 경우는?
			   10,9,8,7,5,4,3,2,1

			   죄다 -1이거든..

			   이런 배열 길이가 10만 이라고 해봐.. 말도안돼.
			   그러면,, 이 경우는 어떻게 빨리 처리할까?

	*/

	var answer []int

	var cache = make(map[int]int) //현재수의 인덱스 - 뒷 큰수에 대한 인덱스.

	for i := 0; i < len(numbers)-1; i++ {
		found := 0

		//어디까지 돌건지 결정하는 지점.
		var endPoint int
		//만약에 맨 처음이거나, 뒷 큰수가 없거나, 뒷 큰수가 있는데 인덱스가 지금 수와 같으면! 그냥 배열 마지막까지 다 봐야 해.
		if cache[i-1] == -1 || cache[i-1] == 0 || i == cache[i-1] {
			endPoint = len(numbers)
		} else {
			endPoint = cache[i-1] + 1
		}

		for j := i + 1; j < endPoint; j++ {
			if numbers[i] < numbers[j] {
				cache[i] = j //일단 캐시에다가 저장해!

				found = 1

				answer = append(answer, numbers[j])
				break
			}
		}

		if found == 0 {
			answer = append(answer, -1)
		}

	}

	answer = append(answer, -1)

	return answer
}

/*
테스트 20 〉	실패 (시간 초과)
테스트 21 〉	실패 (시간 초과)
테스트 22 〉	실패 (시간 초과)
테스트 23 〉	실패 (시간 초과)
*/
