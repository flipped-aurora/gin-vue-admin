import types = require('./types');

declare function isSorted(arr: any[], cmp?: types.AnyFn): boolean;

export = isSorted;
