import types = require('./types');

declare function map<T, TResult>(
    list: types.List<T>,
    iterator: types.ListIterator<T, TResult>,
    context?: any
): TResult[];
declare function map<T, TResult>(
    object: types.Dictionary<T>,
    iterator: types.ObjectIterator<T, TResult>,
    context?: any
): TResult[];

export = map;
