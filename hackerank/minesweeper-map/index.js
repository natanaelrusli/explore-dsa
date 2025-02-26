function createMinesweeperMap(input) {
    const size = 10;
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
    ];

    const lines = input.trim().split('\n');
    const numBombs = parseInt(lines[0]);

    for (let i = 1; i <= numBombs; i++) {
        const [row, col] = lines[i].split(',').map(Number);
        map[row - 1][col - 1] = 'x';
    }

    for (let row = 0; row < size; row++) {
        for (let col = 0; col < size; col++) {
            if (map[row][col] === 'x') continue;

            let bombCount = 0;
            for (const [dr, dc] of directions) {
                const newRow = row + dr;
                const newCol = col + dc;
                if (newRow >= 0 && newRow < size && newCol >= 0 && newCol < size && map[newRow][newCol] === 'x') {
                    bombCount++;
                }
            }

            if (bombCount > 0) {
                map[row][col] = bombCount.toString();
            }
        }
    }

    return map.map(row => row.join('')).join('\n');
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
