import $safeEls = require('./$safeEls');
import types = require('./types');

declare const $event: {
    on(
        element: $safeEls.El,
        event: string,
        selector: string,
        handler: types.AnyFn
    ): void;
    on(element: $safeEls.El, event: string, handler: types.AnyFn): void;
    off(
        element: $safeEls.El,
        event: string,
        selector: string,
        handler: types.AnyFn
    ): void;
    off(element: $safeEls.El, event: string, handler: types.AnyFn): void;
};

export = $event;
