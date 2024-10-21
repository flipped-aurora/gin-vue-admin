import types = require('./types');

declare function h(
    tag: string,
    attrs?: types.PlainObj<any>,
    ...child: Array<string | HTMLElement>
): HTMLElement;

export = h;
