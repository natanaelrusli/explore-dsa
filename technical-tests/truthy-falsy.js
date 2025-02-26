var string = "string";
var number = 0;
var bool = false;

console.log(number || string); // Output: "string"
console.log(number && string); // Output: 0
console.log(bool || string); // Output: true
console.log(bool && string); // Output: "string"

// The || (logical OR) operator returns the first truthy operand. If the first operand is falsy, it returns the second operand.
// The && (logical AND) operator returns the first falsy operand. If the first operand is truthy, it returns the second operand.
