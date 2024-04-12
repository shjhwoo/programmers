package main_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	today     string
	terms     []string
	privacies []string
	expect    []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			today:     "2022.05.19",
			terms:     []string{"A 6", "B 12", "C 3"},
			privacies: []string{"2021.05.02 A", "2021.07.01 B", "2022.02.19 C", "2022.02.20 C"},
			expect:    []int{1, 3},
		},
		{
			today:     "2020.01.01",
			terms:     []string{"Z 3", "D 5"},
			privacies: []string{"2019.01.01 D", "2019.11.15 Z", "2019.08.02 D", "2019.07.01 D", "2018.12.28 Z"},
			expect:    []int{1, 4, 5},
		},
	}

	for _, test := range tests {
		ans := solution(test.today, test.terms, test.privacies)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(today string, terms []string, privacies []string) []int {

	var termMap = getTermMap(terms)

	todayDate, err := time.Parse("20060102", strings.ReplaceAll(today, ".", ""))
	if err != nil {
		fmt.Println("날짜 형식이 잘못되었습니다.", err)
		return nil
	}

	var answer []int

	for idx, priv := range privacies {
		if IsAbleToDiscard(priv, termMap, todayDate) {
			answer = append(answer, idx+1)
		}
	}

	return answer
}

func getTermMap(terms []string) map[string]int {
	var result = make(map[string]int)

	for _, term := range terms {
		tsl := strings.Split(term, " ")

		termName := tsl[0]

		term, err := strconv.Atoi(tsl[1])
		if err != nil {
			continue
		}

		result[termName] = term
	}

	return result
}

func IsAbleToDiscard(privacyInfo string, termMap map[string]int, todayDate time.Time) bool {
	collectDate, termName := getPrivacyInfo(privacyInfo)
	discardDate := getDiscardDate(collectDate, termMap[termName])

	return todayDate.After(discardDate) || todayDate.Equal(discardDate)
}

func getPrivacyInfo(privacyInfo string) (time.Time, string) {
	tsl := strings.Split(privacyInfo, " ")

	collectedDateStr := strings.ReplaceAll(tsl[0], ".", "")
	termName := tsl[1]

	collectedDateTime, err := time.Parse("20060102", collectedDateStr)
	if err != nil {
		return time.Time{}, ""
	}

	return collectedDateTime, termName
}

func getDiscardDate(collectedDate time.Time, term int) time.Time {
	return collectedDate.AddDate(0, term, 0)
}
