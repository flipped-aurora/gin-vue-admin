import types = require('./types');

declare function scrollTo(
    target: Element | string | number,
    options: {
        tolerance?: number;
        duration?: number;
        easing?: string | Function;
        callback?: types.AnyFn;
    }
);

export = scrollTo;
