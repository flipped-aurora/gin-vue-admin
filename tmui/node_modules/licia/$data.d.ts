import $safeEls = require('./$safeEls');
import types = require('./types');

declare function $data(element: $safeEls.El, name: string, value: string): void;
declare function $data(
    element: $safeEls.El,
    attributes: types.PlainObj<string>
): void;
declare function $data(element: $safeEls.El, name: string): string;

export = $data;
