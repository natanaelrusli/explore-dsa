// let a = 30;

// if (true) {
//   let b = 40;
//   console.log(a); // Accessible, prints 30
// }

// console.log(b); // Error: b is not defined (only accessible within the block)

// var x = 10;

// function exampleFunction() {
//   var y = 20;
//   console.log(x); // Accessible, prints 10
// }

// // console.log(y); // Error: y is not defined (only accessible within the function)

// const pi = 3.14;
// // pi = 3.14159; // Error: Assignment to a constant variable

const arr = [3, 4, 5, 6];

for (let i in arr) {
  console.log(i);
}

for (var i of arr) {
  console.log(i);
}
