class MaxPQ {
    private N = 0
    public pq: number[] = []

    public get isEmpty() {
        return this.N === 0
    }

    public insert(value: number) {
        this.pq[++this.N] = value
        this.swim(this.N)
    }

    public swim(k: number) {
        // parent of node at k is at k/2
        while(k > 1 && this.less(k/2|0, k)) {
            this.swap(k, k/2|0)
            k = k/2|0
        }
    }

    public sink(k :number) {
        while(2*k <= this.N) {
            let j = 2*k
            // children of node at k are 2k and 2k+1
            if (j < this.N && this.less(j, j+1)) j++
            if (!this.less(k, j)) break
            this.swap(k, j)
            k = j
        }
    }

    public delMax() {
        if (this.isEmpty) return

        const max = this.pq[1]
        this.swap(1, this.N--)
        this.sink(1)
        delete this.pq[this.N+1]
        return max
    }

    private less(i, j) {
        return this.pq[i] < this.pq[j]
    }

    private swap(i, j) {
        const tmp = this.pq[i]
        this.pq[i] = this.pq[j]
        this.pq[j] = tmp;
    }
}

const pq = new MaxPQ()
console.log(`isEmpty: ${pq.isEmpty}`);
pq.insert(10)
console.log(`isEmpty: ${pq.isEmpty}`);
pq.insert(20)
pq.insert(30)
pq.insert(40)
pq.insert(50)
console.log(pq.pq);

const max = pq.delMax()
console.log(max);