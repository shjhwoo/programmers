const personalities = {
  R: 0,
  T: 0,
  C: 0,
  F: 0,
  J: 0,
  M: 0,
  A: 0,
  N: 0,
};

function solution(survey, choices) {
  let answer = "";
  //성격 유형별로 점수를 합산하는 객체를 하나 만들고

  //설문 배열과 선택한 답지를 정리해서 객체 만들고
  const surveyChoices = {};
  survey.forEach((element, idx) => {
    surveyChoices[element] = choices[idx];
  });
  //이 객체를 순회하면서
  // 성격 유형에다가 점수를 합산한다. (질문지, 선택지 )
  for (const quest in surveyChoices) {
    getPersonalityScore(quest, surveyChoices[quest]);
  }

  const personalitiesArr = Object.entries(personalities);

  personalitiesArr.forEach((personalArr, idx) => {
    if (idx % 2 === 0) {
      const score1 = personalArr[1];
      const score2 = personalitiesArr[idx + 1][1];

      if (score1 >= score2) {
        answer += personalArr[0];
      } else {
        answer += personalitiesArr[idx + 1][0];
      }
    }
  });

  return answer;
}

function getPersonalityScore(quest, choice) {
  console.log(quest, choice, Math.abs(choice - 4));

  if (choice - 4 > 0) {
    personalities[quest[1]] += choice - 4;
  } else if (choice - 4 < 0) {
    personalities[quest[0]] += Math.abs(choice - 4);
  }
}

//console.log(solution(["AN", "CF", "MJ", "RT", "NA"], [5, 3, 2, 7, 5]));

console.log(solution(["TR", "RT", "TR"], [7, 1, 3]));
