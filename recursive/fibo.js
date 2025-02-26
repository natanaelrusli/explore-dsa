function fibo(n) {
  if (n <= 1) {
    return n;
  }
  return fibo(n - 1) + fibo(n - 2);
}

console.log(fibo(5));
console.log(fibo(10));

function fiboMemo(n, memo = {}) {
  if (n <= 1) {
    return n;
  }
  if (n in memo) {
    return memo[n];
  }

  memo[n] = fiboMemo(n - 1, memo) + fiboMemo(n - 2, memo);
  return memo[n];
}

console.log(fiboMemo(5));
console.log(fiboMemo(50));