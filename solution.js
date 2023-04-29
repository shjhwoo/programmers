function solution(plans) {
  //시작해야 되는 과제
  /*
  멈춰둔 과제 저장소: 스택
  앞으로 해야 하는 과제 저장소: 큐?
  */

  const tempStoppedHomeWork = [];

  //과제 시작 시간 기준 오름차순으로 정렬한다
  const notStartedHomeworks = formatPlans(plans).sort(
    (a, b) => a.startTime - b.startTime
  );

  //현재 진행중인 과제
  const onFocus = [];

  //완료한 과제가 쌓이는 곳
  const completed = [];

  while (notStartedHomeworks.length > 0 || tempStoppedHomeWork.length > 0) {
    if (onFocus.length === 0) {
      //맨 처음인 경우
      onFocus.push(notStartedHomeworks.shift());
    } else {
      const {
        startTime: currentHomeWorkStartTime,
        endTime: currentHomeWorkEndTime,
      } = onFocus[0];

      if (notStartedHomeworks.length === 0 && tempStoppedHomeWork.length > 0) {
        return [...completed, ...onFocus, ...tempStoppedHomeWork.reverse()].map(
          (plan) => plan.name
        );
      }

      const { startTime: newHomeWorkStartTIme, endTime: newHomeWorkEndTIme } =
        notStartedHomeworks[0];

      if (currentHomeWorkEndTime <= newHomeWorkStartTIme) {
        //새로 시작할 과제와 잠시 멈춘 과제가 모두 있는 경우
        if (notStartedHomeworks.length > 0 && tempStoppedHomeWork.length > 0) {
          completed.push(onFocus.shift());
          onFocus.push(notStartedHomeworks.shift());
        }

        //잠시 멈춘 과제만 있는 경우
        if (
          notStartedHomeworks.length === 0 &&
          tempStoppedHomeWork.length > 0
        ) {
          completed.push(onFocus.shift());
          onFocus.push({
            ...tempStoppedHomeWork.pop(),
            startTime: newHomeWorkEndTIme,
          });
        }

        //새로 시작할 과제만 있는 경우
        //정상적으로 끝남
        if (
          notStartedHomeworks.length > 0 &&
          tempStoppedHomeWork.length === 0
        ) {
          completed.push(onFocus.shift());
          onFocus.push(notStartedHomeworks.shift());
        }
      } else {
        //새 숙제부터 시작해야해
        const { name, startTime, duration, endTime } = onFocus[0];
        const newDuration =
          duration - (newHomeWorkStartTIme - currentHomeWorkStartTime);
        tempStoppedHomeWork.push({ ...onFocus.shift(), duration: newDuration });
        onFocus.push(notStartedHomeworks.shift());
      }
    }
  }

  if (notStartedHomeworks.length === 0 && tempStoppedHomeWork.length === 0) {
    return [...completed, ...onFocus].map((plan) => plan.name);
  }

  return completed.map((plan) => plan.name);
}

function formatPlans(plans) {
  return plans.map((plan) => {
    const startTime = plan[1];
    const [hh, mm] = startTime.split(":");
    return {
      name: plan[0],
      startTime: new Date(null, null, null, hh, mm),
      durations: plan[2],
      endTime: new Date(
        new Date(null, null, null, hh, mm).getTime() + plan[2] * 60000
      ),
    };
  });
}

const answer = solution([
  ["korean", "11:40", "30"],
  ["english", "12:10", "20"],
  ["math", "12:30", "40"],
]);

console.log(answer);

const answer2 = solution([
  ["science", "12:40", "50"],
  ["music", "12:20", "40"],
  ["history", "14:00", "30"],
  ["computer", "12:30", "100"],
]);

console.log(answer2);

const answer3 = solution([
  ["aaa", "12:00", "20"],
  ["bbb", "12:10", "30"],
  ["ccc", "12:40", "10"],
]);

console.log(answer3);
