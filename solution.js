function solution(s) {
  if (s.length % 2 === 1) {
    return false;
  }

  if (s[0] === ")" || s[s.length - 1] === "(") {
    return false;
  }

  let pstack = [];
  for (const char of s) {
    if (pstack.length === 0 && char === ")") {
      return false;
    }

    pstack.push(char);

    if (pstack.length < 2) {
      continue;
    }

    const lastTwoChar = pstack.slice(pstack.length - 2, pstack.length);
    if (lastTwoChar[0] === "(" && lastTwoChar[1] === ")") {
      pstack = pstack.slice(0, pstack.length - 2);
    }
  }

  return pstack.length === 0;
}
console.log(solution("()()"));

//이게 더 좋다고 함.
function solution(s) {
  let answer = [];
  for (i of s) {
    if (i == "(") answer.push("(");
    else if (answer.length == 0) return false; //중간에 즉시 리턴을 위함,
    else answer.pop();
  }
  return answer.length ? false : true;
}

//스택 말고 메모리 더 줄여서.
function solution2(s) {
  let cnt = 0;
  for (const c of s) {
    if (c === ")") {
      cnt++;
    } else {
      if (cnt === 0) {
        //더 이상 빼줄 수 있는,
        //즉 이전에 들어간 짝 ( 괄호가 없기 때문에 false를 리턴하는것이다.
        return false;
      }
      cnt--;
    }
  }

  return cnt === 0;
}
