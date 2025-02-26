// function asyncExample1() {
//   return new Promise((resolve, reject) => {
//     setTimeout(() => {
//       resolve("Async Example 1");
//     }, 1000);
//   });
// }

// asyncExample1().then((result) => console.log(result));

// function asyncExample2() {
//   return new Promise((resolve, reject) => {
//     setTimeout(() => {
//       reject("Error in Async Example 2");
//     }, 500);
//   });
// }

// asyncExample2().catch((error) => console.error(error));

// console.log("first");

// function asyncExample3() {
//   return new Promise((resolve) => {
//     console.log("Async Example 3");
//     resolve();
//   });
// }

// asyncExample3().then(() => console.log("After Async Example 3"));

// console.log("Start");

// setTimeout(function () {
//   console.log("Timeout 2");
// }, 0);

// Promise.resolve(console.log("resolved")).then(function () {
//   setTimeout(function () {
//     console.log("Timeout 1");
//   }, 0);
//   console.log("wow");
// });

// console.log("End");

const promise = new Promise((res) => res(2));

promise
  .then((v) => {
    console.log(v); // Output: 2
    return v * 2;
  })
  .then((v) => {
    console.log(v); // Output: 4
    return v * 2;
  })
  .finally((v) => {
    console.log(v); // Output: undefined
    return 0;
  })
  .then((v) => {
    console.log(v); // Output: 0
  });
