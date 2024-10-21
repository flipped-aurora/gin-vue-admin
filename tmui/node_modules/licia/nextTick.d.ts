import types = require('./types');

declare function nextTick(cb: types.AnyFn): void;

export = nextTick;
