import types = require('./types');

declare function partial(fn: types.AnyFn, ...partials: any[]): types.AnyFn;

export = partial;
