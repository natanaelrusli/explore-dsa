const input = [1, 3, 6, 4, 1, 2];

function solution(A) {
  const uniquePositiveNumber = [...new Set(A.filter((x) => x > 0))];

  if (uniquePositiveNumber.length === 0) {
    return 1;
  }

  let smallestPositiveNumber = 1;
  while (uniquePositiveNumber.includes(smallestPositiveNumber)) {
    smallestPositiveNumber += 1;
  }

  return smallestPositiveNumber;
}

console.log(solution(input));
