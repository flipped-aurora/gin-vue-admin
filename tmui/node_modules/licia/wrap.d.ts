import types = require('./types');

declare function wrap(fn: types.AnyFn, wrapper: types.AnyFn): types.AnyFn;

export = wrap;
