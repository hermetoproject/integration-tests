const lodash = require('lodash');

let numbers = [1, 2, 3, 4, 5];
let shuffledNumbers = lodash.shuffle(numbers);

console.log("Hello from first-pkg!");

console.log('Original numbers:', numbers);
console.log('Shuffled numbers:', shuffledNumbers);
