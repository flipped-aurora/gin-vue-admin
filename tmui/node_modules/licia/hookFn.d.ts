import types = require('./types');

declare function hookFn<T>(
    fn: T,
    options: {
        before?: types.AnyFn;
        after?: types.AnyFn;
        error?: types.AnyFn;
    }
): T;

export = hookFn;
