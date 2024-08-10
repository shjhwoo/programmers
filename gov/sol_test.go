package main_test

import (
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
			genres: []string{"classic", "pop", "classic", "classic", "pop"}, //  5 1 2 3 4
			plays:  []int{500, 600, 150, 800, 2500},
			expect: []int{4, 1, 3, 0},
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
	SongIdxPlaysMap map[int]int
}

func solution(genres []string, plays []int) []int {

	var genreMap = make(map[string]int)

	for idx, genre := range genres {
		genreMap[genre] += plays[idx]
	}

	var bestAlbum = make(map[int]BestGenRe)

	var idx = 1
	for len(bestAlbum) < 2 {
		var maxPlay int
		var maxGenre string
		for genreK, playSum := range genreMap {
			if maxPlay < playSum {
				maxPlay = playSum
				maxGenre = genreK

			}
		}

		bestAlbum[idx] = BestGenRe{
			GenreName: maxGenre,
		}
		delete(genreMap, maxGenre)
		idx++
	}

	first := bestAlbum[1].GenreName
	second := bestAlbum[2].GenreName
	for idx, genre := range genres {
		if genre == first {
			if val, ok := bestAlbum[1]; ok {
				if val.SongIdxPlaysMap == nil {
					val.SongIdxPlaysMap = make(map[int]int)
				}
				val.SongIdxPlaysMap[idx] = plays[idx]
				bestAlbum[1] = val
			}
		}

		if genre == second {
			if val, ok := bestAlbum[2]; ok {
				if val.SongIdxPlaysMap == nil {
					val.SongIdxPlaysMap = make(map[int]int)
				}
				val.SongIdxPlaysMap[idx] = plays[idx]
				bestAlbum[2] = val
			}
		}
	}

	var answer []int
	for i := 1; i < 3; i++ {
		var top2 []int
		for len(top2) < 2 {

			var max int
			var maxIdx int
			for idx, play := range bestAlbum[i].SongIdxPlaysMap {
				if max < play {
					max = play
					maxIdx = idx
				}
			}

			top2 = append(top2, maxIdx)
			delete(bestAlbum[i].SongIdxPlaysMap, maxIdx)
		}

		answer = append(answer, top2...)
	}

	return answer
}
