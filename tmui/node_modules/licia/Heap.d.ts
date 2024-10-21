import types = require('./types');

declare class Heap {
    size: number;
    constructor(cmp?: types.AnyFn);
    clear(): void;
    add(item: any): number;
    poll(): any;
    peek(): any;
}

export = Heap;
