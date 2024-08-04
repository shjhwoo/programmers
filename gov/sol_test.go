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
		{
			bridge_length: 100,
			weight:        100,
			truck_weights: []int{10},
			expect:        101,
		},
		{
			bridge_length: 100,
			weight:        100,
			truck_weights: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			expect:        110,
		},
	}

	for _, test := range tests {
		ans := solution(test.bridge_length, test.weight, test.truck_weights)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

/*
숙제..
1초에 1칸씩 갈 수 있지만
초 = 거리 이므로
무게 제한에 걸렸을 때 차량들을 한번에 끝까지 보내버리고 그 전체 거리 시간 계산에 더해버리면 줄일 수 있음.

1초씩 처리하고 시간 초과하면 튜닝하려고 했는데 그냥 통과해버려서 넘어가려다가 찝찝해서 튜닝
*/
func solution(bridge_length int, maxWeight int, waiting_trucks []int) int {
	var bridge []int
	bridge = append(bridge, waiting_trucks[0])
	current_weight := waiting_trucks[0]

	var answer = 1
	var truckIdx = 1
	for current_weight > 0 {

		var newTruck int
		if truckIdx < len(waiting_trucks) {
			newTruck = waiting_trucks[truckIdx]
		}

		newWeight := current_weight + newTruck
		if len(bridge) < bridge_length {
			if newWeight <= maxWeight {
				bridge = append(bridge, newTruck)
				current_weight += newTruck
				truckIdx++
			} else if newWeight > maxWeight {
				//bridge = append(bridge, 0) //하나씩 붙이지 말고 남은 여유 구간만큼 0을 다 붙이고..
				bridge = bridge[1:]
				bridge = append(bridge, newTruck)
				current_weight = newTruck
				answer += bridge_length
				continue
			}
		} else if len(bridge) == bridge_length {

			firstTruck := bridge[0]
			if current_weight-firstTruck+newTruck <= maxWeight {
				bridge = bridge[1:]
				bridge = append(bridge, newTruck)
				truckIdx++
				current_weight = current_weight - firstTruck + newTruck
			} else if current_weight-firstTruck+newTruck > maxWeight {
				bridge = bridge[1:]
				bridge = append(bridge, 0)
				current_weight -= firstTruck
			}
		}

		answer++
	}

	return answer
}
