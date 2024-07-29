function solution(s) {
  if (s.length % 2 === 1) {
    return false;
  }

  let pstack = [];
  for (const char of s) {
    pstack.push(char);

    if (pstack.length < 2) {
      continue;
    }

    const lastTwoChar = pstack.slice(pstack.length - 2, pstack.length);
    if (lastTwoChar[0] === "(" && lastTwoChar[1] === ")") {
      pstack = pstack.slice(0, pstack.length - 2);
    }
  }

  console.log(pstack);

  return pstack.length === 0;
}

console.log(solution("()()"));
