import each = require('./each');
import types = require('./types');

declare class Select {
    constructor(selector: string | Element | Document);
    find(selector: string): Select;
    each(fn: types.AnyFn): Select;
}

export = Select;
