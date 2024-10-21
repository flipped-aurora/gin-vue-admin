import { type ElementNode, type RootNode } from '@vue/compiler-core';
export declare function parseVueCode(code: string, isNVue?: boolean): {
    code: string;
    files?: undefined;
    errors?: undefined;
} | {
    code: string;
    files: string[];
    errors: SyntaxError[];
};
export declare function parseBlockCode(ast: RootNode, code: string): string;
export declare function parseWxsNodes(ast: RootNode): ElementNode[];
export declare function parseWxsCode(wxsNodes: ElementNode[], code: string): string;
