import types = require('./types');

declare function remove<T, TResult>(
    list: types.List<T>,
    iterator: types.ListIterator<T, boolean>,
    context?: any
): TResult[];

export = remove;
