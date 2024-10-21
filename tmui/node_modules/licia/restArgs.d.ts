import types = require('./types');

declare function restArgs(fn: types.AnyFn, startIndex?: number): types.AnyFn;

export = restArgs;
