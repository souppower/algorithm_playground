function sort(arr: number[], lo: number, hi: number) {
    if (hi <= lo) return;
    const pivot = partition(arr, lo, hi);
    sort(arr, lo, pivot-1);
    sort(arr, pivot+1, hi);
    return arr;
}

function partition(arr: number[], lo: number, hi: number): number {
    let left = lo;
    let right = hi + 1;

    while (true) {
        while (arr[++left] < arr[lo])  {
            if(left === hi) break;
        }
        while (arr[lo] < arr[--right]) {
            if(right === lo) break;
        }

        if (left >= right) break;

        swap(arr, left, right);
    }


    // swap j with partitioning item
    swap(arr, lo, right)

    // return index of item now know to be in place
    return right;
}

function swap(arr: number[], i: number, j: number) {
    const tmp = arr[i];
    arr[i] = arr[j];
    arr[j] = tmp;
}

const arr = [19, 10, 30, 2, 4, 20, 1, 5];
console.log(`before sort:`);
console.log(arr);
const sorted = sort(arr, 0, arr.length-1)
console.log(`after sort:`);
console.log(sorted);
