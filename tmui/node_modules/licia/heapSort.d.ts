import types = require('./types');

declare function heapSort(arr: any[], cmp?: types.AnyFn): any[];

export = heapSort;
