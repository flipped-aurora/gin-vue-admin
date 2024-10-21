import types = require('./types');

declare function define(
    name: string,
    requires: string[],
    method: types.AnyFn
): void;
declare function define(name: string, method: types.AnyFn): void;

export = define;
