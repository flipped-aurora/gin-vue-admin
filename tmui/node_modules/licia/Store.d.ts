import Emitter = require('./Emitter');
import each = require('./each');

declare class Store extends Emitter {
    constructor(data?: {});
    set(key: string, val: any): void;
    set(values: {}): void;
    get(key: string): any;
    get(keys: string[]): {};
    remove(key: string): void;
    remove(keys: string[]): void;
    clear(): void;
    each(fn: (...args: any[]) => void): void;
}

export = Store;
