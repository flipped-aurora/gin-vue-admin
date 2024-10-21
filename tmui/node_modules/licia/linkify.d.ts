import types = require('./types');

declare function linkify(str: string, hyperlink?: types.AnyFn): string;

export = linkify;
