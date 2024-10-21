import types = require('./types');

declare function strTpl(str: string, data: types.PlainObj<any>): string;

export = strTpl;
