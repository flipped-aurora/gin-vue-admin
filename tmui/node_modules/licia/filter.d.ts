import types = require('./types');

declare function filter<T>(
    list: types.List<T>,
    iterator: types.ListIterator<T, boolean>,
    context?: any
): T[];
declare function filter<T>(
    object: types.Dictionary<T>,
    iterator: types.ObjectIterator<T, boolean>,
    context?: any
): T[];

export = filter;
