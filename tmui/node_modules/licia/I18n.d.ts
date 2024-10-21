import types = require('./types');

declare class I18n {
    constructor(locale: string, langs: types.PlainObj<any>);
    set(locale: string, lang: types.PlainObj<any>): void;
    t(path: string | string[], data?: types.PlainObj<any>): string;
    locale(locale: string): void;
}

export = I18n;
