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
    const width = Math.abs(startY - ball[1]);
    return calcDist(height, width);
  }

  //두 공의 y좌표, n - y좌표 중 최소값 구한다
  if (startY === ball[1]) {
    const height = Math.min(startY, n - startY); //삼각형 높이
    const width = Math.abs(startX - ball[0]);
    return calcDist(height, width);
  }

  //두 공의 좌표가 모두 다른 경우
  const candidate = [];
  //4군데 다 생각해야 함
  //남쪽으로 치는 경우
  //동쪽으로 치는 경우
  //북쪽으로 치는 경우
  //서쪽으로 친느 경우

  //하나만 생각하면
  //탄젠트 값 구한다
  //한쪽 삼각형의 높이를 구한다
  //그 높이와 밑변의 길이 통해서 빗변의 길이를 구한다.
  const southXDist = Math.abs(startX - ball[0]);
  const tan1 = southXDist / (startY + ball[1]);
  const h1 = tan1 * startY;
  const t1 = calcDistForSIngleTriangle(h1, startY);
  const h2 = tan1 * ball[1];
  const t2 = calcDistForSIngleTriangle(h2, ball[1]);
  Math.pow(t1 + t2, 2);

  //동쪽
  const eastYDist = Math.abs(startY - ball[1]);
  const tan2 = eastYDist / (m - startX + m - ball[1]);
}

function calcDist(height, width) {
  return Math.pow(Math.sqrt(height * height + width * width) * 2, 2);
}

function calcDistForSIngleTriangle(height, width) {
  return Math.sqrt(height * height + width * width);
}

//문제 분석

//벽에 한번은 맞아야 함
//입사각 반사각 동일하다

//공1 출발

//공2 도착
