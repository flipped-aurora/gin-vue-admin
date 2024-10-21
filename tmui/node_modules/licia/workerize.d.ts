import types = require('./types');

declare function workerize(fn: types.AnyFn): types.AnyFn;

export = workerize;
