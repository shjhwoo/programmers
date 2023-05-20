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
  //시작 지점을 구한다(배열)
  let startPoint = getStartPoint(park);

  let newPoint = startPoint;
  // routes를 순회하면서 다음 작업을 한다
  for (let i = 0; i < routes.length; i++) {
    newPoint = getNewPoint(newPoint, routes[i], park);
  }

  //모든 요소를 다 순회했을 떄 정답을 반환한다
  return newPoint;
}

function getStartPoint(park) {
  const row = park.findIndex((ele) => ele.includes("S"));
  const col = park[row].split("").findIndex((el) => el === "S");
  return [row, col];
}

function getNewPoint(current, route, park) {
  const [op, dist] = route.split(" ");

  let newPoint = current;

  switch (op) {
    case "S":
      newPoint = [Number(current[0]) + Number(dist), current[1]];
      break;
    case "N":
      newPoint = [Number(current[0]) - Number(dist), current[1]];
      break;
    case "E":
      newPoint = [current[0], Number(current[1]) + Number(dist)];
      break;
    case "W":
      newPoint = [current[0], Number(current[1]) - Number(dist)];
      break;
  }

  if (!isValidPoint(current, op, newPoint, park)) {
    newPoint = current;
  }

  return newPoint;
}

function isValidPoint(oldPoint, op, newPoint, park) {
  const [oldcol, oldrow] = oldPoint;
  const [col, row] = newPoint;
  const endCol = Math.min(col, park[0].length - 1);
  const endRow = Math.min(row, park.length - 1);

  //도착칸이 장애물이 있거나, 또는 이동 시작점부터, 이동 도착점 사이에는 X가 없어야 한다.
  let noXinWay = false;
  switch (op) {
    case "S":
      //row들을 모두 구한다
      //공원에 있는 아래의 줄들에서 도착점까지 전부다 O여야 한다
      //도착점이 공원을 벗어나면 도착점은 공원의 경계점으로 한다
      noXinWay = park.slice(oldrow + 1, endRow).every((r) => r[oldcol] === "O");
      break;
    case "N":
      noXinWay = park.slice(endRow + 1, oldrow).every((r) => r[oldcol] === "O");
      break;
    case "E":
      noXinWay = park[oldrow]
        .slice(oldcol + 1, endCol)
        .split("")
        .every((c) => c === "O");
      break;
    case "W":
      noXinWay = park[oldrow]
        .slice(endCol + 1, oldcol)
        .split("")
        .every((c) => c === "O");
      break;
  }

  return (
    col >= 0 &&
    row >= 0 &&
    col <= park[0].length - 1 &&
    row <= park.length - 1 &&
    noXinWay
  );
}

// const ans1 = solution(["SOO", "OOO", "OOO"], ["E 2", "S 2", "W 1"]);

const ans2 = solution(["SOO", "OXX", "OOO"], ["E 2", "S 2", "W 1"]);

// const ans3 = solution(["OSO", "OOO", "OXO", "OOO"], ["E 2", "S 3", "W 1"]);

console.log(ans2);
