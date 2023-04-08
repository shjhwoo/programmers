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

/*
In general, a JavaScript Map is faster than a JavaScript Object for finding a value
 because Maps are specifically designed for key-value pairs 
 and have faster lookup times than Objects.

Maps use a hash table under the hood to store key-value pairs 
and can provide constant time complexity (O(1)) for finding a value by key. 
On the other hand, Objects use a hash table as well, 
but they have additional features like prototype chaining that can slow down lookups.
*/

function solution(player, calling) {
  const map = new Map();
  const map2 = new Map();
  player.forEach((e, i) => {
    map.set(e, i);
  });
  player.forEach((e, i) => {
    map2.set(i, e);
  });

  for (let i = 0; i < calling.length; i++) {
    if (map.get(calling[i]) !== undefined) {
      const playerName = calling[i];
      const playerScore = map.get(playerName);
      const playerNewScore = playerScore - 1;
      const OtherName = map2.get(playerScore - 1);
      const OtherScore = playerScore - 1;
      const OtherNewScore = playerScore;

      map.set(playerName, playerNewScore);
      map.set(OtherName, OtherNewScore);
      map2.set(playerScore, OtherName);
      map2.set(OtherScore, playerName);
    }
  }
  let arr = [];
  for (let [key, val] of map.entries()) {
    arr[val] = key;
  }
  return arr;
}
