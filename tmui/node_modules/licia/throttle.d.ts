import types = require('./types');

declare function throttle<T extends types.AnyFn>(fn: T, wait: number): T;

export = throttle;
