import types = require('./types');

declare function parallel(tasks: types.AnyFn[], cb?: types.AnyFn): void;

export = parallel;
