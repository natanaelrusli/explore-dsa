// Write a function to generate the first n numbers of the Fibonacci sequence.

/**
 * 
 * @param {*} n 
 * @returns fibonacci sequence
 * 
 * the big O notation for this function is O(n)
 */
function fibonacci(n) {
  // can assign the first two numbers of the sequence to 0 and 1
  // because fibonacci sequence always starts with 0 and 1
  let fib = [];
  for (let i = 2; i < n; i++) {
    if (fib.length === 0) {
      fib.push(0);
      fib.push(1);
    }
    fib.push(fib[i - 1] + fib[i - 2]);
  }

  return fib;
}

console.log(fibonacci(10)); // [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
