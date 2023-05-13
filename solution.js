function solution(plans) {
  //시작해야 되는 과제
  /*
  멈춰둔 과제 저장소: 스택
  앞으로 해야 하는 과제 저장소: 큐?
  */

  const tempStoppedHomeWork = [];

  //과제 시작 시간 기준 오름차순으로 정렬한다

  //정렬할 때 그냥 시 분을 분으로 통일해서 봐야하나 time 함수로 바꿔서 그런가.
  const notStartedHomeworks = formatPlans(plans).sort(
    (a, b) => a.startTime - b.startTime
  );

  //완료한 과제가 쌓이는 곳
  const completed = [];

  while (notStartedHomeworks.length > 0) {
    const onFocus = notStartedHomeworks[0];
    const { name, startTime, duration, endTime } = onFocus;

    // if (notStartedHomeworks.length > 0) {
    const nextPlan = notStartedHomeworks[1];

    if (nextPlan === undefined && tempStoppedHomeWork.length === 0) {
      return [...completed, onFocus].map((plan) => plan.name);
    }

    const { startTime: nextStartTime } = nextPlan;

    if (nextStartTime < endTime) {
      tempStoppedHomeWork.push({
        ...notStartedHomeworks.shift(),
        duration: duration - (nextStartTime - startTime),
      });
      continue;
    }

    //계획대로 진행중이던 과제 종료
    completed.push(notStartedHomeworks.shift());

    //남은 시간동안 과제를 진행
    /*
          a. 앞으로 해야하는 것, 임시로 그만둔것 모두 있음 => 앞으로 해야할것 부터
          b. 앞으로 해야하는 것만 남은 경우
          c. 임시로 그만둔것만 있음 => 다 몰아서 바로 리턴
        */
    while (tempStoppedHomeWork.length > 0) {
      //다음 과제 시작 시간 전까지 남은 시간, 가장 최근에 멈춘 과제를 하는 데 걸리는 시간
      const focus = tempStoppedHomeWork.pop();
      const { name, startTime, duration } = focus;
      if (nextStartTime - endTime >= duration) {
        completed.push(focus);
      } else {
        //멈춘 과제와 새로 시작할 과제 모두 있음
        tempStoppedHomeWork.push({
          ...focus,
          startTime: endTime,
          duration: duration - (nextStartTime - endTime),
          endTime: endTime + duration - (nextStartTime - endTime),
        });
      }
    }
    //}
  }

  console.log("잠시 멈춘 과제:", tempStoppedHomeWork);

  while (tempStoppedHomeWork.length > 0) {
    completed.push(tempStoppedHomeWork.pop());
  }

  return completed.map((plan) => plan.name);
}

function formatPlans(plans) {
  return plans.map((plan) => {
    const startTime = plan[1];
    const [hh, mm] = startTime.split(":");
    return {
      name: plan[0],
      startTime: Math.floor(new Date(2023, 5, 13, hh, mm) / 1000),
      duration: plan[2] * 60,
      endTime: Math.floor(new Date(2023, 5, 13, hh, mm) / 1000) + plan[2] * 60,
    };
  });
}

// const answer = solution([
//   ["korean", "11:40", "30"],
//   ["english", "12:10", "20"],
//   ["math", "23:30", "40"],
// ]);

// console.log(answer);

const answer2 = solution([
  ["science", "12:40", "50"],
  ["music", "12:20", "40"],
  ["history", "14:00", "30"],
  ["computer", "12:30", "100"],
]);

console.log(answer2);

// const answer3 = solution([
//   ["aaa", "12:00", "20"],
//   ["bbb", "12:10", "30"],
//   ["ccc", "12:40", "10"],
// ]);

// console.log(answer3);
