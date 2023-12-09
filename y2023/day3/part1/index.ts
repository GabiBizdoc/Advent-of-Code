import fs from 'node:fs/promises'
import * as console from "console";
import {com} from "../../com";


class Point {
    constructor(public line: number, public col: number) {
    }
}

function getValue(grid: string[][], l: number, c: number) {
    const is = l >= 0 && l < grid.length && c >= 0 && c < grid[0].length
    if (is) {
        const value = grid[l][c]
        return [value, is] as const
    }
    return [null, false] as const
}

async function main() {
    const data = await fs.readFile("./p3/i2.txt")
    const rows = data.toString().split("\n").map(t => t.trim()).filter(Boolean)
    const grid = rows.map(row => row.split(''))

    // const nums = []
    // let status = new Info(false, -1)
    const tmp: Array<Point>[] = []

    let found: Array<Point> = []
    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[i].length; j++) {
            const c = grid[i][j]
            if (com.isDigit(c)) {
                found.push(new Point(i, j))
            } else {
                if (found.length > 0) {
                    tmp.push(found)
                    found = []
                }
            }
        }
        if (found.length > 0) {
            tmp.push(found)
            found = []
        }
    }
    if (found.length > 0) {
        tmp.push(found)
        found = []
    }

    return tmp.filter(found => {
        const first = found.at(0)!
        const last = found.at(-1)!

        let isSolution = false

        function check(t: ReturnType<typeof getValue>) {
            const [value, is] = t
            if (is) {
                if (value !== '.') {
                    isSolution = true
                }
            }
        }

        check(getValue(grid, first.line - 1, first.col - 1))
        check(getValue(grid, first.line + 1, first.col - 1))
        check(getValue(grid, first.line, first.col - 1))

        check(getValue(grid, last.line - 1, last.col + 1))
        check(getValue(grid, last.line + 1, last.col + 1))
        check(getValue(grid, last.line, last.col + 1))


        for (const f of found) {
            check(getValue(grid, f.line + 1, f.col))
            check(getValue(grid, f.line - 1, f.col))
        }

        return isSolution
    }).map(found => {
        return found.reduce((a, c) => a * 10 + Number(grid[c.line][c.col]), 0)
    }).reduce(com.sum, 0)
}

main().then(console.log)
