export declare function isWxs(id: string): boolean;
export declare function isSjs(id: string): boolean;
export declare function isRenderjs(id: string): boolean;
type FilterType = 'wxs' | 'renderjs' | 'sjs';
export declare function parseRenderjs(id: string): {
    type: FilterType;
    name: string;
    filename: string;
} | {
    readonly type: "";
    readonly name: "";
    readonly filename: "";
};
export declare function missingModuleName(type: FilterType, code: string): string;
export declare function parseFilterNames(lang: string, code: string): string[];
export {};
