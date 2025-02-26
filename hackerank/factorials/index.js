// or this function can be declared inside of the extraLongFactorials function
function factorial(num) {
    if (num === 1) return BigInt(1);
    return BigInt(num) * factorial(num - 1);
}

function extraLongFactorials(n) {
  const result = factorial(n);
  console.log(result.toString());
}

extraLongFactorials(3);
