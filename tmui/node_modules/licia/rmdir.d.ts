import types = require('./types');

declare function rmdir(dir: string, cb?: types.AnyFn): void;

export = rmdir;
