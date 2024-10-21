import types = require('./types');

declare function find<T>(
    object: types.List<T>,
    iterator: types.ListIterator<T, boolean>,
    context?: any
): T | undefined;
declare function find<T>(
    object: types.Dictionary<T>,
    iterator: types.ObjectIterator<T, boolean>,
    context?: any
): T | undefined;

export = find;
