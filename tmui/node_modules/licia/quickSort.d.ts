import types = require('./types');

declare function quickSort(arr: any[], cmp?: types.AnyFn): any[];

export = quickSort;
