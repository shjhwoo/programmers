package main_test

import (
	"fmt"
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	genres []string
	plays  []int
	expect []int
}

// 포인트는 인덱스를 같이 저장하는 것..!!
func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			genres: []string{"classic", "pop", "classic", "classic", "pop"},
			plays:  []int{500, 600, 150, 800, 2500},
			expect: []int{4, 1, 3, 0},
		},
		{
			genres: []string{"classic", "pop"},
			plays:  []int{500, 600},
			expect: []int{1, 0},
		},
		{
			genres: []string{"classic", "pop", "classic", "classic", "classic"},
			plays:  []int{500, 600, 500, 1900, 300},
			expect: []int{3, 0, 1},
		},
	}

	for _, test := range tests {
		ans := solution(test.genres, test.plays)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

/*
장르 별로 가장 많이 재생된 노래를 두 개씩

노래 = 고유번호 키

속한 노래가 많이 재생된 장르를 먼저 수록합니다.
장르 내에서 많이 재생된 노래를 먼저 수록합니다.
장르 내에서 재생 횟수가 같은 노래 중에서는 고유 번호가 낮은 노래를 먼저
*/

type BestGenRe struct {
	GenreName       string
	PlaySum         int
	SongIdxPlayList []*SongPlayTime
}

type SongPlayTime struct {
	SongIdx int
	Play    int
}

func solution(genres []string, plays []int) []int {

	var genreMap = make(map[string]*BestGenRe)

	for idx, genre := range genres {
		if genreMap[genre] == nil {
			genreMap[genre] = &BestGenRe{}
			genreMap[genre].GenreName = genre
		}
		genreMap[genre].PlaySum += plays[idx]

		genreMap[genre].SongIdxPlayList = append(genreMap[genre].SongIdxPlayList, &SongPlayTime{
			SongIdx: idx,
			Play:    plays[idx],
		})
	}

	fmt.Println(genreMap, "genreMap 확인중!")

	var topGen []*BestGenRe
	for _, Info := range genreMap {
		topGen = append(topGen, Info)
	}

	sort.Slice(topGen, func(i, j int) bool {
		return topGen[i].PlaySum > topGen[j].PlaySum
	})

	if len(topGen) > 2 {
		topGen = topGen[:2] //상위 2개만.
	}

	fmt.Println(topGen[0].GenreName, topGen[1].GenreName, "topGen")

	var answer []int
	for _, item := range topGen {
		sort.Slice(item.SongIdxPlayList, func(i, j int) bool {
			if item.SongIdxPlayList[i].Play != item.SongIdxPlayList[j].Play {
				return item.SongIdxPlayList[i].Play > item.SongIdxPlayList[j].Play
			}

			return i < j
		})

		if len(item.SongIdxPlayList) > 2 {
			item.SongIdxPlayList = item.SongIdxPlayList[:2]
		}

		for _, item := range item.SongIdxPlayList {
			answer = append(answer, item.SongIdx)
		}

	}

	return answer
}
