import reduce = require('./reduce');
import types = require('./types');

declare function reduceRight<T, TResult>(
    list: types.Collection<T>,
    iterator: types.MemoIterator<T, TResult>,
    memo?: TResult,
    context?: any
): TResult;

export = reduceRight;
