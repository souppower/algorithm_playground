function merge(a1, a2) {
    const a = []
    const t = [...a1, ...a2.reverse()]

    let i = 0
    let j = t.length-1

    for (let k=0; k < t.length; k++) {
        if (t[i] <= t[j]) {
            a[k] = t[i]
            i++
        } else {
            a[k] = t[j]
            j--
        }
    }

    return a
}

const arr1 = [1,5,11]
const arr2 = [2,4,9,12]
const sorted = merge(arr1,arr2)
console.log(sorted);