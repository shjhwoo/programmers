function solution(plans) {
  const answer = [];

  //answer <== 진행 중 과제 잠시멈춘과제(스택)

  //제일 먼저 plans 들을 시작시간 순서로 오름차순 정렬해야 한다
  const formattedPlans = formatPlans(plans);
  //잠시 멈춘 과제 배열을 선언한다
  const tempStopped = new Map(); //{name, remain, idx} 멈춘 과제 이름과 남은 시간, 인덱스를 배열로 저장한다

  //정렬된 과제 배열을 순회한다.
  formattedPlans.forEach((plan, idx) => {
    //지금 시작한 과제
    const [name, startTime, requiredTime] = plan;
    //지금 시작한 과제가 끝나는 시각을 구한다
    const endTime = new Date(startTime.getTime() + requiredTime * 60000);

    if (formattedPlans[idx + 1] !== undefined) {
      //다음에 시작해야 하는 과제
      const [nextName, nextStartTime, nextRequiredTime] =
        formattedPlans[idx + 1];

      if (endTime > nextStartTime) {
        if (!tempStopped.has(name)) {
          tempStopped.set(name, {
            idx,
            remain: requiredTime - getMinDiff(nextStartTime - startTime),
          });
        } else {
          //이미 배열에 있는 경우, 시간 다시 계산해준다
          const idx = tempStopped.get(name).idx;
          const remain = tempStopped.get(name).remain;
          tempStopped.set(name, {
            idx,
            remain: remain - getMinDiff(nextStartTime - startTime),
          });
        }
      }
      //제시간에 끝낸경우
      if (endTime < nextStartTime) {
        answer.push(name);
        //잠시 멈춰둔 과제가 있다면 이어서 진행
        if (tempStopped.size > 0) {
          const recentStoppped = Array.from(tempStopped).pop();
          console.log(recentStoppped);
        }
      }

      if (endTime === nextStartTime) {
        answer.push(name);
        //잠시 멈춰둔 과제가 있다면 새로 시작하는 과제부터 진행
        if (tempStopped.size > 0) {
        }
      }
    }
  });

  return answer;
}

function formatPlans(plans) {
  return plans.map((plan) => {
    const startTime = plan[1];
    const [hh, mm] = startTime.split(":");
    return [plan[0], new Date(null, null, null, hh, mm), plan[2]];
  });
}

function getMinDiff(nextDate, date) {
  return (((nextDate - date) % 86400000) % 3600000) / 60000;
}
