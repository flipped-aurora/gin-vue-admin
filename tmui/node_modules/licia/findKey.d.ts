import types = require('./types');

declare function findKey(
    obj: any,
    predicate: types.AnyFn,
    ctx?: any
): string | void;

export = findKey;
