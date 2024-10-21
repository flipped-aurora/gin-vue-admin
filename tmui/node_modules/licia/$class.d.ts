import $safeEls = require('./$safeEls');

declare const $class: {
    add(element: $safeEls.El, name: string | string[]): void;
    has(element: $safeEls.El, name: string): boolean;
    toggle(element: $safeEls.El, name: string): void;
    remove(element: $safeEls.El, name: string): void;
};

export = $class;
