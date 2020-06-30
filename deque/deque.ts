import {IIterator, Iterator} from "./iterator";

interface Iterable<Item> {
    iterator(): IIterator<Item>;
}

export interface INode<Item> {
    item: Item;
    prev: INode<Item>;
    next: INode<Item>;
}

class Deque<Item> implements Iterable<Item>{
    private first: INode<Item>;
    private last: INode<Item>;
    private n = 0;

    get isEmpty(): boolean {
        return this.n === 0;
    }

    get size(): number {
        return this.n;
    }

    iterator(): IIterator<Item> {
        return new Iterator<Item>(this.first);
    }

    addFirst(item: Item) {
        const old = this.first;
        this.first = {
            item,
            prev: this.last,
            next: old,
        };
        this.first.next.prev = this.first;

        this.n += 1;
    }

    addLast(item: Item) {
        const old = this.last;
        this.last = {
            item,
            prev: old,
            next: this.first,
        }
        this.last.prev.next = this.last;
        this.n += 1;
    }

    removeFirst(): Item {
        const item = this.first.item;
        this.n -= 1;

        if (this.isEmpty) {
            this.first = null;
            this.last = null;
        } else {
            this.first = this.first.next;
            this.first.prev = null;
        }

        return item;
    }

    removeLast(): Item {
        const item = this.last.item;
        this.n -= 1;

        if (this.isEmpty) {
            this.first = null;
            this.last = null;
        } else {
            this.last = this.last.prev;
            this.last.next = null;
        }

        return item;
    }


}

// class Deque<Item> implements Iterable<Item> {
//     iterator(): IIterator<Item> | null {
//         return new Iterator();
//     }
// }

// const it = new Iterator<number>();
