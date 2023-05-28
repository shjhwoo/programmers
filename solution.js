// function solution(plans) {
//   //시작해야 되는 과제
//   /*
//   멈춰둔 과제 저장소: 스택
//   앞으로 해야 하는 과제 저장소: 큐?
//   */

//   const tempStoppedHomeWork = [];

//   //과제 시작 시간 기준 오름차순으로 정렬한다
//   const notStartedHomeworks = formatPlans(plans).sort(
//     (a, b) => a.startTime - b.startTime
//   );

//   //현재 진행중인 과제
//   const onFocus = [];

//   //완료한 과제가 쌓이는 곳
//   const completed = [];

//   while (notStartedHomeworks.length > 0 || tempStoppedHomeWork.length > 0) {
//     if (onFocus.length === 0) {
//       //맨 처음인 경우
//       onFocus.push(notStartedHomeworks.shift());
//     } else {
//       const {
//         startTime: currentHomeWorkStartTime,
//         endTime: currentHomeWorkEndTime,
//       } = onFocus[0];

//       if (notStartedHomeworks.length === 0 && tempStoppedHomeWork.length > 0) {
//         return [...completed, ...onFocus, ...tempStoppedHomeWork.reverse()].map(
//           (plan) => plan.name
//         );
//       }

//       const { startTime: newHomeWorkStartTIme, endTime: newHomeWorkEndTIme } =
//         notStartedHomeworks[0];

//       if (currentHomeWorkEndTime <= newHomeWorkStartTIme) {
//         //새로 시작할 과제와 잠시 멈춘 과제가 모두 있는 경우
//         if (notStartedHomeworks.length > 0 && tempStoppedHomeWork.length > 0) {
//           completed.push(onFocus.shift());
//           onFocus.push(notStartedHomeworks.shift());
//         }

//         //잠시 멈춘 과제만 있는 경우
//         if (
//           notStartedHomeworks.length === 0 &&
//           tempStoppedHomeWork.length > 0
//         ) {
//           completed.push(onFocus.shift());
//           onFocus.push({
//             ...tempStoppedHomeWork.pop(),
//             startTime: newHomeWorkEndTIme,
//           });
//         }

//         //새로 시작할 과제만 있는 경우
//         //정상적으로 끝남
//         if (
//           notStartedHomeworks.length > 0 &&
//           tempStoppedHomeWork.length === 0
//         ) {
//           completed.push(onFocus.shift());
//           onFocus.push(notStartedHomeworks.shift());
//         }
//       } else {
//         //새 숙제부터 시작해야해
//         const { name, startTime, duration, endTime } = onFocus[0];
//         const newDuration =
//           duration - (newHomeWorkStartTIme - currentHomeWorkStartTime);
//         tempStoppedHomeWork.push({ ...onFocus.shift(), duration: newDuration });
//         onFocus.push(notStartedHomeworks.shift());
//       }
//     }
//   }

//   if (notStartedHomeworks.length === 0 && tempStoppedHomeWork.length === 0) {
//     return [...completed, ...onFocus].map((plan) => plan.name);
//   }

//   return completed.map((plan) => plan.name);
// }

// function formatPlans(plans) {
//   return plans.map((plan) => {
//     const startTime = plan[1];
//     const [hh, mm] = startTime.split(":");
//     return {
//       name: plan[0],
//       startTime: new Date(2023, 5, 14, hh, mm),
//       durations: plan[2],
//       endTime: new Date(
//         new Date(2023, 5, 14, hh, mm).getTime() + plan[2] * 60000
//       ),
//     };
//   });
// }

// const answer = solution([
//   ["korean", "11:40", "30"],
//   ["english", "12:10", "20"],
//   ["math", "12:30", "40"],
// ]);

// console.log(answer);

// const answer2 = solution([
//   ["science", "12:40", "50"],
//   ["music", "12:20", "40"],
//   ["history", "14:00", "30"],
//   ["computer", "12:30", "100"],
// ]);

// console.log(answer2);

// const answer3 = solution([
//   ["aaa", "12:00", "20"],
//   ["bbb", "12:10", "30"],
//   ["ccc", "12:40", "10"],
// ]);

// console.log(answer3);

function solution(park, routes) {
  //시작 지점 구하기
  let [row, col] = getStartPoint(park);

  for (let i = 0; i < routes.length; i++) {
    const tempPoint = getNewDestination([row, col], routes[i]);
    if (isInPark(tempPoint, park)) {
      if (NoHuddleInWay([row, col], tempPoint, park)) {
        row = tempPoint[0];
        col = tempPoint[1];
      }
    }
  }
  return [row, col];
}

function getStartPoint(park) {
  const row = park.findIndex((route) => route.includes("S"));
  const col = park[row].split("").findIndex((el) => el === "S");
  return [row, col];
}

function getNewDestination(currentPoint, command) {
  const [curRow, curCol] = currentPoint;

  const [op, n] = command.split(" ");

  switch (op) {
    case "E":
      return [curRow, Number(curCol) + Number(n)];
    case "W":
      return [curRow, Number(curCol) - Number(n)];
    case "S":
      return [Number(curRow) + Number(n), curCol];
    case "N":
      return [Number(curRow) - Number(n), curCol];
  }
}

function isInPark(tempPoint, park) {
  const [tempRow, tempCol] = tempPoint;
  return (
    tempRow >= 0 &&
    tempCol >= 0 &&
    tempRow <= park.length - 1 &&
    tempCol <= park[0].length - 1
  );
} //공원을 안 벗어났는지

function NoHuddleInWay(oldPoint, newPoint, park) {
  const [oldRow, oldCol] = oldPoint;
  const [newRow, newCol] = newPoint;
  // 같은 가로선상에 있는 경우: 공원의 그 줄 슬라이스 해서 X가 없으면 된다
  if (oldRow === newRow) {
    return !park[oldRow]
      .slice(Math.min(oldCol, newCol), Math.max(oldCol, newCol) + 1)
      .includes("X");
  }
  //같은 세로선상에 있는 경우: 공원의 시작로우부터 끝로우 사이에 있는 콜이 모두 X가 아니어야해
  if (oldCol === newCol) {
    return park
      .slice(Math.min(oldRow, newRow), Math.max(oldRow, newRow) + 1)
      .every((route) => route[oldCol] !== "X");
  }
} //진행로에 장애물이 없는지

const test1 =
  solution(["SOO", "OOO", "OOO"], ["E 2", "S 2", "W 1"]).toString() ==
  [2, 1].toString();

const test2 =
  solution(["SOO", "OXX", "OOO"], ["E 2", "S 2", "W 1"]).toString() ===
  [0, 1].toString();

const test3 =
  solution(["OSO", "OOO", "OXO", "OOO"], ["E 2", "S 3", "W 1"]).toString() ==
  [0, 0].toString();

console.log(test1, test2, test3);
