import fs from "node:fs/promises";
import console from "node:console";
import {Timer} from "../../com";
import yargs from "yargs";

function equation(k: number, time: number, distance: number) {
    return k * (time - k) - distance;
}

function findMinMaxK(time: number, distance: number) {

    let cnt = 0
    for (let i = 1; i < time; i++) {
        if (equation(i, time, distance) > 0) {
            cnt += 1
        }
    }
    console.log(time, distance, cnt)

    return cnt
}

const timer = new Timer()

async function solve(inputFilePath: string) {
    const data = await fs.readFile(inputFilePath)
    console.log(timer.lap())

    const rows = data.toString()
        .split('\n')
        .map(t => t.trim())
        .filter(Boolean)
        .map(row => {
           return row.split(":")[1]
               .split(/\s+/)
               .filter(t => t !== '')
               .map(Number)
        })

    const [timeRow, distanceRow] = rows

    const result = timeRow.map((time, i) => {
        const distance = distanceRow[i]

        return findMinMaxK(time, distance)
    }).reduce((a, b) => a * b, 1)

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
