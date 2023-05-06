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
        completed.push(onFocus.shift());

        //새로 시작할 과제와 잠시 멈춘 과제가 모두 있는 경우
        if (notStartedHomeworks.length > 0 && tempStoppedHomeWork.length > 0) {
          onFocus.push(notStartedHomeworks.shift());
        }

        //잠시 멈춘 과제만 있는 경우
        if (
          notStartedHomeworks.length === 0 &&
          tempStoppedHomeWork.length > 0
        ) {
          onFocus.push({
            ...tempStoppedHomeWork.pop(),
            startTime: newHomeWorkEndTIme,
          });
        }

        //새로 시작할 과제만 있는 경우
        if (
          notStartedHomeworks.length > 0 &&
          tempStoppedHomeWork.length === 0
        ) {
          onFocus.push(notStartedHomeworks.shift());
        }
      } else {
        /*
          단순히 특정 과제 단독으로 시작시간, 걸리는시간을 보면 안되고 다음 과제를 시작하게 될 때 까지 걸리는 시간이 소모되는 걸 생각해야함
          테스트3개는 통과인데 제출에서 거의 반타작인 경우 남은시간이 차감되는걸 신경쓰지 않았을 확률이 큼
        */
        //새 숙제부터 시작해야해
        const { name, startTime, duration, endTime } = onFocus[0];
        const remainedDuration =
          duration - (newHomeWorkStartTIme - currentHomeWorkStartTime);
        tempStoppedHomeWork.push({
          ...onFocus.shift(),
          duration: remainedDuration,
        });
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
      duration: plan[2],
      endTime: new Date(
        new Date(null, null, null, hh, mm).getTime() + plan[2] * 60000
      ),
    };
  });
}

const answer = solution([
  ["korean", "11:40", "30"],
  ["english", "12:10", "9999999999999"],
  ["math", "12:30", "40"],
]);

console.log(answer);

const answer2 = solution([
  ["science", "12:40", "50"],
  ["music", "12:20", "40"],
  ["history", "14:00", "99999999999999999999"],
  ["computer", "12:30", "100"],
]);

console.log(answer2);

const answer3 = solution([
  ["aaa", "12:00", "20"],
  ["bbb", "12:10", "30"],
  ["ccc", "12:40", "10"],
]);

console.log(answer3);
