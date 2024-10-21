import types = require('./types');

declare function bind(fn: types.AnyFn, ctx: any, ...args: any[]): types.AnyFn;

export = bind;
