function solution(m, n, startX, startY, balls) {
  var answer = [];
  balls.forEach((ball) => {
    const dist = getShortestDist(m, n, startX, startY, ball);
    answer.push(dist);
  });
  return answer;
}

function getShortestDist(m, n, startX, startY, ball) {
  //두 공이 x좌표가 같은 경우
  if (startX === ball[0]) {
    const candidate = [
      Math.pow(
        startX > ball[0] ? n - startY + (n - ball[1]) : startY + ball[1],
        2
      ),
      calcDist2(Math.abs(2 * m - startX - ball[0]), Math.abs(startY - ball[1])), //동
      calcDist2(Math.abs(startX + ball[0]), Math.abs(startY - ball[1])), //서
    ];
    return Math.min(...candidate);
  }

  //두 공의 y좌표 같은 경우, n - y좌표 중 최소값 구한다
  if (startY === ball[1]) {
    const candidate = [
      Math.pow(
        startX > ball[0] ? m - startX + (m - ball[0]) : startX + ball[0],
        2
      ),
      calcDist2(Math.abs(startX - ball[0]), Math.abs(startY - ball[1] * -1)), //남
      calcDist2(Math.abs(startX - ball[0]), Math.abs(2 * n - startY - ball[1])), //북
    ];
    return Math.min(...candidate);
  }

  //두 공의 좌표가 모두 다른 경우
  const candidate = [
    calcDist2(Math.abs(startX - ball[0]), Math.abs(startY - ball[1] * -1)), //남
    calcDist2(Math.abs(2 * m - startX - ball[0]), Math.abs(startY - ball[1])), //동
    calcDist2(Math.abs(startX - ball[0]), Math.abs(2 * n - startY - ball[1])), //북
    calcDist2(Math.abs(startX + ball[0]), Math.abs(startY - ball[1])), //서
  ];
  return Math.min(...candidate);
}

function calcDist2(w, h) {
  return Math.pow(w, 2) + Math.pow(h, 2);
}

solution(10, 10, 3, 7, [
  [7, 7],
  [2, 7],
  [7, 3],
]);
