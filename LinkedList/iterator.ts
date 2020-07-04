import {INode} from "./SLList";

export interface IIterator<Item> {
    hasNext(): boolean;
    next(): Item;
}

export class Iterator<Item> implements IIterator<Item> {
    private current: INode<Item>;

    constructor(private node: INode<Item>) {
        this.current = node;
    }

    hasNext(): boolean {
        return this.current !== null;
    }

    next(): Item {
        if (!this.hasNext()) {
            return null;
        }

        const item = this.current.item;
        this.current = this.current.next;
        return item;
    }
}
