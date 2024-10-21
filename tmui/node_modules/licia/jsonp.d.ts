import types = require('./types');

declare function jsonp(options: {
    url: string;
    data?: any;
    success?: types.AnyFn;
    param?: string;
    name?: string;
    error?: types.AnyFn;
    complete?: types.AnyFn;
    timeout?: number;
}): void;

export = jsonp;
