function insertion(a, v) {
    if (a.length === 0) {
        a.push(v)
        return;
    }

    let i = a.length-1

    while (true) {
        if (i < 0 || a[i] <= v) break

        // slide by one
        a[i+1] = a[i]
        i--
    }

    a[i+1] = v
}

const arr = [1,2,4,5,6,7]
insertion(arr, 3)
console.log(arr);