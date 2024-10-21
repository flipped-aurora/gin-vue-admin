import types = require('./types');

declare function loadCss(src: string, cb?: types.AnyFn): void;

export = loadCss;
