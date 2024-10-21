import types = require('./types');

declare function after<T extends types.AnyFn>(n: number, fn: T): T;

export = after;
