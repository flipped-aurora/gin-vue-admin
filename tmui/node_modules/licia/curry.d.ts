import types = require('./types');

declare function curry(fn: types.AnyFn): types.AnyFn;

export = curry;
