import types = require('./types');

declare function deprecate(fn: types.AnyFn, msg: string): types.AnyFn;

export = deprecate;
