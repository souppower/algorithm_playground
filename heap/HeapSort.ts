class HeapSort {
    public static sort(pq: number[]) {
        // heapify
        let N = pq.length-1
        for (let k=N/2|0; k >= 0; k--) {
            this.sink(pq, k, N)
        }

        // in-array sort
        while (N > 0) {
            this.swap(pq, 0, N)
            this.sink(pq, 0, --N)
        }
    }

    private static sink(pq: number[], k :number, N: number) {
        while(2*k+1 <= N) {
            let j = 2*k+1
            // children of node at k are 2k+1 and 2k+2
            if (j < N && this.less(pq, j, j+1)) j++
            if (!this.less(pq, k, j)) break
            this.swap(pq, k, j)
            k = j
        }
    }

    private static less(pq, i, j) {
        return pq[i] < pq[j]
    }

    private static swap(pq, i, j) {
        const tmp = pq[i]
        pq[i] = pq[j]
        pq[j] = tmp;
    }
}


const arr = [100,2,10,3,4,20,5,6]
HeapSort.sort(arr)
console.log(arr);
