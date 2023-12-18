import * as fs from 'node:fs/promises'
import * as console from "node:console";
import yargs from "yargs";
import {com, Timer} from "../../com";

const timer = new Timer()

const digits = 'one, two, three, four, five, six, seven, eight, nine, 1, 2, 3, 4, 5, 6, 7, 8, 9'
    .split(', ')
    .map(t => t.trim())

async function solve(inputFilePath: string) {
    const data = await fs.readFile(inputFilePath)
    console.log(timer.lap())
    const rows = data.toString()
        .split('\n')
        .map(t => t.trim())
        .filter(Boolean)
        .map((row, rowNum) => {
            let first = 0
            let last = 0
            let bestIndex: number = -1

            digits.forEach((digit, i) => {
                // first index
                let index = row.indexOf(digit)
                if (index != -1) {
                    if (bestIndex === -1 || bestIndex > index) {
                        first = digit.length == 1 ? +digit : i + 1
                        bestIndex = index
                    }
                }
            })

            if (bestIndex === -1) {
                throw new Error(`digit not found in row ${rowNum}`)
            }

            bestIndex = -1
            digits.forEach((digit, i) => {
                // first index
                let index = row.lastIndexOf(digit)
                if (index != -1) {
                    if (bestIndex === -1 || bestIndex < index) {
                        last = digit.length == 1 ? +digit : i + 1
                        bestIndex = index
                    }
                }
            })

            if (bestIndex === -1) {
                throw new Error(`digit not found in row ${rowNum}`)
            }

            return 10 * first + last
        })
    const result = rows.reduce(com.sum, 0)

    return [result, timer.elapsed()]
}

const argv = yargs(process.argv.slice(2))
    .options({
        file: {
            type: 'string',
            alias: 'f',
            describe: 'Path to the input file',
            demandOption: "input file is missing"
        }
    }).parseSync()

async function main() {
    const solution = await solve(argv.file)
    console.log(...solution)
}

main()
