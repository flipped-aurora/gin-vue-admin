declare function golangify<T, U = Error>(
    fn: (...args: any[]) => Promise<T>
): (...args: any[]) => Promise<[T | undefined, U | null]>;
declare function golangify<T, U = Error>(
    p: Promise<T>
): Promise<[T | undefined, U | null]>;

export = golangify;
