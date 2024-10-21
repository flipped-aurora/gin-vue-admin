import types = require('./types');

declare function every<T>(
    object: types.List<T>,
    iterator?: types.ListIterator<T, boolean>,
    context?: any
): boolean;
declare function every<T>(
    object: types.Dictionary<T>,
    iterator?: types.ObjectIterator<T, boolean>,
    context?: any
): boolean;

export = every;
