import types = require('./types');

declare function delay(fn: types.AnyFn, wait: number, ...args: any[]): void;

export = delay;
