import types = require('./types');

declare function fnParams(fn: types.AnyFn | string): string[];

export = fnParams;
