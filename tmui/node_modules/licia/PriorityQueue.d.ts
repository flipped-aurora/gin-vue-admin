import types = require('./types');

declare class PriorityQueue {
    size: number;
    constructor(cmp?: types.AnyFn);
    clear(): void;
    enqueue(item: any): number;
    dequeue(): any;
    peek(): any;
}

export = PriorityQueue;
