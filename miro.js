//최소한의 이동거리를 구하려면 가장 경계점에 있는 파일이 무엇인지를 알아야 한다
//가장 위 - 가장 왼쪽에 있는 파일 하나.
//가장 아래- 가장 오른쪽에 있는 파일 하나.

//특이 케이스: 파일이 하나뿐인 경우: 바로 처리할 수 있다

function solution(wallpaper) {
  //파일이 단 하나뿐인 경우
  if (wallpaper.join("").match(/#/g)?.length === 1) {
    //파일의 XY 좌표 구하기
    const fileY = wallpaper.findIndex((line) => line.includes("#"));
    const fileX = wallpaper[fileY].split("").findIndex((ch) => ch === "#");

    return [fileY, fileX, fileY + 1, fileX + 1];
  }
  //가장 위에파일 찾기
  const fileLY = wallpaper.findIndex((line) => line.includes("#"));
  //가장 왼쪽에 파일 찾기
  const fileLX = Math.min(
    ...wallpaper
      .map((line) => line.split("").findIndex((ch) => ch === "#"))
      .filter((n) => n >= 0)
  );

  //왼쪽 위 교점 구하기

  //가장 아래 파일 찾기
  const fileRY =
    wallpaper.length -
    wallpaper.reverse().findIndex((line) => line.includes("#"));

  //가장 오른쪽 파일 찾기
  const fileRX = Math.max(
    ...wallpaper
      .filter((line) => line.includes("#"))
      .map(
        (line) =>
          line.length -
          line
            .split("")
            .reverse()
            .findIndex((ch) => ch === "#")
      )
      .filter((n) => n >= 0)
  );
  return [fileLY, fileLX, fileRY, fileRX + 1];
}

solution([
  ".##...##.",
  "#..#.#..#",
  "#...#...#",
  ".#.....#.",
  "..#...#..",
  "...#.#...",
  "....#....",
]);
//y좌표, x좌표 순으로 리턴
//파일 위치가 x: 0 y: 1인 경우

//왼쪽 위: x: 0 y: 1
//오른쪽 위: x: 1 y: 1

//왼쪽 아래: x: 0 y: 2
//오른쪽 아래: x: 1 y: 2

//파일 위치가 x1 y1인 경우
//왼쪽 위: x1 y1
//오른쪽 위: x2 y1
//왼쪽 아래: x1 y2
//오른쪽 아래: x2 y2
