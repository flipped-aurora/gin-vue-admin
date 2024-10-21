import types = require('./types');

declare function reduce<T, TResult>(
    list: types.List<T>,
    iterator: types.MemoIterator<T, TResult>,
    memo?: TResult,
    context?: any
): TResult;
declare function reduce<T, TResult>(
    list: types.Dictionary<T>,
    iterator: types.MemoObjectIterator<T, TResult>,
    memo?: TResult,
    context?: any
): TResult;

export = reduce;
