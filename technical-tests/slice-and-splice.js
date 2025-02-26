// slice
const originalArray1 = [1, 2, 3, 4, 5];

const slicedArray = originalArray1.slice(1, 4);

console.log(originalArray1);
console.log(slicedArray);

// splice
const originalArray2 = [1, 2, 3, 4, 5];

// Using splice to replace elements in the array
const removedItems = originalArray2.splice(1, 2, 1);

console.log("Original Array 2 after splice:", originalArray2);
console.log("Removed items from Original Array 2:", removedItems);
