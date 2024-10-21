import has = require('./has');
import types = require('./types');

declare function memoize(fn: types.AnyFn, hashFn?: types.AnyFn): types.AnyFn;

export = memoize;
