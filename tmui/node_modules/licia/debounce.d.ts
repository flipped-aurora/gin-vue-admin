import types = require('./types');

declare function debounce<T extends types.AnyFn>(fn: T, wait: number): T;

export = debounce;
