import types = require('./types');

declare function safeCb(val?: any, ctx?: any, argCount?: number): types.AnyFn;

export = safeCb;
