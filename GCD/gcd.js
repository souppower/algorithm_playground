function calc(a, b) {
    let aa = a
    let bb = b
    while (0 < bb) {
        const r = aa % bb
        aa = bb
        bb = r
    }
    return aa
}

const val = calc(120, 12)
console.log(val);