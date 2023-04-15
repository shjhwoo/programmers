function solution(name, yearning, photo) {
  //name과 yearning을 맵으로 만들어준다. 이름 => 점수
  const scoreMap = new Map();
  name.forEach((n, idx) => {
    scoreMap.set(n, yearning[idx]);
  });
  //photo 순회하면서 점수 합산한다
  return photo
    .map((p) => p.filter((n) => scoreMap.get(n)))
    .map((p) => getYearningScore(p, scoreMap));
}

function getYearningScore(photo, scoreMap) {
  return photo.reduce((sum, n) => {
    return scoreMap.get(n) !== undefined ? scoreMap.get(n) + sum : sum;
  }, 0);
}
