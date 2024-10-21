import types = require('./types');

declare function sortBy(arr: any, iterator?: types.AnyFn, ctx?: any): any[];

export = sortBy;
