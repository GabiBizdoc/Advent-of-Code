import * as fs from 'node:fs/promises'
import * as console from "node:console";
import yargs from "yargs";
import {com, Timer} from "../../com";

const timer = new Timer()

function parseRow(row: string) {
    const [left, right] = row.split('|')
    const [cards, nums] = left.split(':')
    const winningNumbers = nums.trim().split(/\s+/).map(Number)
    const extractedNumbers = right.trim().split(' ').map(Number)
    const cardId = Number(cards.split(/\s+/)[1])

    return [cardId, winningNumbers, extractedNumbers] as const
}

async function solve(inputFilePath: string)  {
    const data = await fs.readFile(inputFilePath)
    console.log(timer.lap())
    const rows = data.toString().split("\n").map(t => t.trim()).filter(Boolean)

    let totalPoints = 0
    rows.map(parseRow).forEach(row => {
        const [cardId, winningNumbers, extractedNumbers] = row

        let local = 0
        for (const num of winningNumbers) {
            const ind = extractedNumbers.findIndex((t: number) => t === num)
            if (ind != -1) {
                extractedNumbers.splice(ind, 1)
                if (local === 0) {
                    local = 1
                } else  {
                    local *= 2
                }
            }
        }
        totalPoints += local
    })

    return [totalPoints, timer.elapsed()]
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
