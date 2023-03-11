function solution(maps) {
  //출발점: S
  //도착점: E
  //레버: L
  const src = getXYPosition(maps, "S");
  const dst = getXYPosition(maps, "E");
  console.log("출발 및 도착지점", src, dst);
  return getShortestDist(maps, src, dst);
}

function getXYPosition(maps, name) {
  const Y = maps.findIndex((road) => road.includes(name));
  const X = maps[Y].indexOf(name);
  return [X, Y];
}

function getShortestDist(maps, src, dst) {
  const up = [-1, 0];
  const right = [0, 1];
  const down = [1, 0];
  const left = [0, -1];
  const direction = [up, right, down, left];

  const isValid = (a, b) =>
    a >= 0 && b >= 0 && a < maps[0].length && b < maps.length;

  //방문여부 저장할 객체
  const visited = {};
  visited[src] = true;
  const queue = [[src, 0]];
  while (queue.length) {
    const [spot, dist] = queue.shift();
    //해당 지점에서 4방향 모두 검사해서 최단거리가 있을지 확인해야 한다
    //포인트는 반복문을 돌면서 원본의 값이 항상 같아야 한다는 것이다.
    for (let d = 0; d < direction.length; d++) {
      const [x, y] = spot;
      const [dx, dy] = direction[d];

      //새롭게 이동한 가상의 정점이 유효한지 확인한다.
      console.log("새로운정점: ", x + dx, y + dy);
      if (
        isValid(x + dx, y + dy) &&
        (maps[x + dx][y + dy] === "O" || maps[x + dx][y + dy] === "L") &&
        visited[`${x + dx},${y + dy}`] === undefined
      ) {
        if (maps[x + dx][y + dy] === "E") return dist + 1;
        visited[`${x + dx},${y + dy}`] = true;
        queue.push([[x + dx, y + dy], dist + 1]);
      }
    }
    return -1;
  }
}

solution(["SOOOL", "XXXXO", "OOOOO", "OXXXX", "OOOOE"]);
