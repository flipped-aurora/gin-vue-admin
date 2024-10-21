import types = require('./types');

declare function before<T extends types.AnyFn>(n: number, fn: T): T;

export = before;
