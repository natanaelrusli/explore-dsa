function palindromeChecker(str) {
  for (let i = 0; i < str.length; i++) {
    if (str[i] === " ") {
      str = str.slice(0, i) + str.slice(i + 1);
    }
  }

  if (str.toLowerCase() === str.toLowerCase().split("").reverse().join("")) {
    return true;
  }

  return false;
}

console.log(palindromeChecker("A man a plan a canal Panama"));
