import types = require('./types');

declare function findIdx(arr: any[], predicate: types.AnyFn): number;

export = findIdx;
