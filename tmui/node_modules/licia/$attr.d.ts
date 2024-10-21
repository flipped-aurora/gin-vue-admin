import $safeEls = require('./$safeEls');
import types = require('./types');

declare namespace $attr {
    function remove(element: $safeEls.El, name: string): void;
}
declare function $attr(element: $safeEls.El, name: string, value: string): void;
declare function $attr(
    element: $safeEls.El,
    attributes: types.PlainObj<string>
): void;
declare function $attr(element: $safeEls.El, name: string): string;

export = $attr;
