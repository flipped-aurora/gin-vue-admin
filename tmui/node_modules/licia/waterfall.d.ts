import types = require('./types');

declare function waterfall(tasks: types.AnyFn[], cb?: types.AnyFn): void;

export = waterfall;
