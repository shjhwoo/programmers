function solution(m, n, startX, startY, balls) {
  var answer = [];
  balls.forEach((ball) => {
    const dist = getShortestDist(m, n, startX, startY, ball);
    answer.push(dist);
  });
  return answer;
}
//이거 점대칭으로 풀어야함.. 그냥 삼각비 쓰면 정수값 안나와서 다 틀림ㅠㅠ
function getShortestDist(m, n, startX, startY, ball) {
  //두 공이 x좌표가 같은 경우
  const [targetX, targetY] = ball;

  const candidate_up =
    (startX - targetX) ** 2 + (startY - 2 * n + targetY) ** 2;
  const candidate_right =
    (startX - 2 * m + targetX) ** 2 + (startY - targetY) ** 2;

  const candidate_down = (startX - targetX) ** 2 + (startY + targetY) ** 2;

  const candidate_left = (startX + targetX) ** 2 + (startY - targetY) ** 2;

  if (startX === targetX) {
    if (startY < targetY) {
      return Math.min(candidate_right, candidate_down, candidate_left);
    } else {
      return Math.min(candidate_right, candidate_up, candidate_left);
    }
  }

  if (startY === targetY) {
    if (startX < targetX) {
      return Math.min(candidate_up, candidate_down, candidate_left);
    } else {
      return Math.min(candidate_up, candidate_down, candidate_right);
    }
  }

  return Math.min(
    candidate_down,
    candidate_left,
    candidate_right,
    candidate_up
  );
}

solution(10, 10, 3, 7, [
  [7, 7],
  [2, 7],
  [7, 3],
]);
