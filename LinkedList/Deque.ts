import {IIterator, Iterator} from "./iterator";

interface Iterable<Item> {
    iterator(): IIterator<Item>;
}

export interface INode<Item> {
    item: Item | null;
    next: INode<Item>;
}

class SLList<Item> implements Iterable<Item>{
    public size = 0;
    sentinel: INode<Item>;

    constructor() {
        this.sentinel = {
            item: null,
            next: null,
        }
    }

    get isEmpty(): boolean {
        return this.size === 0;
    }

    iterator(): IIterator<Item> {
        return new Iterator<Item>(this.sentinel);
    }

    addFirst(item: Item) {
        const newFirst = {
            item,
            next: this.sentinel.next,
        }
        this.sentinel.next = newFirst;

        this.size += 1;
    }

    addLast(item: Item) {
        let node = this.sentinel;
        while (node.next !== null) {
            node = node.next;
        }
        node.next = {
            item,
            next: null,
        }
        this.size += 1;
    }

    removeFirst(): Item {
        if (this.sentinel.next === null) {
            return null;
        }
        const item = this.sentinel.next.item;
        this.sentinel.next =  this.sentinel.next.next;
        this.size -= 1;

        return item;
    }

    removeLast(): Item {
        if (this.sentinel.next === null) {
            return null;
        }

        let node = this.sentinel;
        let prev = this.sentinel;
        while (node.next !== null) {
            prev = node;
            node = node.next;
        }

        const item = prev.next.item;
        prev.next = null;
        return item;
    }

}

function log(list: any) {
    console.log(JSON.stringify(list, null, 2))
}
