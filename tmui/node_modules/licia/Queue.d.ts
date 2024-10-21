import types = require('./types');

declare class Queue {
    size: number;
    clear(): void;
    enqueue(item: any): number;
    dequeue(): any;
    peek(): any;
    forEach(iterator: types.AnyFn, context?: any): void;
    toArr(): any[];
}

export = Queue;
