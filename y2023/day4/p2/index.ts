import * as fs from 'node:fs/promises'
import * as console from "node:console";
import yargs from "yargs";
import {Timer} from "../../com";

const timer = new Timer()

function parseRow(row: string) {
    const [left, right] = row.split('|')
    const [cards, nums] = left.split(':')
    const winningNumbers = nums.trim().split(/\s+/).map(Number)
    const extractedNumbers = right.trim().split(' ').map(Number)
    const cardId = Number(cards.split(/\s+/)[1])

    return [cardId, winningNumbers, extractedNumbers] as const
}

async function solve(inputFilePath: string) {
    const data = await fs.readFile(inputFilePath)
    console.log(timer.lap())
    const rows = data.toString().split("\n").map(t => t.trim()).filter(Boolean)

    const cards = new Map<number, number>()
    const clones = new Map<number, number>()

    let totalPoints = 0
    rows.map(parseRow).forEach(row => {
        const [cardId, winningNumbers, extractedNumbers] = row

        let points = 0
        for (const num of winningNumbers) {
            const ind = extractedNumbers.findIndex((t: number) => t === num)
            if (ind != -1) {
                extractedNumbers.splice(ind, 1)
                points += 1
            }
        }

        cards.set(cardId, points)
        for (let i = 1; i <= points; i++) {
            const old = clones.get(cardId + i) || 0
            clones.set(cardId + i, old + 1)
        }

        totalPoints += 1
    })

    for (const cardId of clones.keys()) {
        let points = cards.get(cardId)!
        for (let j = 1; j <= points; j++) {
            const nextCardId = cardId + j
            if (clones.has(nextCardId)) {
                clones.set(nextCardId, clones.get(nextCardId)! + clones.get(cardId)!)
            }
        }
        totalPoints += clones.get(cardId)!
    }

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
