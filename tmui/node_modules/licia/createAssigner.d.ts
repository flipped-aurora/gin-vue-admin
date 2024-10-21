import types = require('./types');

declare function createAssigner(
    keysFn: types.AnyFn,
    defaults: boolean
): types.AnyFn;

export = createAssigner;
