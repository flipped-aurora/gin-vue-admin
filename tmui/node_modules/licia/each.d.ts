import types = require('./types');

declare function each<T>(
    list: types.List<T>,
    iterator: types.ListIterator<T, void>,
    ctx?: any
): types.List<T>;
declare function each<T>(
    object: types.Dictionary<T>,
    iterator: types.ObjectIterator<T, void>,
    ctx?: any
): types.Collection<T>;

export = each;
