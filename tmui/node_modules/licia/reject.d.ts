import types = require('./types');

declare function reject<T>(
    list: types.List<T>,
    iterator: types.ListIterator<T, boolean>,
    context?: any
): T[];
declare function reject<T>(
    object: types.Dictionary<T>,
    iterator: types.ObjectIterator<T, boolean>,
    context?: any
): T[];

export = reject;
