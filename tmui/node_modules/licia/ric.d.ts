import types = require('./types');

declare namespace ric {
    function cancel(id: number): void;
}
declare function ric(cb: types.AnyFn): number;

export = ric;
