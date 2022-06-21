const prompt = require('prompt-sync')({fake_val: "OPTIONAL CONFIG VALUES HERE",});
const num1 = parseInt(prompt('Enter the first number '));
const num2 = parseInt(prompt('Enter the second number '));

const sum = num1 + num2;

console.log(`The sum of ${num1} and ${num2} is ${sum}`);
