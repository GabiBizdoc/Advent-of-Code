export module com {
    export const sum = (a: any, b: any) => a + b
    export const isBetween = <T extends string | number>(a: T, b: T) => (c: T) => c >= a && c <= b
    export const isDigit = isBetween<string>('0', '9')
}

export class Timer {
    private start!: number

    constructor() {
        this.reset()
    }

    reset() {
        this.start = this.now
    }

    lap() {
        const elapsed = this.elapsed()
        this.reset()
        return elapsed
    }

    elapsed() {
        const elapsed = this.now - this.start
        return elapsed.toFixed(6) + ' ms'
    }

    private get now() {
        return performance.now()
    }
}

