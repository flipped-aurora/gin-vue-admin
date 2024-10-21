declare namespace meta {
    function remove(nameList: string | string[]): void;
}
declare function meta(): {};
declare function meta(key: string): string;
declare function meta(keys: string[]): {};
declare function meta(key, value): void;
declare function meta(pairs: {}): void;

export = meta;
