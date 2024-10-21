declare namespace stringifyAll {
    function parse(str: string): any;
}
declare function stringifyAll(
    obj: any,
    options?: {
        unenumerable?: boolean;
        symbol?: boolean;
        accessGetter?: boolean;
        timeout?: number;
        depth?: number;
        ignore?: any[];
    }
): string;

export = stringifyAll;
