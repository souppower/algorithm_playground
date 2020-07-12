function sort(arr: number[], lo: number, hi: number) {
    if (hi <= lo) return;

    let lt = lo;
    let gt = hi;

    const v = arr[lo]

    let i = lo;

    while (i <= gt) {
        if (arr[i] < v) swap(arr, lt++, i++);
        else if (arr[i] > v) swap(arr, i, gt--);
        else i++;
    }

    sort(arr, lo, lt-1);
    sort(arr, gt+1, hi);
    return arr;
}

function swap(arr: number[], i: number, j: number) {
    const tmp = arr[i];
    arr[i] = arr[j];
    arr[j] = tmp;
}

const arr = [19, 10, 30, 2, 100, 4, 20, 1, 5];
console.log(`before sort:`);
console.log(arr);
const sorted = sort(arr, 0, arr.length-1)
console.log(`after sort:`);
console.log(sorted);
