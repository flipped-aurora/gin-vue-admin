import types = require('./types');

declare function universalify(fn: types.AnyFn, type: string): types.AnyFn;

export = universalify;
