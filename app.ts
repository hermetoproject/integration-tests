const yargs = require('yargs/yargs')
const { hideBin } = require('yargs/helpers')
const { format } = require('date-in-spanish')  // fecha.format

const argv = yargs(hideBin(process.argv)).argv

const now = Date.now()

console.log(`Hello, ${argv.name ?? 'World'}!`)
console.log(`Today is ${format(now, 'dddd MMMM Do, YYYY')}`)
