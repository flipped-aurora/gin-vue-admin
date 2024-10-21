import type { Plugin } from 'vite';
export declare function isSrcImport(code: string): boolean;
export declare function isSrcImportVue(code: string): boolean;
export declare function uniViteSfcSrcImportPlugin({ onlyVue }?: {
    onlyVue: boolean;
}): Plugin;
