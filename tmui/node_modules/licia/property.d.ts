import types = require('./types');

declare function property(path: string | string[]): types.AnyFn;

export = property;
