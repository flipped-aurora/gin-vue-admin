import has = require('./has');

declare class HashTable {
    constructor(size?: number);
    set(key: string, val: any): void;
    get(key: string): any;
    has(key: string): boolean;
    delete(key: string): void;
}

export = HashTable;
