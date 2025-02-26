function findSecondLargest(arr) {
  let largest = 0;
  let secondLargest = 0;

  for (val of arr) {
    if (val >largest) {
      secondLargest = largest;
      largest = val;
    }
  }

  return secondLargest;
}

console.log(findSecondLargest([4, 2, 7, 1, 9])); 
console.log(findSecondLargest([1, 2, 3, 4])); 
