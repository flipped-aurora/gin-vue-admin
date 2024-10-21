import types = require('./types');

declare const delegate: {
    add(el: Element, type: string, selector: string, cb: types.AnyFn): void;
    remove(el: Element, type: string, selector: string, cb: types.AnyFn): void;
};

export = delegate;
