import types = require('./types');

declare function mapObj<T, TResult>(
    object: types.Dictionary<T>,
    iterator: types.ObjectIterator<T, TResult>,
    context?: any
): types.Dictionary<TResult>;

export = mapObj;
