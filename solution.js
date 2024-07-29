function solution(s) {
  /*
  '(' 또는 ')' 로만 이루어진 문자열 
  
  빠르게 판단: 
  길이가 홀수면 무조건 false
  
  짝수일 때가 문제겠네.
  
  스택을 어떻게 이용할까요? 
  
  ()()    
  
  */

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
