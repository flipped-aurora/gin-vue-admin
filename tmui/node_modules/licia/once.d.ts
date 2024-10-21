import types = require('./types');

declare function once(fn: types.AnyFn): types.AnyFn;

export = once;
