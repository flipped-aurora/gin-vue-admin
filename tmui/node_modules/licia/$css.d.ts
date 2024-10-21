import $safeEls = require('./$safeEls');
import types = require('./types');

declare function $css(element: $safeEls.El, name: string): string;
declare function $css(element: $safeEls.El, name: string, val: string): void;
declare function $css(
    element: $safeEls.El,
    properties: types.PlainObj<string | number>
): void;

export = $css;
