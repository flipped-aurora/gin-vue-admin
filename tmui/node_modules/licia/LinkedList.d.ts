import types = require('./types');

declare namespace LinkedList {
    class Node {
        value: any;
        prev: Node | null;
        next: Node | null;
    }
}
declare class LinkedList {
    size: number;
    head: LinkedList.Node;
    tail: LinkedList.Node;
    push(val: any): number;
    pop(): any;
    unshift(val: any): number;
    shift(): any;
    find(fn: types.AnyFn): LinkedList.Node | void;
    delNode(node: LinkedList.Node): void;
    forEach(iterator: types.AnyFn, ctx?: any);
    toArr(): any[];
}

export = LinkedList;
