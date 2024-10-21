import types = require('./types');

declare function sortKeys(
    obj: object,
    options?: {
        deep?: boolean;
        comparator?: types.AnyFn;
    }
): object;

export = sortKeys;
