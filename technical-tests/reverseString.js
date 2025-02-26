function getEndSymbol(inputStr) {
  let i = inputStr.length - 1;
  // let regex = /^[a-zA-Z]+$/;
  let result = "";

  while (
    !(inputStr[i].toLowerCase() >= "a" && inputStr[i].toLowerCase() <= "z") &&
    i >= 0
  ) {
    result += inputStr[i];
    i--;
  }

  return { result, updatedString: inputStr.slice(0, i + 1) };
}

let inputStr = "is, golang, what.,.,";
let { result: symbol, updatedString } = getEndSymbol(inputStr);

let updatedStrArray = updatedString.split(", ");

let result = updatedStrArray.reverse().join(" ") + symbol;
console.log(result);
