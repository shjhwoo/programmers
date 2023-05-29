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

  const mbti = {
    R: 0,
    T: 0,
    C: 0,
    F: 0,
    J: 0,
    M: 0,
    A: 0,
    N: 0,
  };

  survey.forEach((q, idx) => {
    const [left, right] = q.split("");
    if (choices[idx] > 4) {
      mbti[right] += choices[idx] - 4;
    } else if (choices[idx] < 4) {
      mbti[left] += 4 - choices[idx];
    }
  });

  answer = answer.concat(mbti["R"] >= mbti["T"] ? "R" : "T");
  answer = answer.concat(mbti["C"] >= mbti["F"] ? "C" : "F");
  answer = answer.concat(mbti["J"] >= mbti["M"] ? "J" : "M");
  answer = answer.concat(mbti["A"] >= mbti["N"] ? "A" : "N");

  return answer;
}
//console.log(solution(["AN", "CF", "MJ", "RT", "NA"], [5, 3, 2, 7, 5]));

console.log(solution(["TR", "RT", "TR"], [7, 1, 3]));
