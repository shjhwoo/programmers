function solution(players, callings) {
  let result = Object.fromEntries(players.map((player, idx) => [player, idx]));
  let res2 = Object.fromEntries(players.map((player, idx) => [idx, player]));

  callings.forEach((player, i) => {
    const idx = result[player];
    const prevIdx = idx - 1;

    const prevRes2 = { ...res2 };
    res2[prevIdx] = player;
    res2[idx] = prevRes2[prevIdx];

    const prevResult = { ...result };
    result[player] = prevResult[player] - 1;
    result[res2[idx]] = prevResult[res2[idx]] + 1;

    console.log(result, res2, i);
  });

  return Object.values(res2);
}

const res = solution(
  ["mumu", "soe", "poe", "kai", "mine"],
  ["kai", "kai", "mine", "mine"]
);
console.log(res);
