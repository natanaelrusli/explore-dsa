// TODO: add one more function to rotate the array without using js methods

// The method below is not the most efficient way to rotate an array
// this is generally not recommended for large arrays because unshift can require shifting all elements in the array
// Time complexity is 0(n * k) where n is the length of the array and k is the number of rotations
function rotateArray(arr, k) {
  const n = arr.length;
  k = k % n;

  if (k < 0) {
    k = n + k;
  }

  for (let i = 0; i < k; i++) {
    arr.unshift(arr.pop());
  }
  return arr;
}

function rotateArrayWithSlice(arr, k) {
  const n = arr.length;
  console.log("k", k);
  k = k % n;
  console.log("k after mod", k);
  console.log("n", n);

  if (k < 0) {
    k = n + k;
  }

  return arr.slice(-k).concat(arr.slice(0, -k));
}

// console.log(rotateArray([1, 2, 3, 4, 5], 2));
// console.log(rotateArray([1, 2, 3, 4, 5], 0));
// console.log(rotateArray([1, 2, 3, 4, 5], 5));

console.log(rotateArrayWithSlice([1, 2, 3, 4, 5], 2));
console.log(rotateArrayWithSlice([1, 2, 3, 4, 5], 0));
console.log(rotateArrayWithSlice([1, 2, 3, 4, 5], 5));
console.log(rotateArrayWithSlice([1, 2, 3, 4, 5], 3));
