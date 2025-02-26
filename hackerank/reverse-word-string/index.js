function reverseWordString(str) {
  return str.split(' ').reverse().join(' ');
}

console.log(reverseWordString('Hello World')); // World Hello
console.log(reverseWordString("The quick brown fox jumps over the lazy dog")); // dog lazy the over jumps fox brown quick The
