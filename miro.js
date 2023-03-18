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
    const height = Math.min(startX, m - startX); //삼각형 높이
    const width = Math.abs(startY - ball[1]) / 2;
    return Math.round(calcDist(height, width));
  }

  //두 공의 y좌표, n - y좌표 중 최소값 구한다
  if (startY === ball[1]) {
    const height = Math.min(startY, n - startY); //삼각형 높이
    const width = Math.abs(startX - ball[0]) / 2;
    return Math.round(calcDist(height, width));
  }

  //두 공의 좌표가 모두 다른 경우
  const candidate = [
    calcDist2(Math.abs(startX - ball[0]), startY, ball[1]), //남
    calcDist2(Math.abs(startY - ball[1]), m - startX, m - ball[0]), //동
    calcDist2(Math.abs(startX - ball[0]), n - startY, n - ball[1]), //북
    calcDist2(Math.abs(startY - ball[1]), startX, ball[0]), //서
  ];
  return Math.round(Math.min(...candidate));
}

function calcDist(height, width) {
  return Math.pow(Math.sqrt(height * height + width * width) * 2, 2);
}

function calcDistForSIngleTriangle(height, width) {
  return Math.sqrt(height * height + width * width);
}

function calcDist2(widthDist, theight1, theight2) {
  const tan2 = widthDist / (theight1 + theight2);
  const h1 = tan2 * theight1;
  const t1 = calcDistForSIngleTriangle(h1, theight1);
  const h2 = tan2 * theight2;
  const t2 = calcDistForSIngleTriangle(h2, theight2);
  return Math.pow(t1 + t2, 2);
}
