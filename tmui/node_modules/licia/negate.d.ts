import types = require('./types');

declare function negate<T extends types.AnyFn>(predicate: T): T;

export = negate;
