package main_test

import (
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	fees    []int
	records []string
	expect  []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			fees:    []int{180, 5000, 10, 600},
			records: []string{"05:34 5961 IN", "06:00 0000 IN", "06:34 0000 OUT", "07:59 5961 OUT", "07:59 0148 IN", "18:59 0000 IN", "19:09 0148 OUT", "22:59 5961 IN", "23:00 5961 OUT"},
			expect:  []int{14600, 34400, 5000},
		},
		{
			fees:    []int{120, 0, 60, 591},
			records: []string{"16:00 3961 IN", "16:00 0202 IN", "18:00 3961 OUT", "18:00 0202 OUT", "23:58 3961 IN"},
			expect:  []int{0, 591},
		},
		{
			fees:    []int{1, 461, 1, 10},
			records: []string{"00:00 1234 IN"},
			expect:  []int{14841},
		},
	}

	for _, test := range tests {
		ans := solution(test.fees, test.records)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(fees []int, records []string) []int {
	//일단 들어온 차량번호 => 번호기준 오름차순으로 정리해야함

	recordMap := make(map[int][]string) //차량번호 - 입차, 출차시각 맵
	for _, record := range records {
		timeStamp, carNum := strings.Split(record, " ")[0], strings.Split(record, " ")[1]

		carNumInt, _ := strconv.Atoi(carNum)

		recordMap[carNumInt] = append(recordMap[carNumInt], timeStamp)
	}

	var carNumList []int
	for carNum := range recordMap {
		carNumList = append(carNumList, carNum)
	}

	sort.Slice(carNumList, func(i, j int) bool {
		return carNumList[i] < carNumList[j]
	})

	var answer []int
	for _, carNum := range carNumList {
		//요금계산해서 asnwer에 넣기
		fee := calculateFee(fees, recordMap[carNum])
		answer = append(answer, fee)
	}

	return answer
}

func calculateFee(fees []int, inoutTime []string) int {
	if len(inoutTime)%2 == 1 {
		inoutTime = append(inoutTime, "23:59")
	}

	var totalDuration int
	for i := 0; i < len(inoutTime); i += 2 {
		inTime, _ := time.Parse("15:04", inoutTime[i])
		outTime, _ := time.Parse("15:04", inoutTime[i+1])
		duration := outTime.Sub(inTime)
		durationMinutes := int(duration.Minutes())
		totalDuration += durationMinutes
	}

	baseTime, baseFee, unitTime, unitFee := fees[0], fees[1], fees[2], fees[3]

	if totalDuration <= baseTime {
		return baseFee
	}

	overTime := totalDuration - baseTime

	if overTime%unitTime > 0 {
		return baseFee + ((overTime/unitTime)+1)*unitFee
	}

	return baseFee + (overTime/unitTime)*unitFee
}
