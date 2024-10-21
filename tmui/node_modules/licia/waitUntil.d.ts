import types = require('./types');

declare function waitUntil(
    condition: types.AnyFn,
    timeout?: number,
    interval?: number
): Promise<any>;

export = waitUntil;
