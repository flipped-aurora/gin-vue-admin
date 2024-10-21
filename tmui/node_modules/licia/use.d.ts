import types = require('./types');

declare function use(requires: string[], method: types.AnyFn): void;
declare function use(method: types.AnyFn): void;

export = use;
