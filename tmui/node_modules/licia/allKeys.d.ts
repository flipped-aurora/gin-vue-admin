declare namespace allKeys {
    interface IOptions {
        prototype?: boolean;
        unenumerable?: boolean;
    }
}
declare function allKeys(
    obj: any,
    options: { symbol: true } & allKeys.IOptions
): Array<string | Symbol>;
declare function allKeys(
    obj: any,
    options?: ({ symbol: false } & allKeys.IOptions) | allKeys.IOptions
): string[];

export = allKeys;
