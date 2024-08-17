class MaxHeap {
  constructor() {
    this.heap = [null]; //0일때 예외처리를 하기 힘들어서, 일부러 최상위 노드는 null로 채운거구나..
  }

  len() {
    return this.heap.length;
  }

  push(value) {
    this.heap.push(value);

    let currentIdx = this.heap.length - 1;
    let parentIdx = Math.floor(currentIdx / 2); //만약에 초기가 아얘 빈 배열이었으면 current가 바로 parent가 되어버리니 그 부분 처리하는 게 귀찮아 질 수 있겠구나.

    while (parentIdx > 0 && this.heap[parentIdx] < value) {
      let temp = this.heap[parentIdx];
      this.heap[parentIdx] = value;
      this.heap[currentIdx] = temp;

      currentIdx = parentIdx;
      parentIdx = Math.floor(currentIdx / 2);
    }
  }

  pop() {
    if (this.heap.length === 2) return this.heap.pop();

    const returnValue = this.heap[1];
    this.heap[1] = this.heap.pop();

    let currentIdx = 1;
    let leftIdx = 2;
    let rightIdx = 3;

    while (
      this.heap[currentIdx] < this.heap[leftIdx] ||
      this.heap[currentIdx] < this.heap[rightIdx]
    ) {
      if (this.heap[leftIdx] < this.heap[rightIdx]) {
        const temp = this.heap[currentIdx];
        this.heap[currentIdx] = this.heap[rightIdx];
        this.heap[rightIdx] = temp;
        currentIdx = rightIdx;
      } else {
        const temp = this.heap[currentIdx];
        this.heap[currentIdx] = this.heap[leftIdx];
        this.heap[leftIdx] = temp;
        currentIdx = leftIdx;
      }

      leftIdx = currentIdx * 2;
      rightIdx = currentIdx * 2 + 1;
    }

    return returnValue;
  }
}

function solution(no, works) {
  if (works.reduce((a, b) => a + b) <= no) {
    return 0;
  }

  const heap = new MaxHeap();
  for (const work of works) {
    heap.push(work);
  }

  while (heap.len() > 0 && no > 0) {
    heap.push(heap.pop() - 1);
    no--;
  }

  return heap.heap.reduce((a, b) => a + b * b);
}

console.log(solution(4, [4, 3, 3]));
