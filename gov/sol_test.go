package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	bridge_length int
	weight        int
	truck_weights []int
	expect        int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			bridge_length: 2,
			weight:        10,
			truck_weights: []int{7, 4, 5, 6},
			expect:        8,
		},
		// {
		// 	bridge_length: 100,
		// 	weight:        100,
		// 	truck_weights: []int{10},
		// 	expect:        101,
		// },
		// {
		// 	bridge_length: 100,
		// 	weight:        100,
		// 	truck_weights: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		// 	expect:        110,
		// },
	}

	for _, test := range tests {
		ans := solution(test.bridge_length, test.weight, test.truck_weights)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

/*
일차선 다리를 정해진 순으로 건너기

다리에는 트럭이 최대 bridge_length대 올라갈 수 있으며

다리는 weight 이하까지의 무게를 견딜 수 있습니다.

모든 트럭이 다리를 건너려면 최소 몇 초가 걸리는지

경과 시간	다리를 지난 트럭	다리를 건너는 트럭	대기 트럭
0	[]	[_ _]	[7,4,5,6]
1	[]	[_ 7]	[4,5,6]
2   []  [7 _]   [4,5,6]
3	[7]	[_ 4]	[5,6]
4	[7]	[4,5]	[6]
5	[7,4] [5 _]	[6]
6	[7,4,5]	[_ 6]	[]
7	[7,4,5]	[6 _]	[]
8	[7,4,5,6]	[]	[]
*/
func solution(bridge_length int, maxWeight int, waiting_trucks []int) int {
	bridge := []int{} // 왼쪽이 입구 -------> 오른쪽이 출구
	bridge = append(bridge, waiting_trucks[0])
	current_weight := waiting_trucks[0]

	var answer = 1
	var truckIdx = 1
	for current_weight > 0 {

		//마지막 처리를 어떻게 해야하지..ㅠㅠ
		if truckIdx == len(waiting_trucks) {
			if len(bridge) == 1 {
				answer += bridge_length
			}
			break
		}

		/*
			일단 새로운 트럭을 다리에 밀어 넣었을 때
			1) 여유가 있어서 들어와도 되는 경우
			2) 무게는 감당할 수 있지만 다리 길이가 모자라서 못 들어가는 경우
			3) 무게가 한도를 넘어가는 경우
		*/

		newTruck := waiting_trucks[truckIdx]
		newWeight := current_weight + newTruck

		if len(bridge) < bridge_length {
			if newWeight <= maxWeight { //1) 여유가 있어서 들어와도 되는 경우
				bridge = append(bridge, newTruck)
				current_weight += newTruck
				truckIdx++
			} else if newWeight > maxWeight {
				//대기해야 한다. 가짜 트럭으로 0을 넣어버리기~
				bridge = append(bridge, 0)
			}
		} else if len(bridge) == bridge_length {
			if current_weight-bridge[0]+newTruck <= maxWeight {
				firstTruck := bridge[0]
				bridge = bridge[1:]
				bridge = append(bridge, newTruck)
				truckIdx++
				current_weight = current_weight - firstTruck + newTruck
			} else if current_weight-bridge[0]+newTruck > maxWeight {
				//넣을 수 없다. 대신 공백으로 0을 채운다
				firstTruck := bridge[0]
				bridge = bridge[1:]
				bridge = append(bridge, 0)
				current_weight -= firstTruck
			}
		}

		answer++

	}

	return answer
}
