function solution(id_list, report, k) {
  //중복 신고한 결과 제거
  const trimmedReport = Array.from(new Set(report));
  const reportedusers = trimmedReport.map((rep) => rep.split(" ")[1]);

  //유저 - 정지여부를 정리한 객체 선언
  const isBlocked = {};
  id_list.forEach((id) => {
    isBlocked[id] =
      reportedusers.filter((reporteduser) => reporteduser === id).length >= k;
  });
  console.log(isBlocked, "각 유저별로 정지 여부 확인합니다");

  //유저 - 각 유저가 신고한 사람 배열 정리
  const reporter_targetList = {};
  trimmedReport.forEach((rep) => {
    const [reporter, target] = rep.split(" ");
    if (reporter_targetList[reporter] === undefined) {
      reporter_targetList[reporter] = isBlocked[target] ? 1 : 0;
    } else {
      const add = isBlocked[target] ? 1 : 0;
      reporter_targetList[reporter] = reporter_targetList[reporter] + add;
    }
  });

  console.log(reporter_targetList, "각 유저별로 누구 신고했는지 확인합니다");

  return Object.values(reporter_targetList);
}

//목표
/*
내가 다른 사람을 신고했을 떄 정지된 이용자에 대해서만 신고 결과를 받게 되는데
그런 신고 결과를 받은 횟수를 배열에 담아서 리턴한다

한 유저가 같은 유저를 여러 번 신고할 수도 있지만 그 경우는 한번 신고한 것으로 간주하기 떄문에 report 배열을 
셋으로 만들어서 중복을 제거해버리기.

*/
