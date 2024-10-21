import types = require('./types');

declare function optimizeCb(
    fn: types.AnyFn,
    ctx: any,
    argCount?: number
): types.AnyFn;

export = optimizeCb;
