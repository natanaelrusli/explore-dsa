/**
 * 
 * @param {*} a array that will be rotated 
 * @param {*} d the amount of left rotation
 * @returns an array resulted from the operation
 */
function rotLeft(a, d) {
  if (!a.length) return [];

  for (let i = 0; i < d; i++) {
      a.push(a.splice(0, 1)[0]);
  }
  
  return a;
}

console.log(rotLeft([1, 2, 3, 4, 5], 4));
console.log(rotLeft([], 4));
