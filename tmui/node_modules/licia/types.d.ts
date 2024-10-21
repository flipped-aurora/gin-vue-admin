declare namespace types {
    interface Collection<T> {}
    interface List<T> extends Collection<T> {
        [index: number]: T;
        length: number;
    }
    interface ListIterator<T, TResult> {
        (value: T, index: number, list: List<T>): TResult;
    }
    interface Dictionary<T> extends Collection<T> {
        [index: string]: T;
    }
    interface ObjectIterator<T, TResult> {
        (element: T, key: string, list: Dictionary<T>): TResult;
    }
    interface MemoIterator<T, TResult> {
        (prev: TResult, curr: T, index: number, list: List<T>): TResult;
    }
    interface MemoObjectIterator<T, TResult> {
        (prev: TResult, curr: T, key: string, list: Dictionary<T>): TResult;
    }
    type Fn<T> = (...args: any[]) => T;
    type AnyFn = Fn<any>;
    type PlainObj<T> = { [name: string]: T };
}
declare const types: {};

export = types;
