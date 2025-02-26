/**
    Input: 42145 Output: 54421
    Input: 145263 Output: 654321
    Input: 123456789 Output: 987654321
*/

export function descendingOrder(n: number): number {
  // your code here
  const nString = n.toString();

  for (let i = 0; i < nString.length; i++) {
    console.log(nString[i]);
  }

  return parseInt(nString);
}

descendingOrder(123);
