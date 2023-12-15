function delay(n) {
    return new Promise(resolve => setTimeout(resolve, n))
}

async function doSomething() {
    while (true) {
        console.log("ping")
        await delay(1000)
    }
}

async function main() {
    const somethingPromise = new Promise(async resolve => {
        await doSomething()
        resolve("somethingPromise is done!")
    })

    const waitPromise = new Promise(async resolve => {
        await delay(5 * 1000)
        resolve("waitPromise is done!")
    })

    return await Promise.race([somethingPromise, waitPromise])
}

main().then(console.log)