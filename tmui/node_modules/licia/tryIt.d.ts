import types = require('./types');

declare function tryIt(fn: types.AnyFn, cb?: types.AnyFn): void;

export = tryIt;
