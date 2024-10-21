import types = require('./types');

declare function toAsync(fn: types.AnyFn): types.AnyFn;

export = toAsync;
