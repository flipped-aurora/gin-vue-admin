import types = require('./types');

declare function abbrev(...names: string[]): types.PlainObj<string>;

export = abbrev;
