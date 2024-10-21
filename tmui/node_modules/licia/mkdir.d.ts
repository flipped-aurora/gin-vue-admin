import types = require('./types');

declare namespace mkdir {
    function sync(dir: string, mode?: number): void;
}
declare function mkdir(dir: string, mode?: number, cb?: types.AnyFn): void;
declare function mkdir(dir: string, cb?: types.AnyFn): void;

export = mkdir;
