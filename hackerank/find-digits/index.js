/**
 * 
 * @param {*} n 
 * https://www.hackerrank.com/challenges/find-digits/problem?isFullScreen=true
 */
function findDigits(n) {
  // split and get each digit
  let digits = n.toString().split("");
  let res = 0;

  // use modulo to check if the digit can be evenly divided into n
  for (const digit of digits) {
    if (n % parseInt(digit) === 0) res++;
  }
  
  // count how many digits of n can be evenly divided and return the count integer
  return res;
}

console.log(findDigits(12));
console.log(findDigits(1012));