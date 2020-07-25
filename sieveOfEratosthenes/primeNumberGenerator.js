function generatePrimeNumbers(N) {
    const arr = new Array(N)

    for (let i = 0; i < N; i++) {
        arr[i] = true
    }
    arr[0] = false
    arr[1] = false

    const max = Math.floor(Math.sqrt(N))

    for(let i = 2; i <= max; i++) {
        if(arr[i]) {
            for(let j = i * 2; j <= N; j += i) {
                arr[j] = false
            }
        }
    }

    return arr.reduce((acc, cur, idx) => {
        if (cur) return [...acc, idx]
        return acc
    }, [])
}

const arr = generatePrimeNumbers(100)
console.log(arr);
