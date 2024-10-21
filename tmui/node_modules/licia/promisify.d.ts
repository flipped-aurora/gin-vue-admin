import types = require('./types');

declare function promisify(fn: types.AnyFn, multiArgs?: boolean): types.AnyFn;

export = promisify;
