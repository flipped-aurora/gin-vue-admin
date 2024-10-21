import types = require('./types');

declare function some<T>(
    list: types.List<T>,
    iterator?: types.ListIterator<T, boolean>,
    context?: any
): boolean;
declare function some<T>(
    object: types.Dictionary<T>,
    iterator?: types.ObjectIterator<T, boolean>,
    context?: any
): boolean;

export = some;
