import types = require('./types');

declare function timeTaken(fn: types.AnyFn): number;

export = timeTaken;
