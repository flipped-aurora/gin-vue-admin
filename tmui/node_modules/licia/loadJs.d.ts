import types = require('./types');

declare function loadJs(src: string, cb?: types.AnyFn): void;

export = loadJs;
