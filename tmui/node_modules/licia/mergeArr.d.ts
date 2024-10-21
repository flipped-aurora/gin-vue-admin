declare function mergeArr<T, U>(
    first: ArrayLike<T>,
    ...arrays: ArrayLike<U>[]
): ArrayLike<T | U>;

export = mergeArr;
