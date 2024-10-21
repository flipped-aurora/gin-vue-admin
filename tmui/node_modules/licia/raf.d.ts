import types = require('./types');

declare namespace raf {
    function cancel(id: number): void;
}
declare function raf(cb: types.AnyFn): number;

export = raf;
