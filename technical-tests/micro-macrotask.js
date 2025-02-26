console.log("Start");

setTimeout(function () {
  console.log("Timeout 1");
}, 0);

Promise.resolve().then(function () {
  console.log("Promise 1"); // microTask!
});

Promise.resolve().then(function () {
  console.log("Promise 2"); // microTask!
});

Promise.resolve().then(function () {
  console.log("Promise 3"); // microTask!
});

setTimeout(function () {
  console.log("Timeout 2");
}, 0);

console.log("End");
