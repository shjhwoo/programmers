function solution(k, tangerine) {
  let count = 0,
    result = 0,
    object = {};

  tangerine.forEach((x, idx) => {
    object[[x]] = tangerine.filter((element) => x === element).length;
  });

  let gulnum = Object.entries(object).sort(([, a], [, b]) => b - a);

  gulnum.forEach((i, index) => {
    let num = Number(gulnum[index][1]);
    if (num >= k) {
      result++;
    } else if (count < k && k - count >= num) {
      count = count + num;
      result++;
    }
  });

  return result;
}

console.log(solution(1, [1, 1]));
