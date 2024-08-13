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
	PlaySum         int
	SongIdxPlaysMap map[int]int
}

func solution(genres []string, plays []int) []int {

	var genreMap = make(map[string]*BestGenRe)

	for idx, genre := range genres {
		if genreMap[genre] == nil {
			genreMap[genre] = &BestGenRe{}
		}
		genreMap[genre].PlaySum += plays[idx]

		if genreMap[genre].SongIdxPlaysMap == nil {
			genreMap[genre].SongIdxPlaysMap = make(map[int]int)
		}

		genreMap[genre].SongIdxPlaysMap[idx] = plays[idx]
	}

	// fmt.Println(genreMap, "genreMap 확인중!")

	var i int
	var top2GenreList = []*BestGenRe{}

	for i < 2 {
		var maxGenrePlayTime int
		var topGenreName string
		var topGenre *BestGenRe
		for genreName, bestGenre := range genreMap {
			if maxGenrePlayTime < bestGenre.PlaySum {
				topGenreName = genreName
				topGenre = bestGenre
				maxGenrePlayTime = bestGenre.PlaySum
			}
		}

		if topGenre != nil {
			top2GenreList = append(top2GenreList, topGenre)
			delete(genreMap, topGenreName)
		}
		i++
	}

	// fmt.Println(top2GenreList, "top2GenreList 확인중!")

	var answer []int
	for _, bestGenre := range top2GenreList {
		var i int
		var listSum [][]int
		for i < 2 {
			var maxPlayTime int
			var topSongIdx int
			for songIdx, playTime := range bestGenre.SongIdxPlaysMap {
				if maxPlayTime < playTime {
					maxPlayTime = playTime
					topSongIdx = songIdx
				}

				if maxPlayTime == playTime {
					if topSongIdx > songIdx {
						topSongIdx = songIdx
					}
				}
			}

			if maxPlayTime > 0 {
				delete(bestGenre.SongIdxPlaysMap, topSongIdx)
				listSum = append(listSum, []int{topSongIdx, maxPlayTime})
			}

			i++
		}

		for _, music := range listSum {
			answer = append(answer, music[0])
		}
	}

	return answer
}
