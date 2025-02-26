/**
 * 
 * @param {*} arr 
 * Given an array of size n-1 with numbers from 1 to n, find the missing number.
 * return 0 if no number is missing.
 */

import { bubbleSort } from "../bubble-sort/index.js";

function findMissingNumber(arr) {
  const n = arr.length - 1;
  bubbleSort(arr);
  
  for (let i = 0; i < n; i++) {
    if (i === 0) continue;
    if (arr[i] - arr[i - 1] !== 1) {
      return arr[i] - 1;
    }  
  }

  return 0;
}

console.log(findMissingNumber([1, 2, 4, 6, 5]));
console.log(findMissingNumber([1, 2, 3, 6, 5]));
console.log(findMissingNumber([6, 1, 2, 7, 3, 4]));