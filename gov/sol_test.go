package main_test

import (
	"testing"

	"sort"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	gems   []string
	expect []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			gems:   []string{"DIA", "RUBY", "RUBY", "DIA", "DIA", "EMERALD", "SAPPHIRE", "DIA"},
			expect: []int{3, 7},
		},
		{
			gems:   []string{"AA", "AB", "AC", "AA", "AC"},
			expect: []int{1, 3},
		},
		{
			gems:   []string{"XYZ", "XYZ", "XYZ"},
			expect: []int{1, 1},
		},
		{
			gems:   []string{"ZZZ", "YYY", "NNNN", "YYY", "BBB"},
			expect: []int{1, 5},
		},
	}

	for _, test := range tests {
		assert.DeepEqual(t, test.expect, solution(test.gems))
	}
}

//진열된 모든 종류의 보석을
//적어도 1개 이상 포함하는
//가장 짧은 구간을 찾아서 구매

func solution(gems []string) []int {
	gems = append([]string{""}, gems...) //인덱스 헷갈리지 마라공..

	gemMap := make(map[string]bool)
	for _, gem := range gems {
		if gem == "" {
			continue
		}
		gemMap[gem] = true
	}

	startIdx := 1
	endIdx := startIdx + len(gemMap) - 1

	var candidates [][]int

	for startIdx <= endIdx && endIdx < len(gems) {
		shoppingList := gems[startIdx : endIdx+1]

		if shopListHasAllKind(shoppingList, gemMap) {
			candidates = append(candidates, []int{startIdx, endIdx})
			startIdx++
			endIdx = startIdx + len(gemMap) - 1
		} else {
			//인덱스를 이동(endIdx만)
			endIdx++
		}
	}

	// for i, candidate := range candidates {
	// 	fmt.Println(i, "번째 쇼핑리스트 확인: ", gems[candidate[0]:candidate[1]+1], "인덱스: ", candidate)
	// }

	sort.Slice(candidates, func(i, j int) bool {

		iStIdx := candidates[i][0]
		iEndIdx := candidates[i][1]

		jStIdx := candidates[j][0]
		jEndIdx := candidates[j][1]

		if len(gems[iStIdx:iEndIdx+1]) == len(gems[jStIdx:jEndIdx+1]) {
			return iStIdx < jStIdx
		}
		return len(gems[iStIdx:iEndIdx+1]) < len(gems[jStIdx:jEndIdx+1])
	})

	//fmt.Println(candidates)

	return candidates[0]
}

func shopListHasAllKind(shoppingList []string, gemMap map[string]bool) bool {
	shopMap := make(map[string]bool)

	for _, shopGem := range shoppingList {
		shopMap[shopGem] = true
	}

	for gem := range gemMap {
		if !shopMap[gem] {
			return false
		}
	}

	return true
}

/*
javascript:

function solution(gems) {
    gems = ["", ...gems]


    const gemMap = new Map()
    gems.forEach(gem => {
        if (gem === ""){
            return
        }

        gemMap.set(gem, true)
    })

    const gemMapArr = Array.from(gemMap , ([key, value]) => key)

    let start = 1
    let end = start + gemMap.size - 1

    const shopMap = new Map()
    gems.slice(start, end + 1).forEach(gem => {
        if (!shopMap.get(gem)) {
            shopMap.set(gem,1)
        }else{
            shopMap.set(gem, shopMap.get(gem) + 1)
        }
    })


    let hasAll = shopListHasAllKind(shopMap, gemMapArr)
    if (hasAll) {
        candidates.push([start,end])
        start++
        end = start + gemMap.size - 1

        //맵을 조정한다
        const gemToDel = gems[start]
        if (shopMap.get(gemToDel) && shopMap.get(gemToDel) > 0) {
            shopMap.set(gemToDel, shopMap.get(gemToDel) - 1)
        }

        if (shopMap.get(gemToDel) && shopMap.get(gemToDel) == 0) {
            shopMap.delete(gemToDel)
        }
    }else{
        end++
    }

    const newGem = gems[end]
    if (!shopMap.get(newGem)) {
        shopMap.set(newGem, 1)
    }else{
        shopMap.set(newGem, shopMap.get(newGem) + 1)
    }

    const candidates = []
    while (start <= end && end < gems.length) {

        console.log("start: ",start,"end: ",end,"shopMap: ", shopMap)

    }

    candidates.sort(sortCandidates);

    return candidates[0]
}

//이 체크하는 로직을 바꿔야한다. 새로 추가한 보석이 여기에 들어간는지 안들어간는지 보면됨..
function shopListHasAllKind(shopMap, gemMapArr) {
    for (const gem of gemMapArr){
        if (!shopMap.get(gem)){
            return false
        }
    }

    return true
}

function sortCandidates(can1, can2) {
    const iStartIdx = can1[0];
    const iEndIdx = can1[1];

    const jStartIdx = can2[0];
    const jEndIdx = can2[1];

    const iGemListLength = iEndIdx - iStartIdx;
    const jGemListLength = jEndIdx - jStartIdx;

    // First, sort by the length of the range (smaller range comes first)
    if (iGemListLength !== jGemListLength) {
        return iGemListLength - jGemListLength;
    }

    // If the ranges are of equal length, sort by the starting index
    return iStartIdx - jStartIdx;
})
*/
