import types = require('./types');

declare function callbackify(fn: types.AnyFn): types.AnyFn;

export = callbackify;
