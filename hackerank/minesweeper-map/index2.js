function createMinesweeperMap(input) {
  const size = 10;

  // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/from
  // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/fill
  const map = Array.from({ length: size }, () => Array(size).fill('.'));
  const directions = [
    [-1, -1],
    [-1, 0],
    [-1, 1],
    [0, -1],
    [0, 1],
    [1, -1],
    [1, 0],
    [1, 1]
  ]

  const lines = input.trim().split('\n');
  const numBombs = parseInt(lines[0]);

  for (let i = 1; i <= numBombs; i++) {
    const [row, col] = lines[i].split(',').map(Number);
    map[row-1][col-1] = 'x';
  }

  console.log(map);
}

// Sample input
const input = `
11
1,8
3,1
4,1
5,1
3,6
6,2
6,3
6,10
8,1
10,5
10,6
`;

console.log(createMinesweeperMap(input));
