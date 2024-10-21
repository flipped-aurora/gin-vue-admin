import types = require('./types');

declare class Stack {
    size: number;
    clear(): void;
    push(item: any): number;
    pop(): any;
    peek(): any;
    forEach(iterator: types.AnyFn, context?: any): void;
    toArr(): any[];
}

export = Stack;
