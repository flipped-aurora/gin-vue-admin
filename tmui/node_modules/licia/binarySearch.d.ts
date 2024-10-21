import types = require('./types');

declare function binarySearch(
    array: any[],
    val: any,
    cmp?: types.AnyFn
): number;

export = binarySearch;
