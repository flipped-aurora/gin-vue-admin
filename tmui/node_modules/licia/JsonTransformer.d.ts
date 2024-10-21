import map = require('./map');
import filter = require('./filter');
import types = require('./types');

declare class JsonTransformer {
    constructor(data: any);
    set(key: string, val: any): JsonTransformer;
    get(key?: string): any;
    map(from: string, to: string, fn: types.AnyFn): JsonTransformer;
    map(from: string, fn: types.AnyFn): JsonTransformer;
    filter(from: string, to: string, fn: types.AnyFn): JsonTransformer;
    filter(from: string, fn: types.AnyFn): JsonTransformer;
    remove(keys: string | string[]): JsonTransformer;
    compute(
        from: string | string[],
        to: string,
        fn: types.AnyFn
    ): JsonTransformer;
    compute(from: string, fn: types.AnyFn): JsonTransformer;
    toString(): string;
}

export = JsonTransformer;
