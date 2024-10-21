import types = require('./types');

declare function compose(...fn: types.AnyFn[]): types.AnyFn;

export = compose;
